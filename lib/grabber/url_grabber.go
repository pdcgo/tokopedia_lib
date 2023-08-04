package grabber

import (
	"context"
	"log"
	"sync"

	"github.com/pdcgo/common_conf/pdc_common"
	"github.com/pdcgo/tokopedia_lib/lib/grabber/filter"
	"github.com/pdcgo/tokopedia_lib/lib/grabber/iterator"
	"github.com/pdcgo/tokopedia_lib/lib/helper"
	"github.com/pdcgo/tokopedia_lib/lib/model_public"
	"go.mongodb.org/mongo-driver/mongo"
)

type UrlGrabber struct {
	*BaseGrabber
}

func NewUrlGrabber(base *BaseGrabber) *UrlGrabber {
	return &UrlGrabber{
		base,
	}
}

func (g *UrlGrabber) Run() error {
	filters := []filter.FilterHandler{
		filter.CreateSoldFilter(g.Base),
		filter.CreateSoldPercentageFilter(g.Base),
		filter.CreateStockFilter(g.Base),
		filter.CreatePointFilter(g.Api, g.Base),
		filter.CreateBlacklistUsernameFilter(g.Base),
		filter.CreateLastLoginFilter(g.Base),
		filter.CreateLastReviewFilter(g.Api, g.Base),
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	filterItem := filter.NewFilterItem(ctx, filters...)

	counter := helper.NewCounter()
	lock := sync.Mutex{}

	err := iterator.IterateUrls(g.Base, g.GrabTasker, func(item string) error {
		g.wg.Add(1)

		go func() {
			defer g.wg.Done()
			go func() {
				<-g.limitGuard
			}()

			layoutVar, _ := ParseProductDetailParamsFromUrl(item)
			lock.Lock()
			layout, err := g.Api.PdpGetlayoutQuery(layoutVar)
			lock.Unlock()
			if err != nil {
				pdc_common.ReportError(err)
				return
			}

			pdpVar := &model_public.PdpGetDataP2Var{
				PdpSession: layout.Data.PdpGetLayout.PdpSession,
				ProductID:  layout.Data.PdpGetLayout.BasicInfo.ID,
			}
			pdp, err := g.Api.PdpGetDataP2(pdpVar)
			if err != nil {
				pdc_common.ReportError(err)
				return
			}

			if g.GrabTasker.UseFilter {
				cek, reason, err := filterItem(layout, pdp)
				if err != nil {
					pdc_common.ReportError(err)
					return
				}
				if cek {
					log.Printf("[ %s ] %s", reason, layoutVar.ProductKey)
					return
				}
			}

			err = g.CacheHandler.AddProductItem(g.GrabTasker.Namespace, layout, pdp)
			if err != nil {
				if mongo.IsDuplicateKeyError(err) {
					log.Printf("[ duplicated ] %s - %s", g.GrabTasker.Namespace, layoutVar.ProductKey)
					return
				}
				pdc_common.ReportError(err)
				return
			}
			counter.Add()
			log.Printf("[ scraped ] %d item saved", counter.Count())
		}()

		return nil
	})

	g.wg.Wait()
	return err
}
