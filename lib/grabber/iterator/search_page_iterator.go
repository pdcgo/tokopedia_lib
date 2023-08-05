package iterator

import (
	"context"
	"encoding/json"
	"math"
	"strings"
	"sync"

	"github.com/pdcgo/common_conf/pdc_common"
	"github.com/pdcgo/tokopedia_lib/lib/api_public"
	"github.com/pdcgo/tokopedia_lib/lib/model_public"
)

type SearchPageHandler func(items []*model_public.ProductSearch) error

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
	ctx context.Context,
	searchVar *model_public.SearchProductVar,
	handler SearchPageHandler,
) error {

	itemCount := searchVar.Rows
	currentCount := 0
	wg := sync.WaitGroup{}
	batchLayoutGuard := make(chan int, 10)

Parent:
	for currentCount < itemCount {
		select {
		case <-ctx.Done():
			break Parent
		default:

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

			products := resp.Data.AceSearchProductV4.Data.Products
			prodLength := len(products)
			maxArray := math.Ceil(float64(prodLength) / 10)
			for i := 1; i <= int(maxArray); i++ {
				wg.Add(1)
				batchLayoutGuard <- 1

				select {
				case <-ctx.Done():
					break Parent
				default:
					startIndex := i*10 - 10
					endIndex := i * 10
					if endIndex > prodLength {
						endIndex = prodLength
					}

					go func() {
						defer wg.Done()
						defer func() {
							<-batchLayoutGuard
						}()

						err := handler(products[startIndex:endIndex])
						if err != nil {
							pdc_common.ReportError(err)
							return
						}
					}()
				}
			}

			wg.Wait()

			itemCount = resp.Data.AceSearchProductV4.Header.TotalData
			currentCount = searchVar.Rows * searchVar.Page

			start := searchVar.Page * searchVar.Rows
			searchVar.Page += 1
			searchVar.Start = start

			if itemCount == 0 {
				break Parent
			}

		}
	}

	return nil
}
