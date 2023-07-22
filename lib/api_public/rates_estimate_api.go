package api_public

import (
	"github.com/pdcgo/tokopedia_lib/lib/model_public"
	"github.com/pdcgo/tokopedia_lib/lib/query"
)

func (api *TokopediaApiPublic) RatesEstimateQuery(payload *model_public.RatesEstimateQueryVar) (*model_public.RatesEstimateQueryResp, error) {
	gqlQuery := GraphqlPayload{
		OperationName: "ratesEstimateQuery",
		Variables:     payload,
		Query:         query.RatesEstimateQuery,
	}
	req := api.NewGraphqlReq(&gqlQuery)

	var hasil model_public.RatesEstimateQueryResp
	err := api.SendRequest(req, &hasil)
	return &hasil, err
}
