package tokopedia_lib

import (
	"context"
	"errors"
	"log"

	"github.com/chromedp/cdproto/network"
	"github.com/chromedp/chromedp"
)

var ErrSellerLoginFailed = errors.New("seller login failed")
var ErrNotSellerAccount = errors.New("not seller account")

func (d *DriverAccount) SellerLogin(dctx *DriverContext) error {

	errorChan := make(chan error, 1)

	loginCtx, cancelLoginCtx := context.WithCancel(dctx.Ctx)

	chromedp.Run(dctx.Ctx,
		chromedp.Navigate("https://seller.tokopedia.com/"),
	)

	go func() {
		titleSeller := `//*/div[@data-testid="btnSellerAccount"]`
		chromedp.Run(dctx.Ctx,
			chromedp.WaitVisible(titleSeller, chromedp.BySearch),
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

				err = d.Session.SaveFromDriver(cookies, userAgent)
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
		err := d.MitraLogin(loginCtx)
		if err != nil {
			errorChan <- err
		}

		d.Lock()
		defer d.Unlock()

		chromedp.Run(dctx.Ctx,
			chromedp.Navigate("https://seller.tokopedia.com/"),
		)
	}()

	// check not seller
	go func() {
		submitSeller := `//*/button[@data-testid="btnSubmitShopDomain"]`
		chromedp.Run(dctx.Ctx,
			chromedp.WaitVisible(submitSeller, chromedp.BySearch),
			chromedp.ActionFunc(func(ctx context.Context) error {
				errorChan <- ErrNotSellerAccount
				return nil
			}),
		)
	}()

	select {
	case <-dctx.Ctx.Done():
		return context.Canceled
	case err := <-errorChan:
		if err == nil {
			log.Println(d.Username, "login Berhasil")
		}
		return err
	}

}
