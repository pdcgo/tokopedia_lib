package tokopedia_lib

import (
	"context"
	"fmt"
	"os"
	"strings"
	"sync"
	"time"

	"github.com/chromedp/chromedp"
	"github.com/pdcgo/common_conf/pdc_common"
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
	Status   string
	Pesan    string
}

// type DriverControl struct {
// 	Context context.Context
// 	ErrCtl  chan<- error
// }

// type UserIdData struct {
// 	Value int32 `json:"value"`
// }

// type UserInfo struct {
// 	UserID int32
// }

type BrowserClosed struct {
	sync.Mutex
	Data bool
}

// func (d *DriverAccount) ProfilePath() string {
// 	pathdata, _ := filepath.Abs("/shopee_profile/" + d.Username)
// 	return pathdata
// }

func (d *DriverAccount) CreateContext(headless bool) (*DriverContext, func()) {
	opt := []func(*chromedp.ExecAllocator){
		chromedp.Flag("headless", headless),
		chromedp.Flag("incognito", true),
		// chromedp.Flag("start-maximized", true),
		// chromedp.Flag("disable-blink-features", "AutomationControlled"),

		// chromedp.UserDataDir(pathProfile),
		// chromedp.Flag("profile-directory", "Default"),
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

		DumpDriverCookies(d.Username, &dctx)
		cancelCtx()
		cancelAloc()

	}
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

func (d *DriverAccount) Run(headless bool, actionCallback func(dctx *DriverContext)) {
	dctx, cancel := d.CreateContext(headless)
	defer cancel()

	actionCallback(dctx)

}

// type ApiCookies struct {
// 	U   *url.URL
// 	Jar http.CookieJar
// }

// func (c *ApiCookies) GetCToken() string {
// 	token := c.FindCookie("CTOKEN")
// 	data, _ := url.QueryUnescape(token)
// 	return data
// }

// func (c *ApiCookies) FindCookie(name string) string {
// 	for _, cookie := range c.Jar.Cookies(c.U) {
// 		if cookie.Name == name {
// 			return cookie.Value
// 		}
// 	}
// 	return ""
// }

// func GetCookies(ctx context.Context) *ApiCookies {
// 	var hasil []*http.Cookie

// 	chromedp.Run(ctx,
// 		chromedp.ActionFunc(func(ctx context.Context) error {
// 			cook := network.GetCookies()
// 			cookies, _ := cook.Do(ctx)
// 			for _, co := range cookies {
// 				cookie := http.Cookie{
// 					Name:   co.Name,
// 					Value:  co.Value,
// 					Path:   co.Path,
// 					Domain: co.Domain,
// 					Secure: co.Secure,
// 				}
// 				hasil = append(hasil, &cookie)
// 			}

// 			return nil
// 		}),
// 	)

// 	jar, _ := cookiejar.New(nil)
// 	u, _ := url.Parse("https://seller.shopee.co.id")
// 	jar.SetCookies(u, hasil)

// 	return &ApiCookies{
// 		U:   u,
// 		Jar: jar,
// 	}
// }

//

// type LoginResult struct {
// 	Loginned bool
// 	F02      bool
// }

func NewDriverAccount(username string, password string, secret string, status string, pesan string) (*DriverAccount, error) {

	return &DriverAccount{
		Username: username,
		Password: password,
		Secret:   secret,
		Status:   status,
		Pesan:    pesan,
	}, nil

}

func GetakunFromFile(fname string) ([]*DriverAccount, func(), error) {
	hasil := []*DriverAccount{}
	data, _ := os.ReadFile(fname)
	lines := strings.Split(string(data), "\n")

Parent:
	for _, line := range lines {
		if line == "" {
			continue Parent
		}

		dataline := make([]string, 5)

		fixline := strings.ReplaceAll(line, "\r", "")

		for ind, value := range strings.Split(fixline, "|") {
			dataline[ind] = value
		}

		driver, err := NewDriverAccount(dataline[0], dataline[1], dataline[2], dataline[3], dataline[4])
		if err != nil {
			pdc_common.ReportError(err)
		}

		hasil = append(hasil, driver)
	}

	return hasil, func() {
		f, err := os.OpenFile(fname, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0755)
		if err != nil {
			panic(err)
		}
		defer f.Close()
		for _, driver := range hasil {
			f.WriteString(fmt.Sprintf("%s|%s|%s|%s|%s\n", driver.Username, driver.Password, driver.Secret, driver.Status, driver.Pesan))
		}

	}, nil

}
