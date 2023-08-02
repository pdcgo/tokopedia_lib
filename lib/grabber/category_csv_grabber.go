package grabber

import (
	"log"

	"github.com/pdcgo/common_conf/pdc_common"
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
	limiter, addCount := filter.CreateLimiter(g.Base)
	filters := []filter.FilterHandler{
		limiter,
		filter.CreateSoldFilter(g.Base),
		filter.CreateSoldPercentageFilter(g.Base),
		filter.CreateStockFilter(g.Base),
		filter.CreatePointFilter(g.Api, g.Base),
		filter.CreateBlacklistUsernameFilter(g.Base),
		filter.CreateLastLoginFilter(g.Base),
		filter.CreateLastReviewFilter(g.Api, g.Base),
	}
	filterItem := filter.NewFilterItem(filters...)

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

		searchVar := CreateGrabSearchVar(g.Base)
		categoryId := getCategoryId(category.Url)
		searchVar.CategoryId = categoryId

		err := iterator.IterateSearchPage(g.Api, searchVar, func(item *model_public.ProductSearch) error {

			g.wg.Add(1)
			g.limitGuard <- 1

			go func() {
				defer g.wg.Done()
				defer func() {
					<-g.limitGuard
				}()

				layoutVar, _ := ParseProductDetailParamsFromUrl(item.URL)
				layout, err := g.Api.PdpGetlayoutQuery(layoutVar)
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

				cek, reason, err := filterItem(layout, pdp)

				if err != nil {
					pdc_common.ReportError(err)
					return
				}

				if cek {
					log.Printf("[ %s ] %s", reason, item.Name)
					return
				}

				// TODO:save item
				err = g.CacheHandler.AddProductItem(g.GrabTasker.Namespace, layout, pdp)
				if err != nil {
					pdc_common.ReportError(err)
					return
				}
				addCount()
				// save berhasil -> addCount

			}()

			return nil
		})

		g.wg.Wait()
		return err
	})
}
