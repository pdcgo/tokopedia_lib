package grabber

import (
	"context"
	"net/url"
	"strings"
	"sync"

	"github.com/pdcgo/common_conf/pdc_common"
	"github.com/pdcgo/go_v2_shopeelib/app/upload_app/legacy_source"
	"github.com/pdcgo/go_v2_shopeelib/lib/legacy"
	"github.com/pdcgo/tokopedia_lib/lib/api_public"
	"github.com/pdcgo/tokopedia_lib/lib/grab_handler"
	"github.com/pdcgo/tokopedia_lib/lib/model_public"
)

type Grabber interface {
	Run() error
}

type BaseGrabber struct {
	Api          *api_public.TokopediaApiPublic
	Base         *legacy_source.BaseConfig
	GrabTasker   *legacy.GrabTasker
	CacheHandler *grab_handler.CacheProductHandler

	wg         sync.WaitGroup
	limitGuard chan int
}

func NewBaseGrabber(
	api *api_public.TokopediaApiPublic,
	base *legacy_source.BaseConfig,
	tasker *legacy.GrabTasker,
	cacheHandler *grab_handler.CacheProductHandler,
) *BaseGrabber {

	grabBasic := legacy.NewGrabBasic(base)

	return &BaseGrabber{
		Api:          api,
		Base:         base,
		CacheHandler: cacheHandler,
		GrabTasker:   tasker,
		limitGuard:   make(chan int, grabBasic.Concurrent),
	}
}

func (g *BaseGrabber) GetProductLayout(ctx context.Context, url string) *model_public.PdpGetlayoutQueryResp {
	select {
	case <-ctx.Done():
		return nil
	default:
		layoutVar, _ := ParseProductDetailParamsFromUrl(url)
		layout, err := g.Api.PdpGetlayoutQuery(layoutVar)
		if err != nil {
			pdc_common.ReportError(err)
			return nil
		}
		return layout
	}
}

func (g *BaseGrabber) GetPdpDataP2(ctx context.Context, pdpSession string, prodId string) *model_public.PdpGetDataP2Resp {
	select {
	case <-ctx.Done():
		return nil
	default:
		pdpVar := &model_public.PdpGetDataP2Var{
			PdpSession: pdpSession,
			ProductID:  prodId,
		}
		pdp, err := g.Api.PdpGetDataP2(pdpVar)
		if err != nil {
			pdc_common.ReportError(err)
			return nil
		}
		return pdp
	}
}

func ParseProductDetailParamsFromUrl(uri string) (*model_public.PdpGetlayoutQueryVar, error) {
	u, err := url.Parse(uri)
	if err != nil {
		return nil, err
	}
	path := u.EscapedPath()
	query := u.Query()

	splitPath := strings.Split(path, "/")
	shopDomain := splitPath[len(splitPath)-2]
	productKey := splitPath[len(splitPath)-1]

	payload := &model_public.PdpGetlayoutQueryVar{
		ShopDomain: shopDomain,
		ProductKey: productKey,
		APIVersion: 1,
		ExtParam:   url.QueryEscape(query.Get("extParam")),
	}
	return payload, nil
}

func GenerateShopProductVar() *model_public.ShopProductVar {
	params := &model_public.ShopProductVar{
		Page:           1,
		PerPage:        100,
		EtalaseID:      "etalase",
		Sort:           1,
		Sid:            "",
		UserDistrictID: "176",
		UserCityID:     "2274",
		UserLat:        "",
		UserLong:       "",
	}
	return params
}
