package iterator

import (
	"context"
	"math"

	"github.com/pdcgo/common_conf/pdc_common"
	"github.com/pdcgo/tokopedia_lib/lib/api_public"
	"github.com/pdcgo/tokopedia_lib/lib/model_public"
)

type ProductShopHandler func(items []*model_public.ShopProductData) error

func IterateProductShopPage(
	api *api_public.TokopediaApiPublic,
	ctx context.Context,
	searchVar *model_public.ShopProductVar,
	handler ProductShopHandler,
) error {

	itemCount := searchVar.PerPage
	currentCount := 0

Parent:
	for currentCount < itemCount {
		select {
		case <-ctx.Done():
			break Parent
		default:

			resp, err := api.ShopProducts(searchVar)
			if err != nil {
				return err
			}

			products := resp.Data.GetShopProduct.Data
			prodLength := len(products)
			maxArray := math.Ceil(float64(prodLength) / 10)
			for i := 1; i <= int(maxArray); i++ {
				select {
				case <-ctx.Done():
					break Parent
				default:

					startIndex := i*10 - 10
					endIndex := i * 10
					if endIndex > prodLength {
						endIndex = prodLength
					}

					err := handler(products[startIndex:endIndex])
					if err != nil {
						pdc_common.ReportError(err)
					}
				}
			}

			itemCount = searchVar.PerPage * len(products)
			currentCount = searchVar.PerPage * searchVar.Page
			searchVar.Page += 1

			if itemCount == 0 {
				break Parent
			}

		}
	}

	return nil
}
