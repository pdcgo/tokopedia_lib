package grabber

import (
	"context"

	"github.com/pdcgo/go_v2_shopeelib/app/upload_app/legacy_source"
	"github.com/pdcgo/go_v2_shopeelib/lib/legacy"
	"github.com/pdcgo/tokopedia_lib/lib/grabber/filter"
	"github.com/pdcgo/tokopedia_lib/lib/grabber/iterator"
	"github.com/pdcgo/tokopedia_lib/lib/model_public"
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

	filterText := legacy_source.NewFilterText(g.Base)
	grabBasic := legacy.NewGrabBasic(g.Base)
	grabTokopedia := legacy.NewGrabTokopedia(g.Base)
	markupConfig := legacy.NewLegacyMarkupConfigWithBase(g.Base)

	ctx := context.Background()
	filters := filter.NewGrabFilterBundle(g.Api, g.Base, filterText, grabBasic, grabTokopedia, markupConfig)
	filterItem := filter.NewFilterItem(ctx, filters...)

	fname := g.Base.Path(g.GrabTasker.ProductURL)

	err := iterator.IterateUrls(fname, func(items []string) error {
		return iterator.GetBatchLayout(g.Api, ctx, items, func(layout *model_public.PdpGetlayoutQueryResp) error {

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

				if g.GrabTasker.UseFilter {
					filtered, _ := g.ApplyFilter(ctx, filterItem, layout, pdp)
					if filtered {
						return
					}
				}

				g.SaveItem(ctx, layout, pdp)
			}()

			return nil
		})
	})

	g.wg.Wait()
	return err
}
