package iterator

import (
	"encoding/json"
	"strings"

	"github.com/pdcgo/tokopedia_lib/lib/api_public"
	"github.com/pdcgo/tokopedia_lib/lib/helper"
	"github.com/pdcgo/tokopedia_lib/lib/model_public"
)

type SearchPageHandler func(item *model_public.ProductSearch) error

func createSearchParams(searchVar *model_public.SearchProductVar) (string, error) {
	rawParams, err := json.Marshal(searchVar)
	if err != nil {
		return "", err
	}
	stringParams := string(rawParams)
	replacer := strings.NewReplacer(`"`, "", ":", "=", ",", "&", " ", "", "[", "", "]", "")
	validParams := replacer.Replace(stringParams)

	return validParams[1 : len(validParams)-1], nil
}

func IterateSearchPage(
	api *api_public.TokopediaApiPublic,
	// ctx context.Context,
	limiter *helper.Limiter,
	searchVar *model_public.SearchProductVar,
	handler SearchPageHandler,
) error {

	itemCount := searchVar.Rows
	currentCount := 0

Parent:
	for currentCount < itemCount {
		// select {
		// case <-ctx.Done():
		// 	fmt.Println("stop iterate page")
		// 	break Parent
		// default:
		if limiter.LimitReached() {
			break Parent
		}
		params, err := createSearchParams(searchVar)
		if err != nil {
			return err
		}

		variable := &model_public.ParamsVar{
			Params: params,
		}
		resp, err := api.SearchProductQueryV4(variable)
		if err != nil {
			return err
		}

		// Item:
		for _, item := range resp.Data.AceSearchProductV4.Data.Products {
			// select {
			// case <-ctx.Done():
			// 	fmt.Println("stop iterate product")
			// 	break Parent
			// default:
			if limiter.LimitReached() {
				break Parent
			}
			err := handler(item)
			if err != nil {
				return err
			}
			
			limiter.Add()

			// }
		}

		itemCount = resp.Data.AceSearchProductV4.Header.TotalData
		currentCount = searchVar.Rows * searchVar.Page

		start := searchVar.Page * searchVar.Rows
		searchVar.Page += 1
		searchVar.Start = start

		if itemCount == 0 {
			break Parent
		}
		// }

	}

	return nil
}
