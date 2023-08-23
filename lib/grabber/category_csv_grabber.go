package grabber

import (
	"context"
	"errors"
	"log"

	"github.com/pdcgo/common_conf/pdc_common"
	"github.com/pdcgo/go_v2_shopeelib/app/upload_app/legacy_source"
	"github.com/pdcgo/go_v2_shopeelib/lib/legacy"
	"github.com/pdcgo/tokopedia_lib/lib/csv"
	"github.com/pdcgo/tokopedia_lib/lib/grabber/filter"
	"github.com/pdcgo/tokopedia_lib/lib/grabber/iterator"
	"github.com/pdcgo/tokopedia_lib/lib/model_public"
	"github.com/rs/zerolog"
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

	headerData, err := g.Api.HeaderMainData()
	if err != nil {
		return err
	}
	categories := headerData.Data.CategoryAllListLite.Categories

	return iterator.IterateCategoryCsv(g.Base, func(item *csv.CategoryCsv) error {

		log.Println("[ info ] grab category", item.Name)

		filterLimit, addCount := filter.CreateLimiter(grabBasic)
		filters := []filter.FilterHandler{
			filterLimit,
		}

		filters = append(filters, filtersOpt...)

		category, err := categories.GetCategoryByUrl(item.Url)
		if err != nil || category.ID == 0 {

			if err == nil {
				err = errors.New("no category")
			}

			pdc_common.ReportErrorCustom(err, func(event *zerolog.Event) *zerolog.Event {
				return event.Str("category", item.Name).Str("url", item.Url)
			})

			return nil
		}

		searchVar := model_public.NewGrabSearchProductVar(grabTokopedia)
		searchVar.CategoryId = category.ID

		ctx, cancel := context.WithCancel(context.Background())
		filterItem := filter.NewFilterItem(ctx, filters...)

		err = iterator.IterateSearchPage(g.Api, ctx, searchVar, func(items []*model_public.ProductSearch) error {

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
