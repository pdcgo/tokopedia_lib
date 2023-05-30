package api_public

import (
	"github.com/pdcgo/tokopedia_lib/lib/model"
	"github.com/pdcgo/tokopedia_lib/lib/query"
)

func (api *TokopediaApiPublic) SearchProductQueryV4(payload *model.ParamsVar) (*model.SearchProductQueryV4Resp, error) {
	gqlQuery := GraphqlPayload{
		OperationName: "SearchProductQueryV4",
		Variables:     payload,
		Query:         query.SearchProductQueryV4,
	}
	req := api.NewGraphqlReq(&gqlQuery)

	var hasil model.SearchProductQueryV4Resp
	err := api.SendRequest(req, &hasil)
	return &hasil, err
}

func (api *TokopediaApiPublic) TopadsProductQuery(payload *model.AdParamsVar) (*model.TopadsProductQueryResp, error) {
	gqlQuery := GraphqlPayload{
		OperationName: "TopadsProductQuery",
		Variables:     payload,
		Query:         query.TopadsProductQuery,
	}
	req := api.NewGraphqlReq(&gqlQuery)

	var hasil model.TopadsProductQueryResp
	err := api.SendRequest(req, &hasil)
	return &hasil, err
}
