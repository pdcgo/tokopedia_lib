package api_public

import (
	"github.com/pdcgo/tokopedia_lib/lib/model_public"
	"github.com/pdcgo/tokopedia_lib/lib/query"
)

func (api *TokopediaApiPublic) SearchProductQueryV4(payload *model_public.ParamsVar) (*model_public.SearchProductQueryV4Resp, error) {
	gqlQuery := GraphqlPayload{
		OperationName: "SearchProductQueryV4",
		Variables:     payload,
		Query:         query.SearchProductQueryV4,
	}
	req := api.NewGraphqlReq(&gqlQuery)

	var hasil model_public.SearchProductQueryV4Resp
	err := api.SendRequest(req, &hasil)
	return &hasil, err
}

func (api *TokopediaApiPublic) TopadsProductQuery(payload *model_public.AdParamsVar) (*model_public.TopadsProductQueryResp, error) {
	gqlQuery := GraphqlPayload{
		OperationName: "TopadsProductQuery",
		Variables:     payload,
		Query:         query.TopadsProductQuery,
	}
	req := api.NewGraphqlReq(&gqlQuery)

	var hasil model_public.TopadsProductQueryResp
	err := api.SendRequest(req, &hasil)
	return &hasil, err
}
