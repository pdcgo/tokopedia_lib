package grabber

import (
	"errors"
	"log"
	"strconv"
	"sync"

	"github.com/pdcgo/common_conf/pdc_common"
	"github.com/pdcgo/tokopedia_lib/lib/grabber/filter"
	"github.com/pdcgo/tokopedia_lib/lib/grabber/iterator"
	"github.com/pdcgo/tokopedia_lib/lib/model_public"
	"go.mongodb.org/mongo-driver/mongo"
)

type CategoryGrabber struct {
	*BaseGrabber
}

func NewCategoryGrabber(base *BaseGrabber) *CategoryGrabber {
	return &CategoryGrabber{
		base,
	}
}

func (g *CategoryGrabber) Run() error {
	filterLimit, limiter := filter.CreateLimiter(g.Base)
	filtersOpt := []filter.FilterHandler{
		filter.CreateSoldFilter(g.Base),
		filter.CreateSoldPercentageFilter(g.Base),
		filter.CreateStockFilter(g.Base),
		filter.CreatePointFilter(g.Api, g.Base),
		filter.CreateBlacklistUsernameFilter(g.Base),
		filter.CreateLastLoginFilter(g.Base),
		filter.CreateLastReviewFilter(g.Api, g.Base),
	}
	filters := []filter.FilterHandler{
		filterLimit,
	}

	if g.GrabTasker.UseFilter {
		filters = append(filters, filtersOpt...)
	}

	filterItem := filter.NewFilterItem(filters...)

	lock := sync.Mutex{}

	filterCatId := func(ids []string) []string {
		var res []string
		for _, id := range ids {
			if id != "0" {
				res = append(res, id)
			}
		}
		return res
	}

	catId := filterCatId(g.GrabTasker.TokpedCateg)
	catIdInt, err := strconv.Atoi(catId[len(catId)-1])
	if err != nil {
		return err
	}

	searchVar := CreateGrabSearchVar(g.Base)
	searchVar.CategoryId = catIdInt

	// ctx, cancel := context.WithCancel(context.Background())

	err = iterator.IterateSearchPage(g.Api, limiter, searchVar, func(item *model_public.ProductSearch) error {
		g.wg.Add(1)
		g.limitGuard <- 1

		go func() {
			defer g.wg.Done()
			defer func() {
				<-g.limitGuard
			}()

			layoutVar, _ := ParseProductDetailParamsFromUrl(item.URL)
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

			cek, reason, err := filterItem(layout, pdp)
			if err != nil {
				if errors.Is(filter.ErrLimiterReached, err) {
					// cancel()
					return
				}
				pdc_common.ReportError(err)
				return
			}
			if cek {
				log.Printf("[ %s ] %s", reason, layoutVar.ProductKey)
				return
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

			log.Printf("[ scraped ] item saved")
			// limiter.Add()
		}()

		return nil
	})

	g.wg.Wait()
	return err

}
