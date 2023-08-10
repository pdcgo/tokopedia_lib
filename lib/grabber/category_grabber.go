package grabber

import (
	"context"
	"log"
	"strconv"

	"github.com/pdcgo/go_v2_shopeelib/app/upload_app/legacy_source"
	"github.com/pdcgo/go_v2_shopeelib/lib/legacy"
	"github.com/pdcgo/tokopedia_lib/lib/grabber/filter"
	"github.com/pdcgo/tokopedia_lib/lib/grabber/iterator"
	"github.com/pdcgo/tokopedia_lib/lib/model_public"
)

type CategoryGrabber struct {
	*BaseGrabber
}

func NewCategoryGrabber(base *BaseGrabber) *CategoryGrabber {
	return &CategoryGrabber{
		base,
	}
}

func (g *CategoryGrabber) getCatId() (int, error) {

	catIds := []string{}
	for _, id := range g.GrabTasker.TokpedCateg {
		if id != "0" {
			catIds = append(catIds, id)
		}
	}

	catIdInt, err := strconv.Atoi(catIds[len(catIds)-1])
	return catIdInt, err
}

func (g *CategoryGrabber) Run() error {

	catId, err := g.getCatId()
	if catId == 0 {
		log.Println("[ warning ] tidak ada category dipilih")
		return nil
	}
	if err != nil {
		return nil
	}

	filterText := legacy_source.NewFilterText(g.Base)
	grabBasic := legacy.NewGrabBasic(g.Base)
	grabTokopedia := legacy.NewGrabTokopedia(g.Base)
	markupConfig := legacy.NewLegacyMarkupConfigWithBase(g.Base)

	filterLimit, addCount := filter.CreateLimiter(grabBasic)
	filters := []filter.FilterHandler{
		filterLimit,
	}

	filterOpt := filter.NewGrabFilterBundle(g.Api, g.Base, filterText, grabBasic, grabTokopedia, markupConfig)
	filters = append(filters, filterOpt...)

	ctx, cancel := context.WithCancel(context.Background())
	filterItem := filter.NewFilterItem(ctx, filters...)

	searchVar := CreateGrabSearchVar(grabTokopedia)
	searchVar.CategoryId = catId

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
}
