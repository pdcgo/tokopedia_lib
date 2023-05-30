package api_public

import (
	"github.com/pdcgo/tokopedia_lib/lib/model_public"
	"github.com/pdcgo/tokopedia_lib/lib/query"
)

func (api *TokopediaApiPublic) ShopProducts(payload *model_public.ShopProductVar) (*model_public.ShopProductResp, error) {
	gqlQuery := GraphqlPayload{
		OperationName: "ShopProducts",
		Variables:     payload,
		Query:         query.ShopProducts,
	}
	req := api.NewGraphqlReq(&gqlQuery)

	var hasil model_public.ShopProductResp
	err := api.SendRequest(req, &hasil)
	return &hasil, err
}
