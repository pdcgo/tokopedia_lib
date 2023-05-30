package api_public

import (
	"github.com/pdcgo/tokopedia_lib/lib/model"
	"github.com/pdcgo/tokopedia_lib/lib/query"
)

func (api *TokopediaApiPublic) CategoryAllListLite() (*model.CategoryAllListLiteResp, error) {
	var variable struct{}
	gqlQuery := GraphqlPayload{
		OperationName: "categoryAllListLite",
		Variables:     variable,
		Query:         query.CategoryAllListLite,
	}

	req := api.NewGraphqlReq(&gqlQuery)

	var hasil model.CategoryAllListLiteResp
	err := api.SendRequest(req, &hasil)
	return &hasil, err
}

// category recom
func (api *TokopediaApiPublic) JarvisRecommendation(payload *model.JarvisRecommendationVar) (*model.JarvisRecommendationResp, error) {
	gqlQuery := GraphqlPayload{
		OperationName: "jarvisRecommendation",
		Variables:     payload,
		Query:         query.JarvisRecommendation,
	}
	req := api.NewGraphqlReq(&gqlQuery)

	var hasil model.JarvisRecommendationResp
	err := api.SendRequest(req, &hasil)
	return &hasil, err
}
