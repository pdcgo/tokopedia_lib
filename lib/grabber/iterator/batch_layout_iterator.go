package iterator

import (
	"context"
	"errors"
	"sync"
	"time"

	"github.com/pdcgo/common_conf/pdc_common"
	"github.com/pdcgo/tokopedia_lib/lib/api_public"
	"github.com/pdcgo/tokopedia_lib/lib/model_public"
)

type BatchLayoutHandler func(layout *model_public.PdpGetlayoutQueryResp) error

func IterateBatchLayout(
	api *api_public.TokopediaApiPublic,
	ctx context.Context,
	urls []string,
	handler BatchLayoutHandler,
) (err error) {

	wg := sync.WaitGroup{}
	batchLayoutGuard := make(chan int, 3)

	wg.Add(1)
	batchLayoutGuard <- 1

	go func() {
		defer wg.Done()
		defer func() {
			time.Sleep(time.Second)
			<-batchLayoutGuard
		}()

		var layoutVars []*model_public.PdpGetlayoutQueryVar

	Urls:
		for _, url := range urls {
			select {

			case <-ctx.Done():
				break Urls

			default:
				layoutVar, err := model_public.NewPdpGetlayoutQueryVar(url)
				if err != nil {
					pdc_common.ReportError(err)
					return
				}

				layoutVars = append(layoutVars, layoutVar)
			}
		}

		resp, err := api.PdpGetlayoutQueryBatch(layoutVars)
		if err != nil {
			if errors.Is(api_public.ErrGraphqlBatchNoQuery, err) {
				return
			}
			pdc_common.ReportError(err)
			return
		}

	Resp:
		for _, resp := range resp {
			select {

			case <-ctx.Done():
				break Resp

			default:
				err := handler(resp)
				if err != nil {
					pdc_common.ReportError(err)
					return
				}
			}
		}
	}()

	wg.Wait()

	return nil
}
