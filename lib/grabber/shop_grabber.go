package grabber

import (
	"context"
	"errors"
	"log"
	"sync"

	"github.com/pdcgo/common_conf/pdc_common"
	"github.com/pdcgo/tokopedia_lib/lib/grabber/filter"
	"github.com/pdcgo/tokopedia_lib/lib/grabber/iterator"
	"github.com/pdcgo/tokopedia_lib/lib/helper"
	"github.com/pdcgo/tokopedia_lib/lib/model_public"
	"go.mongodb.org/mongo-driver/mongo"
)

type ShopGrabber struct {
	*BaseGrabber
}

func NewShopGrabber(base *BaseGrabber) *ShopGrabber {
	return &ShopGrabber{
		base,
	}
}

func (g *ShopGrabber) Run() error {
	fname := g.Base.Path(g.GrabTasker.TokoUsername)

	filtersOpt := []filter.FilterHandler{
		filter.CreateBlacklistUsernameFilter(g.Base),
		filter.CreateSoldFilter(g.Base),
		filter.CreateSoldPercentageFilter(g.Base),
		filter.CreateStockFilter(g.Base),
		filter.CreatePointFilter(g.Api, g.Base),
		filter.CreateLastLoginFilter(g.Base),
		filter.CreateLastReviewFilter(g.Api, g.Base),
	}

	lock := sync.Mutex{}
	counter := helper.NewCounter()

	return iterator.IterateShops(g.Api, fname, func(shopCore *model_public.ShopCoreInfoResp) error {
		filterLimit, addCount := filter.CreateLimiter(g.Base)
		filters := []filter.FilterHandler{
			filterLimit,
		}
		if g.GrabTasker.UseFilter {
			filters = append(filters, filtersOpt...)
		}

		shopId := shopCore.Data.Result[0].ShopCore.ShopID
		searchVar := GenerateShopProductVar()
		searchVar.Sid = shopId

		ctx, cancel := context.WithCancel(context.Background())
		filterItem := filter.NewFilterItem(ctx, filters...)

		err := iterator.IterateProductShopPage(g.Api, ctx, searchVar, func(item *model_public.ShopProductData) error {
			g.wg.Add(1)

			go func() {
				defer g.wg.Done()
				go func() {
					<-g.limitGuard
				}()

				lock.Lock()
				layout := g.GetProductLayout(ctx, item.ProductURL)
				lock.Unlock()
				if layout == nil {
					return
				}

				pdp := g.GetPdpDataP2(ctx, layout)
				if layout == nil {
					return
				}

				cek, reason, err := filterItem(layout, pdp)
				if err != nil {
					if errors.Is(filter.ErrLimiterReached, err) {
						cancel()
						return
					}
					if errors.Is(filter.ErrFilterCancel, err) {
						return
					}
					if errors.Is(filter.ErrBlacklistUsername, err) {
						log.Printf("[ %s ] %s", reason, item.Name)
						cancel()
						return
					}
					pdc_common.ReportError(err)
					return
				}
				if cek {
					log.Printf("[ %s ] %s", reason, item.Name)
					return
				}

				err = g.CacheHandler.AddProductItem(g.GrabTasker.Namespace, layout, pdp)
				if err != nil {
					if mongo.IsDuplicateKeyError(err) {
						log.Printf("[ duplicated ] %s - %s", g.GrabTasker.Namespace, item.Name)
						return
					}
					pdc_common.ReportError(err)
					return
				}

				addCount()
				counter.Add()
				log.Printf("[ scraped ] %d item saved", counter.Count())
			}()

			return nil
		})

		g.wg.Wait()
		return err
	})
}
