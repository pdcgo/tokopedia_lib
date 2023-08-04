package grabber

import (
	"context"
	"errors"
	"log"
	"strconv"
	"sync"

	"github.com/pdcgo/common_conf/pdc_common"
	"github.com/pdcgo/tokopedia_lib/lib/grabber/filter"
	"github.com/pdcgo/tokopedia_lib/lib/grabber/iterator"
	"github.com/pdcgo/tokopedia_lib/lib/helper"
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
	filterLimit, addCount := filter.CreateLimiter(g.Base)
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

	ctx, cancel := context.WithCancel(context.Background())
	filterItem := filter.NewFilterItem(ctx, filters...)

	counter := helper.NewCounter()
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
		cancel()
		return err
	}

	searchVar := CreateGrabSearchVar(g.Base)
	searchVar.CategoryId = catIdInt

	err = iterator.IterateSearchPage(g.Api, ctx, searchVar, func(item *model_public.ProductSearch) error {
		g.wg.Add(1)
		g.limitGuard <- 1

		go func() {
			defer g.wg.Done()
			defer func() {
				<-g.limitGuard
			}()

			lock.Lock()
			layout := g.GetProductLayout(ctx, item.URL)
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

}
