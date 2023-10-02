package iterator

import (
	"context"
	"errors"
	"log"
	"sync"

	"github.com/pdcgo/common_conf/pdc_common"
	"github.com/pdcgo/tokopedia_lib/lib/api_public"
	"github.com/pdcgo/tokopedia_lib/lib/model_public"
	"github.com/rs/zerolog"
)

type SearchPageHandler func(items []*model_public.ProductSearch) error
type IterateConfig struct {
	ChuckSize      int
	ConcurentGuard chan int
}

// deprecated
func IterateSearchPage(
	cfg *IterateConfig, // tidak berguna
	api *api_public.TokopediaApiPublic,
	ctx context.Context,
	searchVar *model_public.SearchProductVar,
	handler SearchPageHandler,
) error {

	itemCount := searchVar.Rows
	currentCount := 0

Parent:
	for currentCount < itemCount {
		var wg sync.WaitGroup
		select {
		case <-ctx.Done():
			break Parent
		default:

			variable := &model_public.ParamsVar{
				Params: searchVar.GetQuery(),
			}
			resp, err := api.SearchProductQueryV4(variable)
			if err != nil {
				return err
			}

			products := resp.Data.AceSearchProductV4.Data.Products
			for _, items := range products.Chunks(cfg.ChuckSize) {
				select {
				case <-ctx.Done():
					return ctx.Err()

				default:
					wg.Add(1)
					ps := items
					go func() {
						defer wg.Done()
						err := handler(ps)
						if err != nil {
							pdc_common.ReportError(err)
						}
					}()
				}

			}

			if err != nil {
				if errors.Is(err, context.Canceled) || errors.Is(err, context.DeadlineExceeded) {
					break Parent
				} else {
					pdc_common.ReportError(err)
				}
			}

			itemCount = resp.Data.AceSearchProductV4.Header.TotalData
			currentCount = searchVar.Rows * searchVar.Page

			start := searchVar.Page * searchVar.Rows
			searchVar.Page += 1
			searchVar.Start = start

			if itemCount == 0 {
				break Parent
			}

		}

		wg.Wait()
	}

	return nil
}

type ContextError interface {
	GetCtx() context.Context
	SetError(err error)
	SetErrorCustom(err error, handler func(event *zerolog.Event) *zerolog.Event)
	Cancel()
}

func V2IterateSearchPage(
	ctxErr ContextError,
	chunkSize int,
	api *api_public.TokopediaApiPublic,
	searchVar *model_public.SearchProductVar,
) (<-chan []*model_public.ProductSearch, error) {

	ctx := ctxErr.GetCtx()
	currentCount := 0
	chunkChan := make(chan []*model_public.ProductSearch, 100)

	variable := &model_public.ParamsVar{
		Params: searchVar.GetQuery(),
	}
	resp, err := api.SearchProductQueryV4(variable)
	if err != nil {
		return chunkChan, err
	}
	itemCount := resp.Data.AceSearchProductV4.Header.TotalData

	if itemCount == 0 {
		return chunkChan, errors.New(" count 0")
	}

	go func() {
		defer close(chunkChan)
	Parent:
		for currentCount < itemCount {
			select {
			case <-ctx.Done():
				break Parent
			default:

				variable := &model_public.ParamsVar{
					Params: searchVar.GetQuery(),
				}

				searchVar.Page += 1
				currentCount = searchVar.Rows * searchVar.Page
				start := searchVar.Page * searchVar.Rows
				searchVar.Start = start

				resp, err := api.SearchProductQueryV4(variable)
				if err != nil {
					ctxErr.SetError(err)
					return
				}
				products := resp.Data.AceSearchProductV4.Data.Products
				if len(products) == 0 {
					break Parent
				}

				for _, items := range products.Chunks(chunkSize) {
					chunkChan <- items

				}

			}
		}

		log.Println("done iterate chunk")

	}()

	return chunkChan, nil
}
