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

	hasMore := true
Parent:
	for hasMore {
		select {
		case <-ctx.Done():
			break Parent
		default:
			resp, err := api.ShopProducts(searchVar)
			if err != nil {
				return err
			}

			for _, item := range resp.Data.GetShopProduct.Data {
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

			if resp.Data.GetShopProduct.Links.Next != "" {
				hasMore = false
			}

			searchVar.Page += 1
		}

	}

	return nil
}
