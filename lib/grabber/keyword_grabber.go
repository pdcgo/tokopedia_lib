package grabber

import (
	"context"
	"net/url"

	"github.com/pdcgo/go_v2_shopeelib/app/upload_app/legacy_source"
	"github.com/pdcgo/go_v2_shopeelib/lib/legacy"
	"github.com/pdcgo/tokopedia_lib/lib/grabber/filter"
	"github.com/pdcgo/tokopedia_lib/lib/grabber/iterator"
	"github.com/pdcgo/tokopedia_lib/lib/model_public"
)

type KeywordGrabber struct {
	*BaseGrabber
}

func NewKeywordGrabber(base *BaseGrabber) *KeywordGrabber {
	return &KeywordGrabber{
		base,
	}
}

func (g *KeywordGrabber) Run() error {

	filterText := legacy_source.NewFilterText(g.Base)
	grabBasic := legacy.NewGrabBasic(g.Base)
	grabTokopedia := legacy.NewGrabTokopedia(g.Base)
	markupConfig := legacy.NewLegacyMarkupConfigWithBase(g.Base)

	filtersOpt := filter.NewGrabFilterBundle(g.Api, g.Base, filterText, grabBasic, grabTokopedia, markupConfig)

	fname := g.Base.Path(g.GrabTasker.Keyword)

	return iterator.IterateKeywords(fname, func(item string) error {

		filterLimit, addCount := filter.CreateLimiter(grabBasic)
		filters := []filter.FilterHandler{
			filterLimit,
		}

		if g.GrabTasker.UseFilter {
			filters = append(filters, filtersOpt...)
		}

		searchVar := model_public.NewGrabSearchProductVar(grabTokopedia)
		searchVar.Query = url.QueryEscape(item)

		ctx, cancel := context.WithCancel(context.Background())
		filterItem := filter.NewFilterItem(ctx, filters...)

		cfg := iterator.IterateConfig{
			ChuckSize:      3,
			ConcurentGuard: make(chan int, 5),
		}
		err := iterator.IterateSearchPage(&cfg, g.Api, ctx, searchVar, func(items []*model_public.ProductSearch) error {

			var urls []string
			for _, item := range items {
				urls = append(urls, item.URL)
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
