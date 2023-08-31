package iterator

import (
	"net/url"
	"strconv"
	"strings"

	"github.com/pdcgo/tokopedia_lib/lib/api_public"
	"github.com/pdcgo/tokopedia_lib/lib/model_public"
	"github.com/pdcgo/tokopedia_lib/lib/query"
)

type ShopCoreItem interface {
	GetShopUrl() string
	SetShopCore(data *model_public.ShopCoreInfoResp)
	GetShopID() int
	SetStatistic(data *model_public.ShopStatisticQueryResp)
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
		payloads := []*api_public.GraphqlPayload{}
		for items := range itemsChan {
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
				items[ind].SetShopCore(core)
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
		payloads := []*api_public.GraphqlPayload{}
		for items := range itemsChan {
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
				items[ind].SetStatistic(stat)
			}
		}

	}()

	return filteredChan, nil

}
