package api_public

import (
	"github.com/pdcgo/tokopedia_lib/lib/model_public"
	"github.com/pdcgo/tokopedia_lib/lib/query"
)

func (api *TokopediaApiPublic) VariantCombinationQuery(payload *model_public.VariantCategoryCombinationVar) (*model_public.VariantCategoryCombinationResp, error) {
	gqlQuery := GraphqlPayload{
		OperationName: "variantCategoryCombination",
		Variables:     payload,
		Query:         query.VariantCategoryCombination,
	}
	req := api.NewGraphqlReq(&gqlQuery)

	var hasil model_public.VariantCategoryCombinationResp
	err := api.SendRequest(req, &hasil)
	return &hasil, err
}
