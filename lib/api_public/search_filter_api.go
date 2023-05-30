package api_public

import (
	"github.com/pdcgo/tokopedia_lib/lib/model_public"
	"github.com/pdcgo/tokopedia_lib/lib/query"
)

func (api *TokopediaApiPublic) FilterSortProductQuery(payload *model_public.ParamsVar) (*model_public.FilterSortProductResp, error) {
	gqlQuery := GraphqlPayload{
		OperationName: "FilterSortProductQuery",
		Variables:     payload,
		Query:         query.FilterSortProductQuery,
	}
	req := api.NewGraphqlReq(&gqlQuery)

	var hasil model_public.FilterSortProductResp
	err := api.SendRequest(req, &hasil)
	return &hasil, err
}
