package grabber

import (
	"context"
	"strconv"

	"github.com/pdcgo/go_v2_shopeelib/app/upload_app/legacy_source"
	"github.com/pdcgo/go_v2_shopeelib/lib/legacy"
	"github.com/pdcgo/tokopedia_lib/lib/grabber/filter"
	"github.com/pdcgo/tokopedia_lib/lib/grabber/iterator"
	"github.com/pdcgo/tokopedia_lib/lib/model_public"
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

	filterText := legacy_source.NewFilterText(g.Base)
	grabBasic := legacy.NewGrabBasic(g.Base)
	grabTokopedia := legacy.NewGrabTokopedia(g.Base)
	markupConfig := legacy.NewLegacyMarkupConfigWithBase(g.Base)

	filtersOpt := filter.NewGrabFilterBundle(g.Api, g.Base, filterText, grabBasic, grabTokopedia, markupConfig)

	return iterator.IterateShops(g.Api, fname, func(shopCore *model_public.ShopCoreInfoResp) error {

		filterLimit, addCount := filter.CreateLimiter(grabBasic)
		filters := []filter.FilterHandler{
			filterLimit,
		}

		if g.GrabTasker.UseFilter {
			filters = append(filters, filtersOpt...)
		}

		shopId := strconv.Itoa(shopCore.Data.Result[0].ShopCore.ShopID)

		searchVar := model_public.NewShopProductVar(shopId)

		ctx, cancel := context.WithCancel(context.Background())
		filterItem := filter.NewFilterItem(ctx, filters...)

		err := iterator.IterateProductShopPage(g.Api, ctx, searchVar, func(items []*model_public.ShopProductData) error {

			var urls []string
			for _, item := range items {
				urls = append(urls, item.ProductURL)
			}

			return iterator.GetBatchLayout(g.Api, ctx, urls, func(layout *model_public.PdpGetlayoutQueryResp) error {

				g.wg.Add(1)
				g.limitGuard <- 1

				go func() {
					defer g.wg.Done()
					go func() {
						<-g.limitGuard
					}()

					pdp := g.GetPdpDataP2(ctx, layout)
					if pdp == nil {
						return
					}

					filtered, finished := g.ApplyFilter(ctx, filterItem, layout, pdp)
					if finished {
						cancel()
					}
					if filtered {
						return
					}

					g.limitLock.Lock()
					defer g.limitLock.Unlock()

					saved := g.SaveItem(ctx, layout, pdp)
					if saved {
						limitReached := addCount()
						if limitReached {
							cancel()
						}
					}
				}()

				return nil
			})
		})

		g.wg.Wait()
		return err
	})
}
