package api_public

import (
	"github.com/pdcgo/tokopedia_lib/lib/model"
	"github.com/pdcgo/tokopedia_lib/lib/query"
)

func (api *TokopediaApiPublic) ShopCoreInfo(payload *model.ShopCoreInfoVar) (*model.ShopCoreInfoResp, error) {
	gqlQuery := GraphqlPayload{
		OperationName: "ShopInfoCore",
		Variables:     payload,
		Query:         query.ShopCoreInfo,
	}

	req := api.NewGraphqlReq(&gqlQuery)

	var hasil model.ShopCoreInfoResp
	err := api.SendRequest(req, &hasil)
	return &hasil, err
}

func (api *TokopediaApiPublic) ShopStatisticQuery(payload *model.ShopStatisticQueryVar) (*model.ShopStatisticQueryResp, error) {
	gqlQuery := GraphqlPayload{
		OperationName: "ShopStatisticQuery",
		Variables:     payload,
		Query:         query.ShopStatisticQuery,
	}
	req := api.NewGraphqlReq(&gqlQuery)

	var hasil model.ShopStatisticQueryResp
	err := api.SendRequest(req, &hasil)
	return &hasil, err
}

func (api *TokopediaApiPublic) GetShopOperationalHourStatus(payload *model.ShopIdVar) (*model.GetShopOperationalHourStatusResp, error) {
	gqlQuery := GraphqlPayload{
		OperationName: "GetShopOperationalHourStatus",
		Variables:     payload,
		Query:         query.GetShopOperationalHourStatus,
	}
	req := api.NewGraphqlReq(&gqlQuery)

	var hasil model.GetShopOperationalHourStatusResp
	err := api.SendRequest(req, &hasil)
	return &hasil, err
}

func (api *TokopediaApiPublic) VoucherListQuery(payload *model.ShopIdVarInt) (*model.ShopVoucherQueryResp, error) {
	gqlQuery := GraphqlPayload{
		OperationName: "VoucherListQuery",
		Variables:     payload,
		Query:         query.VoucherListQuery,
	}
	req := api.NewGraphqlReq(&gqlQuery)

	var hasil model.ShopVoucherQueryResp
	err := api.SendRequest(req, &hasil)
	return &hasil, err
}

func (api *TokopediaApiPublic) ShopNote(payload *model.ShopNoteVar) (*model.ShopNoteResp, error) {
	gqlQuery := GraphqlPayload{
		OperationName: "ShopNote",
		Variables:     payload,
		Query:         query.ShopNote,
	}
	req := api.NewGraphqlReq(&gqlQuery)

	var hasil model.ShopNoteResp
	err := api.SendRequest(req, &hasil)
	return &hasil, err
}

func (api *TokopediaApiPublic) ShopPageLayoutV2(payload *model.ShopPageGetLayoutV2Var) (*model.ShopPageGetLayoutV2Resp, error) {
	gqlQuery := GraphqlPayload{
		OperationName: "ShopPageGetLayoutV2",
		Variables:     payload,
		Query:         query.ShopPageGetLayoutV2,
	}
	req := api.NewGraphqlReq(&gqlQuery)

	var hasil model.ShopPageGetLayoutV2Resp
	err := api.SendRequest(req, &hasil)
	return &hasil, err
}
