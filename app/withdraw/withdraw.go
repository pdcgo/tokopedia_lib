package withdraw

import (
	"context"
	"crypto/sha256"
	"errors"
	"fmt"
	"strconv"
	"time"

	"github.com/chromedp/cdproto/page"
	"github.com/chromedp/cdproto/runtime"
	"github.com/chromedp/chromedp"
	"github.com/pdcgo/tokopedia_lib"
	"github.com/pdcgo/tokopedia_lib/lib/api"
	"github.com/sethvargo/go-retry"
)

var ErrSaldoKosong = errors.New("saldo kosong")
var ErrWithdraw = errors.New("withdraw error")
var ErrHashedPin = errors.New("error hashing pin")

type Withdraw struct {
	Driver *tokopedia_lib.DriverAccount

	Api         *api.TokopediaApi
	User        *api.UserDataQueryResp
	Balance     *api.WithdrawBalanceResp
	Bank        *api.GetBankWDV2
	GenerateKey *api.WindrawnGenerateKeyResp
}

func NewWithdraw(driver *tokopedia_lib.DriverAccount, api *api.TokopediaApi) (*Withdraw, error) {
	w := &Withdraw{
		Api:    api,
		Driver: driver,
	}

	err := w.InitDataWithdraw()
	if err != nil {
		return nil, err
	}

	return w, nil
}

func (w *Withdraw) InitDataWithdraw() error {
	user, err := w.Api.UserDataQuery()
	if err != nil {
		return err
	}
	w.User = user

	balance, err := w.Api.WithDrawBalance()
	if err != nil {
		return err
	}
	w.Balance = balance
	if balance.Data.MidasGetAllDepositAmount.SellerAll == 0 {
		return ErrSaldoKosong
	}

	banks, err := w.Api.BankListQuery(false)
	if err != nil {
		return err
	}

	bank := banks.Data.GetBankListWDV2.Data.GetDefaultBank()
	w.Bank = bank

	return nil
}

func (w *Withdraw) GetGenerateKey() (*api.WindrawnGenerateKeyResp, error) {
	_, err := w.Api.WithdrawOtpRequest(w.User.Data.User.Phone)
	if err != nil {
		return nil, err
	}

	generateKey, err := w.Api.WindrawnGenerateKey()
	if err != nil {
		return nil, err
	}

	w.GenerateKey = generateKey

	return generateKey, nil
}

func (w *Withdraw) StartWithdraw(pinHashed string) error {
	otpValidateVariable := api.NewOtpValidateVariable(
		w.User.Data.User.Phone,
		strconv.Itoa(w.Bank.BankAccountID),
		pinHashed, w.GenerateKey.Data.GenerateKey)
	otpValidate, err := w.Api.WithdrawOtpValidate(otpValidateVariable)
	if err != nil {
		return err
	}

	withdrawvariable := api.NewWithdrawVariable(
		w.User.Data.User,
		w.Bank, otpValidate.Data.OTPValidate,
		strconv.Itoa(w.Balance.Data.MidasGetAllDepositAmount.SellerAll))
	_, err = w.Api.WithdrawSaldoMutation(withdrawvariable)

	return err
}

func (w *Withdraw) Run() error {
	withdraw := func() error {
		err := w.Driver.Run(false, func(dctx *tokopedia_lib.DriverContext) error {
			return w.Withdraw(dctx)
		})

		return err
	}

	b := retry.NewFibonacci(time.Second)
	err := retry.Do(context.Background(), retry.WithMaxRetries(2, b), func(ctx context.Context) error {
		err := withdraw()
		if err != nil {
			return retry.RetryableError(err)
		}

		return nil
	})

	return err
}

func (wd *Withdraw) SetupWindowProperty(ctx context.Context) {
	script := `Object.defineProperties(window, {
		textEncoder: {
			get: () => new TextEncoder(),
		},
		publicKeyEncoder: {
			get: () =>  function(e) {
				for (var t = window.atob(e), n = t.length, r = new ArrayBuffer(n), i = new Uint8Array(r), o = 0, a = n; o < a; o += 1)
					i[o] = t.charCodeAt(o);
				return r
				},
		},
		encryptToString: {
			get: () => function(e) {
				for (var t = "", n = new Uint8Array(e), r = 0; r < n.byteLength; r += 1)
					t += String.fromCharCode(n[r]);
				return window.btoa(t)
			}, 
		},
		otpAlgorithm: {
			get: () => Object.assign({}, {"name": "RSA-OAEP", "hash": "SHA-256"}),
		}
	});`

	chromedp.Run(
		ctx,
		chromedp.ActionFunc(func(ctx context.Context) error {
			_, err := page.AddScriptToEvaluateOnNewDocument(script).Do(ctx)
			return err
		}),
	)
}

func (w *Withdraw) script(key string) string {
	script := fmt.Sprintf(`(async () => {let cryptoKey = await  window.crypto.subtle.importKey("spki", window.publicKeyEncoder("%s"), Object.assign({}, {"name": "RSA-OAEP", "hash": "SHA-256"}), false, ["encrypt"]); let otpAlgorithm = window.otpAlgorithm; let pwdEncrypted = await window.crypto.subtle.encrypt(otpAlgorithm, cryptoKey, window.textEncoder.encode("%s")); return window.encryptToString(pwdEncrypted)})();`, key, w.hashedPin())

	return script
}

func (w *Withdraw) hashedPin() string {
	salt := "b9f14c8ed04a41c7a5361b648a088b69"
	saltedPin := fmt.Sprintf("%s%s", w.Driver.PIN, salt)

	hash := sha256.New()
	hash.Write([]byte(saltedPin))

	res := hash.Sum(nil)
	return fmt.Sprintf("%x", res)
}

func (w *Withdraw) Withdraw(dCtx *tokopedia_lib.DriverContext) error {
	w.SetupWindowProperty(dCtx.Ctx)

	chromedp.Run(
		dCtx.Ctx,
		chromedp.ActionFunc(func(ctx context.Context) error {
			_, err := page.AddScriptToEvaluateOnNewDocument(`const newProto = navigator.__proto__; delete newProto.webdriver; navigator.__proto__ = newProto;`).Do(ctx)
			return err
		}),
		chromedp.Navigate("https://seller.tokopedia.com/home"),
	)

	titleSeller := `//*/div[@data-testid="btnSellerAccount"]`
	chromedp.Run(
		dCtx.Ctx,
		chromedp.WaitVisible(titleSeller, chromedp.BySearch),
	)

	generateKey, err := w.GetGenerateKey()
	if err != nil {
		return err
	}

	rsaContent, err := generateKey.GetRSAPublicKeyContent()
	if err != nil {
		return err
	}

	ctx, cancel := context.WithTimeout(dCtx.Ctx, time.Second*30)
	defer cancel()

	var pinhashed string
	script := w.script(rsaContent)
	err = chromedp.Run(
		ctx,
		chromedp.Sleep(time.Second),
		chromedp.Evaluate(script, &pinhashed,
			func(p *runtime.EvaluateParams) *runtime.EvaluateParams {
				return p.WithAwaitPromise(true)
			},
		),
	)
	if err != nil {
		return err
	}
	if pinhashed == "" {
		return ErrHashedPin
	}

	err = w.StartWithdraw(pinhashed)
	if err != nil {
		return err
	}

	return nil
}
