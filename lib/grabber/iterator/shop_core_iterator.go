package iterator

import (
	"net/url"
	"strconv"
	"strings"

	"github.com/pdcgo/common_conf/pdc_common"
	"github.com/pdcgo/tokopedia_lib/lib/api_public"
	"github.com/pdcgo/tokopedia_lib/lib/model_public"
	"github.com/pdcgo/tokopedia_lib/lib/query"
	"github.com/rs/zerolog"
)

type ShopCoreItem interface {
	GetShopUrl() string
	SetShopCore(data *model_public.ShopCoreInfoResp) error
	GetShopID() int
	SetStatistic(data *model_public.ShopStatisticQueryResp) error
}

func BatchShopCore[T ShopCoreItem](
	ctxErr *ContextError,
	itemsChan <-chan []T,
	limitTask int,
	limitChan int,
	api *api_public.TokopediaApiPublic,
) (<-chan []T, error) {
	filteredChan := make(chan []T, limitChan)

	go func() {
		defer close(filteredChan)

		for items := range itemsChan {
			payloads := []*api_public.GraphqlPayload{}
			for _, item := range items {
				uri, err := url.Parse(item.GetShopUrl())

				if err != nil {
					ctxErr.SendError(err)
					return
				}
				domain := strings.ReplaceAll(uri.Path, "/", "")
				pay := model_public.ShopCoreInfoVar{
					ID:     0,
					Domain: domain,
				}

				query := api_public.GraphqlPayload{
					Variables: &pay,
					Query:     query.ShopCoreInfo,
				}

				payloads = append(payloads, &query)
			}

			hasil := []*model_public.ShopCoreInfoResp{}
			req, err := api.NewGraphqlReqBatch("ShopInfoCore", payloads)
			if err != nil {
				ctxErr.SendError(err)
				return
			}
			err = api.SendRequest(req, &hasil)
			if err != nil {
				ctxErr.SendError(err)
				return
			}

			for ind, core := range hasil {
				err := items[ind].SetShopCore(core)
				if err != nil {
					pdc_common.ReportErrorCustom(err, func(event *zerolog.Event) *zerolog.Event {

						return event.Interface("core", hasil).Interface("itemcore", core)
					})
					ctxErr.SendError(err)
					return
				}
			}

			select {
			case filteredChan <- items:
				continue
			case <-ctxErr.Ctx.Done():
				continue
			}
		}

	}()

	return filteredChan, nil

}

func BatchShopStatistic[T ShopCoreItem](
	ctxErr *ContextError,
	itemsChan <-chan []T,
	limitTask int,
	limitChan int,
	api *api_public.TokopediaApiPublic,
) (<-chan []T, error) {
	filteredChan := make(chan []T, limitChan)

	go func() {
		defer close(filteredChan)

		for items := range itemsChan {
			payloads := []*api_public.GraphqlPayload{}
			for _, item := range items {

				shopID := item.GetShopID()
				shopIDstr := strconv.Itoa(shopID)

				pay := model_public.ShopStatisticQueryVar{
					ShopID:    shopID,
					ShopIDStr: shopIDstr,
				}

				query := api_public.GraphqlPayload{
					OperationName: "ShopStatisticQuery",
					Variables:     &pay,
					Query:         query.ShopStatisticQuery,
				}

				payloads = append(payloads, &query)
			}

			hasil := []*model_public.ShopStatisticQueryResp{}
			req, err := api.NewGraphqlReqBatch("ShopStatisticQuery", payloads)
			if err != nil {
				ctxErr.SendError(err)
				return
			}
			err = api.SendRequest(req, &hasil)
			if err != nil {
				ctxErr.SendError(err)
				return
			}

			for ind, stat := range hasil {
				err := items[ind].SetStatistic(stat)
				if err != nil {
					ctxErr.SendError(err)
					return
				}
			}

			select {
			case filteredChan <- items:
				continue
			case <-ctxErr.Ctx.Done():
				continue
			}
		}

	}()

	return filteredChan, nil

}
