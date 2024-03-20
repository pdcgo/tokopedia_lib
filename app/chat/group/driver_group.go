package group

import (
	"context"
	"errors"
	"net/http"
	"sync"

	"github.com/pdcgo/common_conf/pdc_common"
	"github.com/pdcgo/tokopedia_lib"
	"github.com/pdcgo/tokopedia_lib/lib/api"
)

type DriverGroup struct {
	sync.RWMutex
	data     map[string]*tokopedia_lib.DriverAccount
	apicache map[string]*api.TokopediaApi
}

func NewDriverGroup() *DriverGroup {
	return &DriverGroup{
		data:     map[string]*tokopedia_lib.DriverAccount{},
		apicache: map[string]*api.TokopediaApi{},
	}
}

func (g *DriverGroup) AddDriver(username string, password string, secret string) error {
	g.Lock()
	defer g.Unlock()

	driver, err := tokopedia_lib.NewDriverAccount(username, password, secret)
	if err != nil {
		return err
	}

	g.data[username] = driver
	return nil
}

var ErrNoDriver = errors.New("driver not found")

type DriverApiHandler func(driver *tokopedia_lib.DriverAccount, api *api.TokopediaApi) error

func (g *DriverGroup) WithDriverApi(username string, handler DriverApiHandler) (err error) {
	g.RLock()
	defer g.RUnlock()

	driver := g.data[username]
	if driver == nil {
		return ErrNoDriver
	}

	acapi := g.apicache[username]
	if acapi == nil {
		acapi, _, err = driver.CreateApi()
		if err != nil {
			return err
		}

		g.apicache[username] = acapi
	}

	err = handler(driver, acapi)
	return err
}

func (g *DriverGroup) Reset() {
	g.Lock()
	defer g.Unlock()

	g.data = map[string]*tokopedia_lib.DriverAccount{}
}

func (g *DriverGroup) reqSaldoSuccess(session tokopedia_lib.DriverSession) (success bool) {
	url := "https://www.tokopedia.com/payment/deposit?nref=dside"
	req, err := http.NewRequest(http.MethodGet, url, nil)

	req.Header.Add("User-Agent", session.UserAgent())
	session.AddToHttpRequest(req)

	if err != nil {
		pdc_common.ReportError(err)
		return
	}

	res, err := api.ClientApi.Do(req)
	if err != nil {
		pdc_common.ReportError(err)
		return
	}

	if res.StatusCode != 200 {
		return
	}

	success = true
	return
}

func (g *DriverGroup) OpenDriver(username string) (context.CancelFunc, error) {
	g.RLock()
	defer g.RUnlock()

	driver := g.data[username]
	if driver == nil {
		return func() {}, ErrNoDriver
	}

	acapi, saveSession, err := driver.CreateApi()
	if err != nil {
		return func() {}, err
	}
	defer saveSession()

	ctx, cancel := context.WithCancel(context.Background())
	driver.ParentCtx = ctx

	_, err = acapi.IsAutheticated()
	if errors.Is(err, api.ErrNoShopid) {
		driver.Session.DeleteSession()
	}

	_, err = acapi.ShopInfoByID()
	if errors.Is(err, api.ErrIsNotAuthorized) {
		driver.Session.DeleteSession()
	}

	saldoSuccess := g.reqSaldoSuccess(driver.Session)
	if !saldoSuccess {
		driver.Session.DeleteSession()
	}

	go driver.Run(false, func(dctx *tokopedia_lib.DriverContext) error {
		driver.SellerLogin(dctx)
		<-ctx.Done()
		return nil
	})

	return cancel, nil
}
