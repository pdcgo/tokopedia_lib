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
	data *DriverApiData
}

func NewDriverGroup(data *DriverApiData) *DriverGroup {
	return &DriverGroup{
		data: data,
	}
}

func (g *DriverGroup) AddDriverApi(username string, password string, secret string) (dapi *DriverApi, err error) {
	g.Lock()
	defer g.Unlock()

	driver, err := tokopedia_lib.NewDriverAccount(username, password, secret)
	if err != nil {
		return
	}

	acapi, saveSession, err := driver.CreateApi()
	if err != nil {
		return
	}
	defer saveSession()

	dapi = g.data.Add(driver, acapi)
	return
}

func (g *DriverGroup) WithDriverApi(shopid int, handler func(dapi *DriverApi) error) error {
	g.RLock()
	defer g.RUnlock()

	dapi, err := g.data.Get(shopid)
	if err != nil {
		return err
	}

	return handler(dapi)
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

	dapi, err := g.data.Get(shopid)
	if err != nil {
		return func() {}, err
	}

	ctx, cancel := context.WithCancel(context.Background())
	dapi.Driver.ParentCtx = ctx

	_, err = dapi.Api.IsAutheticated()
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
