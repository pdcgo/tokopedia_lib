package api_public

import (
	"github.com/pdcgo/tokopedia_lib/lib/model_public"
	"github.com/pdcgo/tokopedia_lib/lib/query"
)

func (api *TokopediaApiPublic) CategoryAllListLite() (*model_public.CategoryAllListLiteResp, error) {
	var variable struct{}
	gqlQuery := GraphqlPayload{
		OperationName: "categoryAllListLite",
		Variables:     variable,
		Query:         query.CategoryAllListLite,
	}

	req := api.NewGraphqlReq(&gqlQuery)

	var hasil model_public.CategoryAllListLiteResp
	err := api.SendRequest(req, &hasil)
	return &hasil, err
}

// category recom
func (api *TokopediaApiPublic) JarvisRecommendation(payload *model_public.JarvisRecommendationVar) (*model_public.JarvisRecommendationResp, error) {
	gqlQuery := GraphqlPayload{
		OperationName: "jarvisRecommendation",
		Variables:     payload,
		Query:         query.JarvisRecommendation,
	}
	req := api.NewGraphqlReq(&gqlQuery)

	var hasil model_public.JarvisRecommendationResp
	err := api.SendRequest(req, &hasil)
	return &hasil, err
}
