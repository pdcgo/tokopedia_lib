package api

import (
	"github.com/pdcgo/tokopedia_lib/lib/model"
	"github.com/pdcgo/tokopedia_lib/lib/query"
)

func (api *TokopediaApi) CategoryAllListLite() (*model.CategoryAllListLiteResp, error) {
	var variable interface{}
	gqlQuery := GraphqlPayload{
		OperationName: "categoryAllListLite",
		Variables:     variable,
		Query:         query.CategoryAllListLite,
	}

	req := api.NewGraphqlReq(&gqlQuery)

	var hasil []*model.CategoryAllListLiteResp
	err := api.SendRequest(req, &hasil)
	return hasil[0], err
}

// category recom
func (api *TokopediaApi) JarvisRecommendation(payload *model.JarvisRecommendationVar) (*model.JarvisRecommendationResp, error) {
	gqlQuery := GraphqlPayload{
		OperationName: "jarvisRecommendation",
		Variables:     payload,
		Query:         query.JarvisRecommendation,
	}
	req := api.NewGraphqlReq(&gqlQuery)

	var hasil []*model.JarvisRecommendationResp
	err := api.SendRequest(req, &hasil)
	return hasil[0], err
}
