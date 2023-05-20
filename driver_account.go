package tokopedia_lib

import (
	"context"
	"sync"
	"time"

	"github.com/chromedp/chromedp"
)

type DriverContext struct {
	sync.Mutex
	Logined bool
	Ctx     context.Context
}

type DriverAccount struct {
	Username string
	Password string
	Secret   string
	DevMode  bool
	Proxy    string
}

type BrowserClosed struct {
	sync.Mutex
	Data bool
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

	if d.DevMode {
		opt = append(opt,
			chromedp.Flag("auto-open-devtools-for-tabs", true),
		)
	}
	if d.Proxy != "" {
		opt = append(opt, chromedp.ProxyServer(d.Proxy))
	}

	parentCtx := context.Background()

	ctxall, cancelAloc := chromedp.NewExecAllocator(
		parentCtx,
		opt...,
	)

	ctx, cancelCtx := chromedp.NewContext(ctxall)

	dctx := DriverContext{
		Logined: false,
		Ctx:     ctx,
	}
	SetDriverCookies(d.Username, &dctx)

	// checking jaga2 jika close manual browser nya
	isClosed := BrowserClosed{
		Data: false,
	}
	go func() {
		<-ctx.Done()

		isClosed.Lock()
		defer isClosed.Unlock()

		isClosed.Data = true
	}()

	return &dctx, func() {
		isClosed.Lock()
		defer isClosed.Unlock()

		if isClosed.Data {
			return
		}

		cancelCtx()
		cancelAloc()

	}
}

func (driver *DriverAccount) mitraLogin(ctx context.Context) error {
	chromedp.Run(ctx, chromedp.Navigate("https://mitra.tokopedia.com"))
	errChan := make(chan error, 1)

	go func() {
		pathemail := `//*/input[@name="login"]`
		selanjutnya := `//*/button[@id="button-submit"]`
		pathpass := `//*/input[@id="login-widget-password"]`
		// masuk := `//*/span[@aria-label="login-button"]`
		pathauthentica := `//*/div[@aria-label="google_authenticator"]`
		tabakun := `//*/div[@data-testid="tabHomeAkunSaya"]`

		chromedp.Run(ctx,
			chromedp.WaitVisible(tabakun, chromedp.BySearch),
			chromedp.Click(tabakun, chromedp.BySearch),
			chromedp.WaitVisible(pathemail, chromedp.BySearch),
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

func (driver *DriverAccount) ExecLogin(dctx *DriverContext) (bool, error) {
	cCtx, cancel := context.WithTimeout(dctx.Ctx, time.Minute*3)
	defer cancel()

	waitdata := make(chan int, 1)
	logined := false

	SetDriverCookies(driver.Username, dctx)

	chromedp.Run(cCtx, chromedp.Navigate("https://mitra.tokopedia.com"))

	go func() {
		pathemail := `//*/input[@name="login"]`
		selanjutnya := `//*/button[@id="button-submit"]`
		pathpass := `//*/input[@id="login-widget-password"]`
		// masuk := `//*/span[@aria-label="login-button"]`
		pathauthentica := `//*/div[@aria-label="google_authenticator"]`
		tabakun := `//*/div[@data-testid="tabHomeAkunSaya"]`

		chromedp.Run(cCtx,
			chromedp.WaitVisible(tabakun, chromedp.BySearch),
			chromedp.Click(tabakun, chromedp.BySearch),
			chromedp.WaitVisible(pathemail, chromedp.BySearch),
			chromedp.SendKeys(pathemail, driver.Username, chromedp.BySearch),
			chromedp.Click(selanjutnya, chromedp.BySearch),
			chromedp.WaitVisible(pathpass, chromedp.BySearch),
			chromedp.SendKeys(pathpass, driver.Password, chromedp.BySearch),
			chromedp.Click(selanjutnya, chromedp.BySearch),
			chromedp.WaitVisible(pathauthentica, chromedp.BySearch),
			chromedp.Click(pathauthentica, chromedp.BySearch),
			chromedp.WaitVisible(tabakun, chromedp.BySearch),
		)
		logined = true
		waitdata <- 1
	}()

	go func() {
		pathotp := `//*/input[@autocomplete="one-time-code"]`

		chromedp.Run(cCtx,
			chromedp.WaitVisible(pathotp, chromedp.BySearch),
			chromedp.ActionFunc(func(ctx context.Context) error {
				otp, _ := GetTotp(driver.Secret)
				return chromedp.Run(ctx, chromedp.SendKeys(pathotp, otp, chromedp.BySearch))
			}),
		)
	}()

	// go func() {
	// 	sidebar := `//*/div[@data-testid="imgSellerSidebarProfile"]`
	// 	chromedp.Run(cCtx, chromedp.WaitVisible(sidebar, chromedp.BySearch))
	// 	waitdata <- 1
	// }()

	select {
	case <-cCtx.Done():
		break
	case <-waitdata:
		logined = true
	}

	return logined, nil
}

func (d *DriverAccount) Run(headless bool, actionCallback func(dctx *DriverContext) error) error {
	dctx, cancel := d.CreateContext(headless)
	defer cancel()

	return actionCallback(dctx)

}

func NewDriverAccount(username string, password string, secret string) (*DriverAccount, error) {

	return &DriverAccount{
		Username: username,
		Password: password,
		Secret:   secret,
	}, nil

}
