package iterator

import (
	"context"

	"github.com/pdcgo/tokopedia_lib/lib/api_public"
	"github.com/pdcgo/tokopedia_lib/lib/model_public"
)

type ProductShopHandler func(item *model_public.ShopProductData) error

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
			for _, item := range products {
				select {
				case <-ctx.Done():
					break Parent
				default:
					err := handler(&item)
					if err != nil {
						return err
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
