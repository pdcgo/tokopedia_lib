package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pdcgo/tokopedia_lib/app/chat/model"
	"github.com/pdcgo/tokopedia_lib/app/chat/repo"
	"github.com/pdcgo/tokopedia_lib/app/chat/service"
	"github.com/pdcgo/v2_gots_sdk"
)

type AccountApi struct {
	BaseApi
	accountRepo    *repo.AccountRepo
	accountService *service.AccountService
}

func NewAccountApi(accountRepo *repo.AccountRepo, accountService *service.AccountService) *AccountApi {
	return &AccountApi{
		accountRepo:    accountRepo,
		accountService: accountService,
	}
}

func (api *AccountApi) list(ctx *gin.Context) {

	query := repo.ListAccountFilter{}
	err := ctx.BindQuery(&query)
	if err != nil {
		ctx.JSON(api.BaseResponseBadRequest(err))
		return
	}

	accounts, err := api.accountRepo.List(&query)
	if err != nil {
		ctx.JSON(api.BaseResponseInternalServerError(err))
		return
	}

	ctx.JSON(http.StatusOK, accounts)
}

type AddAccountPayload struct {
	Akun    service.Account `json:"akun"`
	Name    string          `json:"name"`
	Cookies []any           `json:"cookies"`
}

func (api *AccountApi) add(ctx *gin.Context) {

	payload := AddAccountPayload{}
	err := ctx.BindJSON(&payload)
	if err != nil {
		ctx.JSON(api.BaseResponseBadRequest(err))
		return
	}

	err = api.accountService.AddAccount(payload.Akun, payload.Name)
	if err != nil {
		ctx.JSON(api.BaseResponseInternalServerError(err))
		return
	}

	ctx.JSON(api.BaseResponseSuccess())
}

func (api *AccountApi) Register(group *v2_gots_sdk.SdkGroup) {

	group.Register(&v2_gots_sdk.Api{
		Method:       http.MethodGet,
		RelativePath: "",
		Payload:      repo.ListAccountFilter{},
		Response:     []model.Account{},
	}, api.list)

	group.Register(&v2_gots_sdk.Api{
		Method:       http.MethodPost,
		RelativePath: "",
		Payload:      AddAccountPayload{},
		Response:     BaseResponse{},
	}, api.add)
}
