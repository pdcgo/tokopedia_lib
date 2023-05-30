package api_public

import (
	"github.com/pdcgo/tokopedia_lib/lib/model_public"
	"github.com/pdcgo/tokopedia_lib/lib/query"
)

func (api *TokopediaApiPublic) ShopCoreInfo(payload *model_public.ShopCoreInfoVar) (*model_public.ShopCoreInfoResp, error) {
	gqlQuery := GraphqlPayload{
		OperationName: "ShopInfoCore",
		Variables:     payload,
		Query:         query.ShopCoreInfo,
	}

	req := api.NewGraphqlReq(&gqlQuery)

	var hasil model_public.ShopCoreInfoResp
	err := api.SendRequest(req, &hasil)
	return &hasil, err
}

func (api *TokopediaApiPublic) ShopStatisticQuery(payload *model_public.ShopStatisticQueryVar) (*model_public.ShopStatisticQueryResp, error) {
	gqlQuery := GraphqlPayload{
		OperationName: "ShopStatisticQuery",
		Variables:     payload,
		Query:         query.ShopStatisticQuery,
	}
	req := api.NewGraphqlReq(&gqlQuery)

	var hasil model_public.ShopStatisticQueryResp
	err := api.SendRequest(req, &hasil)
	return &hasil, err
}

func (api *TokopediaApiPublic) GetShopOperationalHourStatus(payload *model_public.ShopIdVar) (*model_public.GetShopOperationalHourStatusResp, error) {
	gqlQuery := GraphqlPayload{
		OperationName: "GetShopOperationalHourStatus",
		Variables:     payload,
		Query:         query.GetShopOperationalHourStatus,
	}
	req := api.NewGraphqlReq(&gqlQuery)

	var hasil model_public.GetShopOperationalHourStatusResp
	err := api.SendRequest(req, &hasil)
	return &hasil, err
}

func (api *TokopediaApiPublic) VoucherListQuery(payload *model_public.ShopIdVarInt) (*model_public.ShopVoucherQueryResp, error) {
	gqlQuery := GraphqlPayload{
		OperationName: "VoucherListQuery",
		Variables:     payload,
		Query:         query.VoucherListQuery,
	}
	req := api.NewGraphqlReq(&gqlQuery)

	var hasil model_public.ShopVoucherQueryResp
	err := api.SendRequest(req, &hasil)
	return &hasil, err
}

func (api *TokopediaApiPublic) ShopNote(payload *model_public.ShopNoteVar) (*model_public.ShopNoteResp, error) {
	gqlQuery := GraphqlPayload{
		OperationName: "ShopNote",
		Variables:     payload,
		Query:         query.ShopNote,
	}
	req := api.NewGraphqlReq(&gqlQuery)

	var hasil model_public.ShopNoteResp
	err := api.SendRequest(req, &hasil)
	return &hasil, err
}

func (api *TokopediaApiPublic) ShopPageLayoutV2(payload *model_public.ShopPageGetLayoutV2Var) (*model_public.ShopPageGetLayoutV2Resp, error) {
	gqlQuery := GraphqlPayload{
		OperationName: "ShopPageGetLayoutV2",
		Variables:     payload,
		Query:         query.ShopPageGetLayoutV2,
	}
	req := api.NewGraphqlReq(&gqlQuery)

	var hasil model_public.ShopPageGetLayoutV2Resp
	err := api.SendRequest(req, &hasil)
	return &hasil, err
}
