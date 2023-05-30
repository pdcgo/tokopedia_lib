package api_public

import (
	"github.com/pdcgo/tokopedia_lib/lib/model"
	"github.com/pdcgo/tokopedia_lib/lib/query"
)

func (api *TokopediaApiPublic) PdpGetlayoutQuery(payload *model.PdpGetlayoutQueryVar) (*model.PdpGetlayoutQueryResp, error) {
	gqlQuery := GraphqlPayload{
		OperationName: "PDPGetLayoutQuery",
		Variables:     payload,
		Query:         query.PdpGetLayoutQuery,
	}

	req := api.NewGraphqlReq(&gqlQuery)

	headers := map[string]string{
		"X-Tkpd-Akamai": "pdpGetLayout",
		"X-Device":      "desktop",
	}
	for key, val := range headers {
		req.Header.Set(key, val)
	}

	var hasil model.PdpGetlayoutQueryResp
	err := api.SendRequest(req, &hasil)
	return &hasil, err
}

func (api *TokopediaApiPublic) PdpGetDataP2(payload *model.PdpGetDataP2Var) (*model.PdpGetDataP2Resp, error) {
	gqlQuery := GraphqlPayload{
		OperationName: "PDPGetDataP2",
		Variables:     payload,
		Query:         query.PdpGetDataP2,
	}
	req := api.NewGraphqlReq(&gqlQuery)

	var hasil model.PdpGetDataP2Resp
	err := api.SendRequest(req, &hasil)
	return &hasil, err
}

func (api *TokopediaApiPublic) PdpShopNote(payload *model.ShopIdVar) (*model.PdpShopNoteResp, error) {
	gqlQuery := GraphqlPayload{
		OperationName: "PDPShopNote",
		Variables:     payload,
		Query:         query.PdpShopNote,
	}

	req := api.NewGraphqlReq(&gqlQuery)

	var hasil model.PdpShopNoteResp
	err := api.SendRequest(req, &hasil)
	return &hasil, err
}

func (api *TokopediaApiPublic) ProductRatingandTopics(payload *model.ProductIdVar) (*model.ProductRatingandTopicsResp, error) {
	gqlQuery := GraphqlPayload{
		OperationName: "productRatingAndTopics",
		Variables:     payload,
		Query:         query.ProductRatingandTopics,
	}

	req := api.NewGraphqlReq(&gqlQuery)

	var hasil model.ProductRatingandTopicsResp
	err := api.SendRequest(req, &hasil)
	return &hasil, err
}

func (api *TokopediaApiPublic) PdpGetReiewImageQuery(payload *model.PdpGetReiewImageQueryVar) (*model.PdpGetReiewImageQueryResp, error) {
	gqlQuery := GraphqlPayload{
		OperationName: "PDPGetReviewImageQuery",
		Variables:     payload,
		Query:         query.PdpGetReiewImageQuery,
	}
	req := api.NewGraphqlReq(&gqlQuery)

	var hasil model.PdpGetReiewImageQueryResp
	err := api.SendRequest(req, &hasil)
	return &hasil, err

}

func (api *TokopediaApiPublic) ProductReviewList(payload *model.ProductReviewListVar) (*model.ProductReviewListResp, error) {
	gqlQuery := GraphqlPayload{
		OperationName: "productReviewList",
		Variables:     payload,
		Query:         query.ProductReviewList,
	}
	req := api.NewGraphqlReq(&gqlQuery)

	var hasil model.ProductReviewListResp
	err := api.SendRequest(req, &hasil)
	return &hasil, err

}

func (api *TokopediaApiPublic) RecomWidget(payload *model.RecomWidgetVar) (*model.RecomWidgetResp, error) {
	gqlQuery := GraphqlPayload{
		OperationName: "RecomWidget",
		Variables:     payload,
		Query:         query.RecomWidget,
	}
	req := api.NewGraphqlReq(&gqlQuery)

	var hasil model.RecomWidgetResp
	err := api.SendRequest(req, &hasil)
	return &hasil, err
}
