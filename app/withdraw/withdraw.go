package withdraw

import (
	"context"
	"errors"
	"strconv"
	"time"

	"github.com/chromedp/cdproto/network"
	"github.com/chromedp/cdproto/page"
	"github.com/chromedp/chromedp"
	"github.com/pdcgo/tokopedia_lib"
	"github.com/pdcgo/tokopedia_lib/lib/api"
	"github.com/sethvargo/go-retry"
)

type Withdraw struct {
	*tokopedia_lib.DriverAccount
}

func NewWithdraw(driver *tokopedia_lib.DriverAccount) *Withdraw {
	return &Withdraw{
		DriverAccount: driver,
	}
}

func (w *Withdraw) RunWithApi() error {
	tApi, saveSession, err := w.CreateApi()
	if err != nil {
		return err
	}
	defer saveSession()

	user, err := tApi.UserDataQuery()
	if err != nil {
		return err
	}

	balance, err := tApi.GetBalance()
	if err != nil {
		return err
	}

	banks, err := tApi.BankListQuery(false)
	if err != nil {
		return err
	}

	bank := banks.Data.GetBankListWDV2.Data.GetDefaultBank()

	_, err = tApi.WithdrawOtpRequest(user.Data.User.Phone)
	if err != nil {
		return err
	}

	generateKey, err := tApi.WindrawnGenerateKey()
	if err != nil {
		return err
	}

	otpValidateVariable, err := api.NewOtpValidateVariable(user.Data.User.Phone, strconv.Itoa(bank.BankAccountID), w.PIN, generateKey.Data.GenerateKey)
	if err != nil {
		return err
	}
	otpValidate, err := tApi.WithdrawOtpValidate(otpValidateVariable)
	if err != nil {
		return err
	}

	withdrawvariable := api.NewWithdrawVariable(user.Data.User, bank, otpValidate.Data.OTPValidate, strconv.Itoa(balance.Data.Balance.SellerAll))
	withdrawSaldo, err := tApi.WithdrawSaldoMutation(withdrawvariable)
	if err != nil {
		return err
	}

	if withdrawSaldo.Data.RichieSubmitWithdrawal.Status != "success" {
		return errors.New(withdrawSaldo.Data.RichieSubmitWithdrawal.MessageError)
	}

	return nil
}

func (w *Withdraw) RunWithDriver() error {

	withdraw := func() error {
		err := w.Run(false, func(dctx *tokopedia_lib.DriverContext) error {
			return w.Withdraw(dctx)
		})
		return err
	}

	b := retry.NewFibonacci(time.Second)
	err := retry.Do(context.Background(), retry.WithMaxRetries(3, b), func(ctx context.Context) error {
		err := withdraw()
		if err != nil {
			if err == ErrSaldoKosong {
				return err
			}

			return retry.RetryableError(err)
		}

		return nil
	})

	return err
}

var ErrSaldoKosong = errors.New("saldo kosong")
var ErrWithdraw = errors.New("withdraw error")

func (w *Withdraw) Withdraw(dCtx *tokopedia_lib.DriverContext) error {
	errorChan := make(chan error, 1)

	chromedp.Run(
		dCtx.Ctx,
		chromedp.ActionFunc(func(ctx context.Context) error {
			_, err := page.AddScriptToEvaluateOnNewDocument(`const newProto = navigator.__proto__;
			delete newProto.webdriver;
			navigator.__proto__ = newProto;`).Do(ctx)
			return err
		}),
		chromedp.Navigate("https://seller.tokopedia.com/"),
	)

	loginCtx, cancelLoginCtx := context.WithCancel(dCtx.Ctx)

	go func() {
		defer cancelLoginCtx()

		titleSeller := `//*/div[@data-testid="btnSellerAccount"]`
		chromedp.Run(
			dCtx.Ctx,
			chromedp.WaitVisible(titleSeller, chromedp.BySearch),
		)

		ctx, cancel := context.WithTimeout(dCtx.Ctx, time.Minute)
		defer cancel()

		saldo := ""
		chromedp.Run(
			ctx,
			chromedp.WaitEnabled("//*/span[@data-testid='txtSellerSidebarValueSaldo']"),
			chromedp.TextContent("//*/span[@data-testid='txtSellerSidebarValueSaldo']", &saldo, chromedp.BySearch),
		)
		if saldo == "Rp0" {
			errorChan <- ErrSaldoKosong
			return
		}

		withdrawUri := "https://www.tokopedia.com/payment/deposit?nref=dside"
		creditsUri := "https://ta.tokopedia.com/v2/manage/credits"
		chromedp.Run(
			dCtx.Ctx,
			chromedp.Sleep(time.Second),
			chromedp.Navigate(creditsUri),
			chromedp.Sleep(time.Second),
			chromedp.Navigate(withdrawUri),
		)

		go func() {
			for {
				chromedp.Run(
					ctx,
					chromedp.WaitEnabled("//*/button/span[contains(text(), 'Tarik Saldo')]/..", chromedp.BySearch),
					chromedp.Click("//*/button/span[contains(text(), 'Tarik Saldo')]/..", chromedp.BySearch),
				)
			}
		}()

		go func() {
			for {
				chromedp.Run(
					ctx,
					chromedp.WaitEnabled(".unf-coachmark__next-button", chromedp.ByQuery),
					chromedp.Click(".unf-coachmark__next-button", chromedp.ByQuery),
				)
			}
		}()

		wdSteps := func() error {
			ctx, cancel := context.WithTimeout(dCtx.Ctx, time.Minute)
			defer cancel()

			tarikSaldoBtn := "//*/span[contains(text(), 'Tarik Saldo')]"
			penghasilanBtn := "#unf-tabitem-coachmark1-1"
			withdrawAllBtn := "//*/div/span[@data-testid='wd-withdraw-all']"
			chromedp.Run(
				ctx,
				chromedp.WaitVisible(penghasilanBtn, chromedp.ByID),
				chromedp.Click(penghasilanBtn, chromedp.ByID),
				chromedp.WaitVisible(withdrawAllBtn, chromedp.BySearch),
				chromedp.Click(withdrawAllBtn, chromedp.BySearch),
				chromedp.Sleep(time.Second),
				chromedp.Click(withdrawAllBtn, chromedp.BySearch),
				chromedp.Sleep(time.Second),
				chromedp.WaitEnabled(tarikSaldoBtn+"/../..", chromedp.BySearch),
				chromedp.Click(tarikSaldoBtn+"/../..", chromedp.BySearch),
			)

			ctx, cancel = context.WithTimeout(dCtx.Ctx, time.Second*20)
			defer cancel()

			pinInput := "//*/input[@aria-label='pin input']"
			backToSaldoBtn := "//*/button[@data-testid='wd-btn-back-to-deposit']"
			return chromedp.Run(
				ctx,
				chromedp.WaitVisible(pinInput, chromedp.BySearch),
				chromedp.Sleep(time.Second),
				chromedp.SendKeys(pinInput, w.PIN, chromedp.BySearch),
				chromedp.WaitEnabled(backToSaldoBtn, chromedp.BySearch),
			)
		}

		err := wdSteps()
		if err != nil {
			chromedp.Run(
				dCtx.Ctx,
				chromedp.WaitVisible("#unf-tabitem-coachmark1-0", chromedp.ByID),
				chromedp.Click("#unf-tabitem-coachmark1-0", chromedp.ByID),
				chromedp.Sleep(time.Second),
				chromedp.Click("#unf-tabitem-coachmark1-0", chromedp.ByID),
				chromedp.Sleep(time.Second),
			)
			err := wdSteps()
			errorChan <- err
			return
		}

		errorChan <- err
	}()

	go func() {
		masukTitle := `//*/h3[contains(text(), "Masuk")]`
		chromedp.Run(
			loginCtx,
			chromedp.WaitReady(masukTitle, chromedp.BySearch),
		)
		err := w.MitraLogin(loginCtx)
		if err != nil {
			errorChan <- err
		}

		chromedp.Run(dCtx.Ctx,
			chromedp.Navigate("https://seller.tokopedia.com/"),
			chromedp.ActionFunc(func(ctx context.Context) error {
				cookies, err := network.GetCookies().Do(ctx)
				if err != nil {
					errorChan <- err
					return err
				}

				var userAgent string
				err = chromedp.Evaluate("navigator.userAgent", &userAgent).Do(ctx)
				if err != nil {
					errorChan <- err
					return err
				}

				err = w.Session.SaveFromDriver(cookies, userAgent)
				if err != nil {
					errorChan <- err
				}
				dCtx.Logined = true
				cancelLoginCtx()
				return nil
			}),
		)
	}()

	select {
	case <-dCtx.Ctx.Done():
		return context.Canceled
	case err := <-errorChan:
		return err
	}
}
