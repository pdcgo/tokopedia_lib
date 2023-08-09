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
func (api *TokopediaApiPublic) JarvisRecommendation(prodname string) (*model_public.JarvisRecommendationResp, error) {
	variable := model_public.JarvisRecommendationVar{
		ProductName: prodname,
	}
	gqlQuery := GraphqlPayload{
		OperationName: "jarvisRecommendation",
		Variables:     variable,
		Query:         query.JarvisRecommendation,
	}
	req := api.NewGraphqlReq(&gqlQuery)

	var hasil model_public.JarvisRecommendationResp
	err := api.SendRequest(req, &hasil)
	return &hasil, err
}

func (api *TokopediaApiPublic) HeaderMainData() (*model_public.HeaderMainDataResp, error) {
	var variable struct{}

	gqlQuery := GraphqlPayload{
		OperationName: "headerMainData",
		Variables:     variable,
		Query:         query.HeaderMainData,
	}

	req := api.NewGraphqlReq(&gqlQuery)

	var hasil model_public.HeaderMainDataResp
	err := api.SendRequest(req, &hasil)
	return &hasil, err
}
