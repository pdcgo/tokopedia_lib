package tokopedia_lib

import (
	"context"
	"errors"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/chromedp/cdproto/network"
	"github.com/chromedp/chromedp"
	"github.com/pdcgo/common_conf/pdc_common"
	"github.com/pdcgo/tokopedia_lib/lib/api"
	"github.com/sethvargo/go-retry"
)

type DriverContext struct {
	sync.Mutex
	Logined bool
	Ctx     context.Context
}

type DriverSession interface {
	SetCookieToDriver(ctx context.Context) error
	Load() error
	DeleteSession() error
	SaveSession() error
	Sync() error
	Update(cookies []*http.Cookie) error
	AddToHttpRequest(req *http.Request)
	UserAgent() string
	SaveFromDriver(cookies []*network.Cookie, ua string) error
	GetCookies() []*http.Cookie
}

type DriverAccount struct {
	sync.RWMutex
	Username  string          `json:"username"`
	Password  string          `json:"password"`
	Secret    string          `json:"secret"`
	PIN       string          `json:"pin"`
	DevMode   bool            `json:"-"`
	Proxy     string          `json:"-"`
	Session   DriverSession   `json:"-"`
	ParentCtx context.Context `json:"-"`
}

type BrowserClosed struct {
	sync.Mutex
	Data bool
}

func (d *DriverAccount) waitAcceptCookies(ctx context.Context) {
	popuponetrust := `//*/div[@class="ot-sdk-container"]`
	acceptonetrust := `//*/button[@id="onetrust-accept-btn-handler"]`

	wait := func(wctx context.Context) error {
		err := chromedp.Run(wctx,
			chromedp.WaitVisible(popuponetrust, chromedp.BySearch),
			chromedp.WaitVisible(acceptonetrust, chromedp.BySearch),
			chromedp.Sleep(time.Second),
			chromedp.Click(acceptonetrust, chromedp.BySearch),
			chromedp.Sleep(time.Second),
		)

		return err
	}

	go func() {
		// accept pertama
		d.Lock()
		func() {
			defer d.Unlock()

			wctx, cancel := context.WithTimeout(ctx, time.Second*10)
			defer cancel()

			wait(wctx)
		}()

		// accept lagi jika ada
	Parent:
		for {
			select {
			case <-ctx.Done():
				break Parent

			default:
				wait(ctx)
				time.Sleep(time.Second)
			}
		}
	}()
}

func (d *DriverAccount) CreateContext(headless bool) (*DriverContext, func()) {
	opt := []func(*chromedp.ExecAllocator){
		chromedp.Flag("headless", headless),
		chromedp.Flag("incognito", true),
		// chromedp.UserAgent("Mozilla/5.0 (Macintosh; Intel Mac OS X 10_14_5) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/77.0.3830.0 Safari/536.36"),
		// chromedp.Flag("start-maximized", true),
		// chromedp.Flag("disable-blink-features", "AutomationControlled"),

		// chromedp.UserDataDir(pathProfile),
		// chromedp.Flag("profile-directory", "Default"),
	}

	if headless {
		opt = append(opt, chromedp.UserAgent("Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/121.0.0.0 Safari/537.36"))
	}

	if d.DevMode {
		opt = append(opt,
			chromedp.Flag("auto-open-devtools-for-tabs", true),
		)
	}
	if d.Proxy != "" {
		opt = append(opt, chromedp.ProxyServer(d.Proxy))
	}

	parentCtx := context.Background()
	if d.ParentCtx != nil {
		parentCtx = d.ParentCtx
	}

	ctxall, cancelAloc := chromedp.NewExecAllocator(
		parentCtx,
		opt...,
	)

	ctx, cancelCtx := chromedp.NewContext(ctxall)

	dctx := DriverContext{
		Logined: false,
		Ctx:     ctx,
	}

	d.Session.SetCookieToDriver(dctx.Ctx)
	d.waitAcceptCookies(dctx.Ctx)

	return &dctx, func() {

		d.SaveSession(&dctx)
		cancelCtx()
		cancelAloc()

	}
}

func (d *DriverAccount) SaveSession(dctx *DriverContext) error {
	return chromedp.Run(dctx.Ctx,
		chromedp.ActionFunc(func(ctx context.Context) error {
			cookies, err := network.GetCookies().Do(ctx)
			if err != nil {
				return err
			}

			var userAgent string
			err = chromedp.Evaluate("navigator.userAgent", &userAgent).Do(ctx)
			if err != nil {
				return err
			}

			err = d.Session.SaveFromDriver(cookies, userAgent)
			if err != nil {
				return err
			}
			return nil
		}),
	)

}

func (driver *DriverAccount) SetPIN(pin string) {
	driver.PIN = pin
}

func (driver *DriverAccount) MitraLogin(ctx context.Context) error {
	chromedp.Run(ctx, chromedp.Navigate("https://mitra.tokopedia.com"))
	errChan := make(chan error, 1)

	tabakun := `//*/button[@data-testid="tabHomeAkunSaya"]`
	pathlogout := `//*/h4[contains(text(), "Keluar Akun")]`

	go func() {
		driver.RLock()
		defer driver.RUnlock()

		chromedp.Run(ctx,
			chromedp.WaitVisible(tabakun, chromedp.BySearch),
			chromedp.Click(tabakun, chromedp.BySearch),
			chromedp.WaitReady(pathlogout, chromedp.BySearch),
		)
		errChan <- nil
	}()

	go func() {
		driver.RLock()
		defer driver.RUnlock()

		banner := `//*/img[@class="success fade"]`
		chromedp.Run(ctx,

			chromedp.WaitVisible(banner, chromedp.BySearch),
			chromedp.WaitVisible(tabakun, chromedp.BySearch),
			chromedp.Click(tabakun, chromedp.BySearch),
			chromedp.Click(tabakun, chromedp.BySearch),
			chromedp.WaitReady(pathlogout, chromedp.BySearch),
		)
		errChan <- nil
	}()

	go func() {
		driver.RLock()
		defer driver.RUnlock()

		pathemail := `//*/input[@name="login"]`
		selanjutnya := `//*/button[@id="button-submit"]`
		pathpass := `//*/input[@id="login-widget-password"]`
		// masuk := `//*/span[@aria-label="login-button"]`
		pathauthentica := `//*/section[@aria-label="google_authenticator"]`

		chromedp.Run(ctx,
			chromedp.WaitReady(pathemail, chromedp.BySearch),

			chromedp.SendKeys(pathemail, driver.Username, chromedp.BySearch),
			chromedp.Click(selanjutnya, chromedp.BySearch),
			chromedp.WaitVisible(pathpass, chromedp.BySearch),
			chromedp.SendKeys(pathpass, driver.Password, chromedp.BySearch),
			chromedp.Click(selanjutnya, chromedp.BySearch),
			chromedp.WaitVisible(pathauthentica, chromedp.BySearch),
			chromedp.Click(pathauthentica, chromedp.BySearch),
			chromedp.WaitVisible(tabakun, chromedp.BySearch),
		)
		errChan <- nil
	}()

	go func() {
		driver.RLock()
		defer driver.RUnlock()

		pathotp := `//*/input[@autocomplete="one-time-code"]`

		chromedp.Run(ctx,
			chromedp.WaitVisible(pathotp, chromedp.BySearch),
			chromedp.ActionFunc(func(ctx context.Context) error {
				otp, _ := GetTotp(driver.Secret)
				return chromedp.Run(ctx, chromedp.SendKeys(pathotp, otp, chromedp.BySearch))
			}),
		)
	}()

	return <-errChan
}

func (d *DriverAccount) Run(headless bool, actionCallback func(dctx *DriverContext) error) error {
	dctx, cancel := d.CreateContext(headless)
	defer cancel()

	return actionCallback(dctx)

}

func (d *DriverAccount) CreateApi() (*api.TokopediaApi, func(), error) {

	loginBrowser := func() error {
		err := d.Run(false, func(dctx *DriverContext) error {
			return d.SellerLogin(dctx)
		})

		if err != nil {
			return err
		}
		return d.Session.Load()
	}

	var sapi *api.TokopediaApi
	var saveSession = func() {
		log.Println(d.Username, "save session")
		err := d.Session.SaveSession()
		if err != nil {
			pdc_common.ReportError(err)
		}
	}

	b := retry.NewFibonacci(time.Second)
	err := retry.Do(context.Background(), retry.WithMaxRetries(3, b), func(ctx context.Context) error {

		err := d.Session.Load()
		if errors.Is(err, ErrSessionNotFound) {
			errlogin := loginBrowser()
			if errlogin != nil {
				return errlogin
			}
			return retry.RetryableError(err)
		}

		acapi := api.NewTokopediaApi(d.Session)
		sapi = acapi
		_, err = acapi.IsAutheticated()

		if err != nil {

			if errors.Is(err, api.ErrNoShopid) {
				d.Session.DeleteSession()
			}

			errlogin := loginBrowser()
			if errlogin != nil {
				return errlogin
			}
			// return nil
			return retry.RetryableError(err)
		}

		// check is seller account
		data, _ := acapi.ShopInfoByID()
		if data == nil || len(data.Data.ShopInfoByID.Result) == 0 {
			return ErrNotSellerAccount
		}

		return nil
	})

	if err != nil {
		return nil, func() {}, err
	}

	return sapi, saveSession, nil
}

func NewDriverAccount(username string, password string, secret string) (*DriverAccount, error) {
	sess := NewSession(username)
	err := sess.Load()

	if errors.Is(err, ErrSessionNotFound) {
		err = nil
	}

	return &DriverAccount{
		Username: username,
		Password: password,
		Secret:   secret,
		Session:  sess,
	}, err

}
