package api_public

import (
	"github.com/pdcgo/tokopedia_lib/lib/model_public"
	"github.com/pdcgo/tokopedia_lib/lib/query"
)

func (api *TokopediaApiPublic) GetDrogonAnnotation(payload *model_public.GetDrogonAnnotationVar) (*model_public.GetDrogonAnnotationResp, error) {
	gqlQuery := GraphqlPayload{
		OperationName: "getDrogonAnnotation",
		Variables:     payload,
		Query:         query.GetDrogonAnnotation,
	}
	req := api.NewGraphqlReq(&gqlQuery)

	var hasil model_public.GetDrogonAnnotationResp
	err := api.SendRequest(req, &hasil)
	return &hasil, err
}
