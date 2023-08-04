package grabber

import (
	"context"
	"errors"
	"log"
	"net/url"
	"sync"

	"github.com/pdcgo/common_conf/pdc_common"
	"github.com/pdcgo/tokopedia_lib/lib/grabber/filter"
	"github.com/pdcgo/tokopedia_lib/lib/grabber/iterator"
	"github.com/pdcgo/tokopedia_lib/lib/helper"
	"github.com/pdcgo/tokopedia_lib/lib/model_public"
	"go.mongodb.org/mongo-driver/mongo"
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
	filtersOpt := []filter.FilterHandler{
		filter.CreateSoldFilter(g.Base),
		filter.CreateSoldPercentageFilter(g.Base),
		filter.CreateStockFilter(g.Base),
		filter.CreatePointFilter(g.Api, g.Base),
		filter.CreateBlacklistUsernameFilter(g.Base),
		filter.CreateLastLoginFilter(g.Base),
		filter.CreateLastReviewFilter(g.Api, g.Base),
	}

	lock := sync.Mutex{}
	counter := helper.NewCounter()

	return iterator.IterateKeywords(g.Base, g.GrabTasker, func(item string) error {
		filterLimit, addCount := filter.CreateLimiter(g.Base)
		filters := []filter.FilterHandler{
			filterLimit,
		}
		if g.GrabTasker.UseFilter {
			filters = append(filters, filtersOpt...)
		}

		searchVar := CreateGrabSearchVar(g.Base)
		searchVar.Query = url.QueryEscape(item)

		ctx, cancel := context.WithCancel(context.Background())
		filterItem := filter.NewFilterItem(ctx, filters...)

		err := iterator.IterateSearchPage(g.Api, ctx, searchVar, func(item *model_public.ProductSearch) error {
			g.wg.Add(1)

			go func() {
				defer g.wg.Done()
				go func() {
					<-g.limitGuard
				}()

				lock.Lock()
				layout := g.GetProductLayout(ctx, item.URL)
				lock.Unlock()
				if layout == nil {
					return
				}

				var pdpSess string
				var prodId string
				if layout.Data.PdpGetLayout.PdpSession != "" {
					pdpSess = layout.Data.PdpGetLayout.PdpSession
				}
				if layout.Data.PdpGetLayout.BasicInfo.ID != "" {
					prodId = layout.Data.PdpGetLayout.BasicInfo.ID
				}
				pdp := g.GetPdpDataP2(ctx, pdpSess, prodId)
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
	})
}
