package api_public

import (
	"github.com/pdcgo/tokopedia_lib/lib/model_public"
	"github.com/pdcgo/tokopedia_lib/lib/query"
)

func (api *TokopediaApiPublic) LocDisctricRecommendation(payload *model_public.LocDisctricRecommendationVar) (*model_public.LocDisctricRecommendationResp, error) {
	gqlQuery := GraphqlPayload{
		OperationName: "keroDistrictRecommendation",
		Variables:     payload,
		Query:         query.LocDisctricRecommendation,
	}
	req := api.NewGraphqlReq(&gqlQuery)

	var hasil model_public.LocDisctricRecommendationResp
	err := api.SendRequest(req, &hasil)
	return &hasil, err
}
