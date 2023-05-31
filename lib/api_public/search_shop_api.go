package api_public

import (
	"github.com/pdcgo/tokopedia_lib/lib/model_public"
	"github.com/pdcgo/tokopedia_lib/lib/query"
)

func (api *TokopediaApiPublic) AceSearchShop(payload *model_public.ParamsVar) (*model_public.AceSearchShopResp, error) {
	gqlQuery := GraphqlPayload{
		OperationName: "AceSearchShop",
		Variables:     payload,
		Query:         query.AceSearchShop,
	}
	req := api.NewGraphqlReq(&gqlQuery)

	var hasil model_public.AceSearchShopResp
	err := api.SendRequest(req, &hasil)
	return &hasil, err

}
