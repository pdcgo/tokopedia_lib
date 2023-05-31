package api_public

import (
	"github.com/pdcgo/tokopedia_lib/lib/model_public"
	"github.com/pdcgo/tokopedia_lib/lib/query"
)

func (api *TokopediaApiPublic) DiscussionDataByProductID(payload *model_public.DiscussionDataProductByIDVar) (*model_public.DiscussionDataProductByIDResp, error) {
	gqlQuery := GraphqlPayload{
		OperationName: "DiscussionDataByProductID",
		Variables:     payload,
		Query:         query.DiscussionDataByProductID,
	}
	req := api.NewGraphqlReq(&gqlQuery)

	var hasil model_public.DiscussionDataProductByIDResp
	err := api.SendRequest(req, &hasil)
	return &hasil, err
}
