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
)

var ErrSaldoKosong = errors.New("saldo kosong")
var ErrHashedPin = errors.New("error hashing pin")

type Withdraw struct {
	Api *api.TokopediaApi
}

func NewWithdraw(api *api.TokopediaApi) *Withdraw {
	w := &Withdraw{
		Api: api,
	}

	return w
}

func (w *Withdraw) GetGenerateKey(phone string) (*api.WindrawnGenerateKeyResp, error) {
	_, err := w.Api.WithdrawOtpRequest(phone)
	if err != nil {
		return nil, err
	}

	generateKey, err := w.Api.WindrawnGenerateKey()
	if err != nil {
		return nil, err
	}

	return generateKey, nil
}

func (w *Withdraw) StartWithdraw(phone, pinHashed, h, amount string) error {

	banks, err := w.Api.BankListQuery(false)
	if err != nil {
		return err
	}

	bank := banks.Data.GetBankListWDV2.Data.GetDefaultBank()

	otpValidateVariable := api.NewOtpValidateVariable(phone, strconv.Itoa(bank.BankAccountID), pinHashed, h)
	otpValidate, err := w.Api.WithdrawOtpValidate(otpValidateVariable)
	if err != nil {
		return err
	}

	withdrawvariable := api.NewWithdrawVariable(bank, otpValidate.Data.OTPValidate, amount)
	_, err = w.Api.WithdrawSaldoMutation(withdrawvariable)

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

func (w *Withdraw) GetHashedPIN(ctx context.Context, pin string, key string) (string, error) {
	salt := "b9f14c8ed04a41c7a5361b648a088b69"
	saltedPin := fmt.Sprintf("%s%s", pin, salt)

	hash := sha256.New()
	hash.Write([]byte(saltedPin))

	res := hash.Sum(nil)
	hashedPin := fmt.Sprintf("%x", res)

	script := fmt.Sprintf(`(async () => {let cryptoKey = await  window.crypto.subtle.importKey("spki", window.publicKeyEncoder("%s"), Object.assign({}, {"name": "RSA-OAEP", "hash": "SHA-256"}), false, ["encrypt"]); let otpAlgorithm = window.otpAlgorithm; let pwdEncrypted = await window.crypto.subtle.encrypt(otpAlgorithm, cryptoKey, window.textEncoder.encode("%s")); return window.encryptToString(pwdEncrypted)})();`, key, hashedPin)

	w.SetupWindowProperty(ctx)

	chromedp.Run(
		ctx,
		chromedp.Navigate("https://seller.tokopedia.com/home"),
	)

	titleSeller := `//*/div[@data-testid="btnSellerAccount"]`
	chromedp.Run(
		ctx,
		chromedp.WaitVisible(titleSeller, chromedp.BySearch),
	)

	ctx, cancel := context.WithTimeout(ctx, time.Second*30)
	defer cancel()

	var pinhashed string
	err := chromedp.Run(
		ctx,
		chromedp.Sleep(time.Second),
		chromedp.Evaluate(script, &pinhashed,
			func(p *runtime.EvaluateParams) *runtime.EvaluateParams {
				return p.WithAwaitPromise(true)
			},
		),
	)
	if err != nil {
		return "", err
	}
	if pinhashed == "" {
		return "", ErrHashedPin
	}

	return pinhashed, nil
}

func (w *Withdraw) Run(dCtx *tokopedia_lib.DriverContext, pin string, report *WithdrawReport) error {
	user, err := w.Api.UserDataQuery()
	if err != nil {
		return err
	}

	balance, err := w.Api.WithDrawBalance()
	if err != nil {
		return err
	}

	report.Jumlah = balance.Data.MidasGetAllDepositAmount.SellerAllFmt

	if balance.Data.MidasGetAllDepositAmount.SellerAll == 0 {
		return ErrSaldoKosong
	}

	generateKey, err := w.GetGenerateKey(user.Data.User.Phone)
	if err != nil {
		return err
	}

	rsaContent, err := generateKey.GetRSAPublicKeyContent()
	if err != nil {
		return err
	}

	pinhashed, err := w.GetHashedPIN(dCtx.Ctx, pin, rsaContent)
	if err != nil {
		return err
	}

	err = w.StartWithdraw(user.Data.User.Phone, pinhashed, generateKey.Data.GenerateKey.H, strconv.Itoa(balance.Data.MidasGetAllDepositAmount.SellerAll))
	if err != nil {
		return err
	}

	return nil
}
