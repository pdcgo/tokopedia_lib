package api

import (
	"github.com/pdcgo/tokopedia_lib/lib/model"
	"github.com/pdcgo/tokopedia_lib/lib/query"
)

func (api *TokopediaApi) ChatGetGroupSticker(typevar int) (*model.ChatGetGroupStickerResp, error) {
	gqlQuery := GraphqlPayload{
		OperationName: "chatGetGroupSticker",
		Variables: model.TypeVar{
			Type: typevar,
		},
		Query: query.ChatGetGroupSticker,
	}
	req := api.NewGraphqlReq(&gqlQuery)

	var hasil model.ChatGetGroupStickerResp
	err := api.SendRequest(req, &hasil)
	return &hasil, err
}

func (api *TokopediaApi) ChatGetBundleSticker(payload *model.ChatGetBundleStickerVar) (*model.ChatGetBundleStickerResp, error) {
	gqlQuery := GraphqlPayload{
		OperationName: "chatGetBundleSticker",
		Variables:     payload,
		Query:         query.ChatGetBundleSticker,
	}

	req := api.NewGraphqlReq(&gqlQuery)

	var hasil model.ChatGetBundleStickerResp
	err := api.SendRequest(req, &hasil)
	return &hasil, err
}
