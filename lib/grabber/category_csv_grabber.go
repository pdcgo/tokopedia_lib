package grabber

import (
	"context"
	"log"

	"github.com/pdcgo/go_v2_shopeelib/app/upload_app/legacy_source"
	"github.com/pdcgo/go_v2_shopeelib/lib/legacy"
	"github.com/pdcgo/tokopedia_lib/lib/csv"
	"github.com/pdcgo/tokopedia_lib/lib/grabber/filter"
	"github.com/pdcgo/tokopedia_lib/lib/grabber/iterator"
	"github.com/pdcgo/tokopedia_lib/lib/model_public"
)

type CategoryCsvGrabber struct {
	*BaseGrabber
}

func NewCategoryCsvGrabber(base *BaseGrabber) *CategoryCsvGrabber {
	return &CategoryCsvGrabber{base}
}

func (g *CategoryCsvGrabber) Run() error {

	filterText := legacy_source.NewFilterText(g.Base)
	grabBasic := legacy.NewGrabBasic(g.Base)
	grabTokopedia := legacy.NewGrabTokopedia(g.Base)
	markupConfig := legacy.NewLegacyMarkupConfigWithBase(g.Base)

	filtersOpt := filter.NewGrabFilterBundle(g.Api, g.Base, filterText, grabBasic, grabTokopedia, markupConfig)

	categories, err := g.Api.HeaderMainData()
	if err != nil {
		return err
	}

	getCategoryId := func(url string) int {
		res := csv.GetCategoryByUrl(categories.Data.CategoryAllListLite.Categories, url)
		category := <-res
		return category.ID
	}

	return iterator.IterateCategoryCsv(g.Base, func(category *csv.CategoryCsv) error {

		log.Println("[ info ] grab category", category.Name)

		filterLimit, addCount := filter.CreateLimiter(grabBasic)
		filters := []filter.FilterHandler{
			filterLimit,
		}

		filters = append(filters, filtersOpt...)

		searchVar := model_public.NewGrabSearchProductVar(grabTokopedia)
		categoryId := getCategoryId(category.Url)
		searchVar.CategoryId = categoryId

		ctx, cancel := context.WithCancel(context.Background())
		filterItem := filter.NewFilterItem(ctx, filters...)

		err := iterator.IterateSearchPage(g.Api, ctx, searchVar, func(items []*model_public.ProductSearch) error {

			var urls []string
			for _, item := range items {
				urls = append(urls, item.URL)
			}

			return iterator.IterateBatchLayout(g.Api, ctx, urls, func(layout *model_public.PdpGetlayoutQueryResp) error {
				g.wg.Add(1)
				g.limitGuard <- 1

				go func() {
					defer g.wg.Done()
					defer func() {
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
