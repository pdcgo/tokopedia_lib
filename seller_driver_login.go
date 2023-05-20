package tokopedia_lib

import (
	"context"
	"errors"

	"github.com/chromedp/chromedp"
)

var ErrSellerLoginFailed = errors.New("seller login failed")

func (d *DriverAccount) SellerLogin(dctx *DriverContext) error {
	// cek, err := d.ExecLogin(dctx)

	errorChan := make(chan error, 1)

	// if !cek {
	// 	return ErrSellerLoginFailed
	// }

	// if err != nil {
	// 	return err
	// }
	loginCtx, cancelLoginCtx := context.WithCancel(dctx.Ctx)

	chromedp.Run(dctx.Ctx,
		chromedp.Navigate("https://seller.tokopedia.com/"),
	)

	go func() {
		titleSeller := `//*/div[@data-testid="btnSellerAccount"]`
		chromedp.Run(dctx.Ctx,
			chromedp.WaitVisible(titleSeller, chromedp.BySearch),
			chromedp.ActionFunc(func(ctx context.Context) error {

				err := d.SaveCookiesBrowser(ctx)
				if err != nil {
					errorChan <- err
				}

				dctx.Logined = true
				cancelLoginCtx()
				errorChan <- nil
				return nil
			}),
		)
	}()

	go func() {
		masukTitle := `//*/h3[contains(text(), "Masuk")]`
		chromedp.Run(loginCtx,
			chromedp.WaitReady(masukTitle, chromedp.BySearch),
		)
		err := d.mitraLogin(loginCtx)
		if err != nil {
			errorChan <- err
		}

		chromedp.Run(dctx.Ctx,
			chromedp.Navigate("https://seller.tokopedia.com/"),
		)
	}()

	return <-errorChan

}
