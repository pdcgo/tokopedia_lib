package api_public

import (
	"github.com/pdcgo/tokopedia_lib/lib/model"
	"github.com/pdcgo/tokopedia_lib/lib/query"
)

func (api *TokopediaApiPublic) ShopProducts(payload *model.ShopProductVar) (*model.ShopProductResp, error) {
	gqlQuery := GraphqlPayload{
		OperationName: "ShopProducts",
		Variables:     payload,
		Query:         query.ShopProducts,
	}
	req := api.NewGraphqlReq(&gqlQuery)

	var hasil model.ShopProductResp
	err := api.SendRequest(req, &hasil)
	return &hasil, err
}
