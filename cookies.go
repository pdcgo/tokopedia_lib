package tokopedia_lib

import (
	"context"
	"encoding/json"
	"errors"
	"os"
	"path/filepath"
	"time"

	"github.com/chromedp/cdproto/cdp"
	"github.com/chromedp/cdproto/network"
	"github.com/chromedp/chromedp"
	"github.com/pdcgo/common_conf/pdc_common"
)

// TODO: cookies dump afternya tidak ada
var BaseSessionPath = "/tokopedia_session_new/"

func init() {
	pathdata, _ := filepath.Abs(BaseSessionPath)
	if _, err := os.Stat(pathdata); errors.Is(err, os.ErrNotExist) {
		os.MkdirAll(pathdata, os.ModeDir)
	}
}

func GetDriverCacheCookies(username string) []*network.Cookie {
	var cookies []*network.Cookie
	pathdata := filepath.Join(BaseSessionPath, username+".json")

	if _, err := os.Stat(pathdata); errors.Is(err, os.ErrNotExist) {
		return cookies
	}

	file, err := os.ReadFile(pathdata)
	if err != nil {
		pdc_common.ReportError(err)
	}

	err = json.Unmarshal(file, &cookies)
	if err != nil {
		pdc_common.ReportError(err)
	}

	return cookies
}

func SetDriverCookies(username string, dcontext *DriverContext) {

	chromedp.Run(
		dcontext.Ctx,
		chromedp.ActionFunc(func(ctx context.Context) error {
			cookies := GetDriverCacheCookies(username)

			for _, cookie := range cookies {
				expr := cdp.TimeSinceEpoch(time.Now().Add(180 * 24 * time.Hour))

				err := network.SetCookie(cookie.Name, cookie.Value).
					WithDomain(cookie.Domain).
					WithPath(cookie.Path).
					WithHTTPOnly(cookie.HTTPOnly).
					WithSecure(cookie.Secure).
					WithExpires(&expr).
					Do(ctx)

				if err != nil {
					if !errors.Is(context.Canceled, err) {
						pdc_common.ReportError(err)
					}

				}
			}
			return nil
		}),
	)
}

func (d *DriverAccount) SaveCookiesBrowser(ctx context.Context) error {
	return chromedp.Run(ctx,
		chromedp.ActionFunc(func(ctx context.Context) error {
			cookies, err := network.GetCookies().Do(ctx)
			if err != nil {
				pdc_common.ReportError(err)
				return err
			}

			pathdata := filepath.Join(BaseSessionPath, d.Username+".json")
			file, err := os.OpenFile(pathdata, os.O_RDWR|os.O_CREATE|os.O_TRUNC, os.ModePerm)
			if err != nil {
				return pdc_common.ReportError(err)
			}
			defer file.Close()

			err = json.NewEncoder(file).Encode(cookies)
			if err != nil {
				return pdc_common.ReportError(err)
			}

			return nil
		}),
	)
}
