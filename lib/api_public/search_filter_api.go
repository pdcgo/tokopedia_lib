package api_public

import (
	"github.com/pdcgo/tokopedia_lib/lib/model"
	"github.com/pdcgo/tokopedia_lib/lib/query"
)

func (api *TokopediaApiPublic) FilterSortProductQuery(payload *model.ParamsVar) (*model.FilterSortProductResp, error) {
	gqlQuery := GraphqlPayload{
		OperationName: "FilterSortProductQuery",
		Variables:     payload,
		Query:         query.FilterSortProductQuery,
	}
	req := api.NewGraphqlReq(&gqlQuery)

	var hasil model.FilterSortProductResp
	err := api.SendRequest(req, &hasil)
	return &hasil, err
}
