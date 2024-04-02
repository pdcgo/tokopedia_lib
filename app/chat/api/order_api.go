package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	chat_model "github.com/pdcgo/tokopedia_lib/app/chat/model"
	"github.com/pdcgo/tokopedia_lib/app/chat/repo"
	"github.com/pdcgo/tokopedia_lib/app/chat/service"
	"github.com/pdcgo/tokopedia_lib/lib/model"
	"github.com/pdcgo/tokopedia_lib/lib/query"
	"github.com/pdcgo/v2_gots_sdk"
	"github.com/pdcgo/v2_gots_sdk/pdc_api"
)

type OrderApi struct {
	BaseApi
	orderRepo    *repo.OrderRepo
	orderService *service.OrderService
}

func NewOrderApi(orderRepo *repo.OrderRepo, orderService *service.OrderService) *OrderApi {

	return &OrderApi{
		orderRepo:    orderRepo,
		orderService: orderService,
	}
}

func (api *OrderApi) list(ctx *gin.Context) {

	query := repo.ListOrderFilter{
		SortType: "desc",
		SortBy:   "created",
		TypeDate: "created",
	}
	err := ctx.BindQuery(&query)
	if err != nil {
		ctx.JSON(api.BaseResponseBadRequest(err))
		return
	}

	res, err := api.orderRepo.Paginate(&query)
	if err != nil {
		ctx.JSON(api.BaseResponseInternalServerError(err))
		return
	}

	ctx.JSON(http.StatusOK, res)
}

func (api *OrderApi) unpaid(ctx *gin.Context) {

	q := BaseQuery{}
	err := ctx.BindQuery(&q)
	if err != nil {
		ctx.JSON(api.BaseResponseBadRequest(err))
		return
	}

	payload := query.OrderPendingListQuery{}
	err = ctx.BindJSON(&payload)
	if err != nil {
		ctx.JSON(api.BaseResponseBadRequest(err))
		return
	}

	res, err := api.orderService.GetUnpaid(q.Shopid, &payload)
	if err != nil {
		ctx.JSON(api.BaseResponseInternalServerError(err))
		return
	}

	ctx.JSON(http.StatusOK, res)
}

func (api *OrderApi) Register(group *v2_gots_sdk.SdkGroup) {

	group.Register(&pdc_api.Api{
		Method:       http.MethodGet,
		RelativePath: "",
		Query:        repo.ListOrderFilter{},
		Response:     repo.PaginationResult[chat_model.Order]{},
	}, api.list)

	group.Register(&pdc_api.Api{
		Method:       http.MethodPost,
		RelativePath: "unpaid",
		Query:        BaseQuery{},
		Payload:      query.OrderPendingListQuery{},
		Response:     model.OrderListWaitingPaymentRes{},
	}, api.unpaid)
}
