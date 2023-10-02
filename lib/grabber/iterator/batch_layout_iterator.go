package iterator

import (
	"context"
	"errors"
	"log"
	"sync"

	"github.com/pdcgo/common_conf/pdc_common"
	"github.com/pdcgo/tokopedia_lib/lib/api_public"
	"github.com/pdcgo/tokopedia_lib/lib/model_public"
)

type BatchLayoutHandler func(layout *model_public.PdpGetlayoutQueryResp) error

// deprecated
func GetBatchLayout(
	api *api_public.TokopediaApiPublic,
	ctx context.Context,
	urls []string,
	handler BatchLayoutHandler,
) (err error) {

	// api, _ := api_public.NewTokopediaApiPublic()

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
				return err
			}

			layoutVars = append(layoutVars, layoutVar)
		}
	}

	resp, err := api.PdpGetlayoutQueryBatch(layoutVars)
	if err != nil {
		if errors.Is(api_public.ErrGraphqlBatchNoQuery, err) {
			return err
		}
		pdc_common.ReportError(err)
		return err
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
				return err
			}
		}
	}

	return nil
}

func V2GetBatchLayout(
	searchItem <-chan []*model_public.ProductSearch,
	limit int,
	taskcount int,
	ctxErr ContextError,
	api *api_public.TokopediaApiPublic,
) (<-chan *model_public.PdpGetlayoutQueryResp, error) {

	result := make(chan *model_public.PdpGetlayoutQueryResp, limit)

	var gr sync.WaitGroup
	gr.Add(taskcount)

	for c := 0; c < taskcount; c++ {
		go func() {
			defer gr.Done()
		Parent:
			for items := range searchItem {
				var layoutVars []*model_public.PdpGetlayoutQueryVar
				for _, item := range items {
					layoutVar, err := model_public.NewPdpGetlayoutQueryVar(item.URL)
					if err != nil {
						ctxErr.SetError(err)
						continue
					}

					layoutVars = append(layoutVars, layoutVar)
				}

				resp, err := api.PdpGetlayoutQueryBatch(layoutVars)
				if err != nil {
					ctxErr.SetError(err)
					continue

				}
				for _, item := range resp {
					select {
					case <-ctxErr.GetCtx().Done():
						break Parent
					default:
						result <- item
					}

				}
			}
		}()
	}

	go func() {
		gr.Wait()
		log.Println("layout iterator selesai")
		close(result)
	}()

	return result, nil
}
