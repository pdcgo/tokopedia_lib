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

type DriverApi struct {
	Api    *api.TokopediaApi
	Driver *tokopedia_lib.DriverAccount
}

type DriverGroup struct {
	sync.RWMutex
	driverLock  sync.Mutex
	data        map[string]*DriverApi
	usernamemap map[int]string
}

func NewDriverGroup() *DriverGroup {
	return &DriverGroup{
		driverLock:  sync.Mutex{},
		data:        map[string]*DriverApi{},
		usernamemap: map[int]string{},
	}
}

func (g *DriverGroup) AddDriverApi(username string, password string, secret string) error {
	g.driverLock.Lock()
	defer g.driverLock.Unlock()

	driver, err := tokopedia_lib.NewDriverAccount(username, password, secret)
	if err != nil {
		return err
	}

	acapi, saveSession, err := driver.CreateApi()
	if err != nil {
		return err
	}
	defer saveSession()

	shopid := acapi.AuthenticatedData.UserShopInfo.Info.ShopID
	g.usernamemap[int(shopid)] = username
	g.data[username] = &DriverApi{
		Api:    acapi,
		Driver: driver,
	}
	return nil
}

var ErrNoDriver = errors.New("driver not found")

type DriverApiHandler func(driver *tokopedia_lib.DriverAccount, api *api.TokopediaApi) error

func (g *DriverGroup) WithDriverApi(username string, handler func(dapi *DriverApi) error) (err error) {
	g.RLock()
	defer g.RUnlock()

	dapi := g.data[username]
	if dapi == nil {
		return ErrNoDriver
	}

	return handler(dapi)
}

func (g *DriverGroup) WithDriverApiByShopid(shopid int, handler func(username string, dapi *DriverApi) error) error {
	username := g.usernamemap[shopid]
	return g.WithDriverApi(username, func(dapi *DriverApi) error {
		return handler(username, dapi)
	})
}

func (g *DriverGroup) Reset() {
	g.Lock()
	defer g.Unlock()

	g.data = map[string]*DriverApi{}
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

func (g *DriverGroup) OpenDriver(shopid int) (context.CancelFunc, error) {
	g.RLock()
	defer g.RUnlock()

	username := g.usernamemap[shopid]
	dapi := g.data[username]
	if dapi == nil {
		return func() {}, ErrNoDriver
	}

	ctx, cancel := context.WithCancel(context.Background())
	dapi.Driver.ParentCtx = ctx

	_, err := dapi.Api.IsAutheticated()
	if errors.Is(err, api.ErrNoShopid) {
		dapi.Driver.Session.DeleteSession()
	}

	_, err = dapi.Api.ShopInfoByID()
	if errors.Is(err, api.ErrIsNotAuthorized) {
		dapi.Driver.Session.DeleteSession()
	}

	saldoSuccess := g.reqSaldoSuccess(dapi.Driver.Session)
	if !saldoSuccess {
		dapi.Driver.Session.DeleteSession()
	}

	go dapi.Driver.Run(false, func(dctx *tokopedia_lib.DriverContext) error {
		dapi.Driver.SellerLogin(dctx)
		<-ctx.Done()
		return nil
	})

	return cancel, nil
}
