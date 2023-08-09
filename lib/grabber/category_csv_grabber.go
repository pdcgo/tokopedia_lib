package grabber

import (
	"context"
	"errors"
	"log"

	"github.com/pdcgo/common_conf/pdc_common"
	"github.com/pdcgo/tokopedia_lib/lib/csv"
	"github.com/pdcgo/tokopedia_lib/lib/grabber/filter"
	"github.com/pdcgo/tokopedia_lib/lib/grabber/iterator"
	"github.com/pdcgo/tokopedia_lib/lib/helper"
	"github.com/pdcgo/tokopedia_lib/lib/model_public"
	"go.mongodb.org/mongo-driver/mongo"
)

type CategoryCsvGrabber struct {
	*BaseGrabber
}

func NewCategoryCsvGrabber(base *BaseGrabber) *CategoryCsvGrabber {
	return &CategoryCsvGrabber{base}
}

func (g *CategoryCsvGrabber) Run() error {
	filtersOpt := []filter.FilterHandler{
		filter.CreateSoldFilter(g.Base),
		filter.CreateSoldPercentageFilter(g.Base),
		filter.CreateStockFilter(g.Base),
		filter.CreatePointFilter(g.Api, g.Base),
		filter.CreateBlacklistUsernameFilter(g.Base),
		filter.CreateLastLoginFilter(g.Base),
		filter.CreateLastReviewFilter(g.Api, g.Base),
	}

	counter := helper.NewCounter()

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
		filterLimit, addCount := filter.CreateLimiter(g.Base)
		filters := []filter.FilterHandler{
			filterLimit,
		}
		filters = append(filters, filtersOpt...)

		searchVar := CreateGrabSearchVar(g.Base)
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
