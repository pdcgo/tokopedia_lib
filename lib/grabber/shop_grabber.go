package grabber

import (
	"context"
	"errors"
	"log"

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

		err := iterator.IterateProductShopPage(g.Api, ctx, searchVar, func(items []*model_public.ShopProductData) error {
			var urls []string
			for _, item := range items {
				urls = append(urls, item.ProductURL)
			}

			return iterator.IterateBatchLayout(g.Api, ctx, urls, func(layout *model_public.PdpGetlayoutQueryResp) error {

				g.wg.Add(1)
				g.limitGuard <- 1

				go func() {
					defer g.wg.Done()
					go func() {
						<-g.limitGuard
					}()

					name := layout.Data.PdpGetLayout.BasicInfo.Alias
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
							log.Printf("[ %s ] %s", reason, name)
							cancel()
							return
						}
						pdc_common.ReportError(err)
						return
					}
					if cek {
						log.Printf("[ %s ] %s", reason, name)
						return
					}

					err = g.CacheHandler.AddProductItem(g.GrabTasker.Namespace, layout, pdp)
					if err != nil {
						if mongo.IsDuplicateKeyError(err) {
							log.Printf("[ duplicated ] %s - %s", g.GrabTasker.Namespace, name)
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
		})

		g.wg.Wait()
		return err
	})
}
