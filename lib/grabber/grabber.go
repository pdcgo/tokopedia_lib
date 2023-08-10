package grabber

import (
	"context"
	"errors"
	"log"
	"sync"
	"sync/atomic"

	"github.com/pdcgo/common_conf/pdc_common"
	"github.com/pdcgo/go_v2_shopeelib/app/upload_app/legacy_source"
	"github.com/pdcgo/go_v2_shopeelib/lib/legacy"
	"github.com/pdcgo/tokopedia_lib/lib/api_public"
	"github.com/pdcgo/tokopedia_lib/lib/grab_handler"
	"github.com/pdcgo/tokopedia_lib/lib/grabber/filter"
	"github.com/pdcgo/tokopedia_lib/lib/model_public"
	"go.mongodb.org/mongo-driver/mongo"
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
	limitLock  sync.Mutex
	total      int32
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
		layoutVar, _ := model_public.NewPdpGetlayoutQueryVar(url)
		layout, err := g.Api.PdpGetlayoutQuery(layoutVar)

		if err != nil {
			pdc_common.ReportError(err)
			return nil
		}
		return layout
	}
}

func (g *BaseGrabber) GetPdpDataP2(ctx context.Context, layout *model_public.PdpGetlayoutQueryResp) *model_public.PdpGetDataP2Resp {
	select {

	case <-ctx.Done():
		return nil

	default:
		var pdpSess string
		var prodId string
		if layout.Data.PdpGetLayout.PdpSession != "" {
			pdpSess = layout.Data.PdpGetLayout.PdpSession
		}
		if layout.Data.PdpGetLayout.BasicInfo.ID != "" {
			prodId = layout.Data.PdpGetLayout.BasicInfo.ID
		}

		pdpVar := &model_public.PdpGetDataP2Var{
			PdpSession: pdpSess,
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

func (g *BaseGrabber) ApplyFilter(
	ctx context.Context,
	filterItem filter.FilterHandler,
	layout *model_public.PdpGetlayoutQueryResp,
	pdp *model_public.PdpGetDataP2Resp,
) (filtered, finished bool) {

	select {

	case <-ctx.Done():
		return true, false

	default:

		cek, reason, err := filterItem(layout, pdp)
		productName := layout.Data.PdpGetLayout.GetProductName()

		if err != nil {
			if errors.Is(filter.ErrLimiterReached, err) {
				return true, true
			}
			if errors.Is(filter.ErrFilterCancel, err) {
				return false, false
			}

			pdc_common.ReportError(err)
			return true, false
		}

		if cek {
			log.Printf("[ %s ] %s", reason, productName)
			return
		}

		return false, false
	}
}

func (g *BaseGrabber) SaveItem(
	ctx context.Context,
	layout *model_public.PdpGetlayoutQueryResp,
	pdp *model_public.PdpGetDataP2Resp,
) (saved bool) {

	select {

	case <-ctx.Done():
		return false

	default:
		namespace := g.GrabTasker.Namespace
		productName := layout.Data.PdpGetLayout.GetProductName()
		err := g.CacheHandler.AddProductItem(namespace, layout, pdp)

		if err != nil {
			if mongo.IsDuplicateKeyError(err) {
				log.Printf("[ duplicated ] %s - %s", namespace, productName)
				return
			}

			pdc_common.ReportError(err)
			return false
		}

		atomic.AddInt32(&g.total, 1)
		total := atomic.LoadInt32(&g.total)
		log.Printf("[ scraped ] %d item saved", total)

		return true
	}
}
