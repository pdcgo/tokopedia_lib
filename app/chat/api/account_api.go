package api

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/pdcgo/tokopedia_lib/app/chat/config"
	"github.com/pdcgo/tokopedia_lib/app/chat/group"
	"github.com/pdcgo/tokopedia_lib/app/chat/model"
	"github.com/pdcgo/tokopedia_lib/app/chat/repo"
	"github.com/pdcgo/tokopedia_lib/app/chat/service"
	tokpedapi "github.com/pdcgo/tokopedia_lib/lib/api"
	apimodel "github.com/pdcgo/tokopedia_lib/lib/model"
	"github.com/pdcgo/v2_gots_sdk"
	"github.com/pdcgo/v2_gots_sdk/pdc_api"
)

type AccountApi struct {
	BaseApi
	accountRepo    *repo.AccountRepo
	accountService *service.AccountService
	driverGroup    *group.DriverGroup
	initConfig     *config.InitConfig
}

func NewAccountApi(
	accountRepo *repo.AccountRepo,
	accountService *service.AccountService,
	driverGroup *group.DriverGroup,
	initConfig *config.InitConfig,
) *AccountApi {

	return &AccountApi{
		accountRepo:    accountRepo,
		accountService: accountService,
		driverGroup:    driverGroup,
		initConfig:     initConfig,
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

type AccountRes struct {
	Address []apimodel.ShopLocationLegacy `json:"address"`
	Akun    *model.Account                `json:"akun"`
}

func (api *AccountApi) get(ctx *gin.Context) {

	shopid, err := strconv.Atoi(ctx.Param("shopid"))
	if err != nil {
		ctx.JSON(api.BaseResponseBadRequest(err))
		return
	}

	account, err := api.accountRepo.GetChatAccount(api.initConfig.ActiveGroup, shopid)
	if err != nil {
		ctx.JSON(api.BaseResponseInternalServerError(err))
		return
	}

	username := account.GetUsername()
	api.driverGroup.WithDriverApi(username, func(tokpedapi *tokpedapi.TokopediaApi) error {

		locationAll, err := tokpedapi.GetShopLocationAll(shopid)
		if err != nil {
			ctx.JSON(api.BaseResponseInternalServerError(err))
			return err
		}

		locations := locationAll.Data.ShopLocGetAllLocations.Data.Warehouses.GetLocations()
		ctx.JSON(http.StatusOK, AccountRes{
			Address: locations,
			Akun:    account,
		})
		return nil
	})
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

func (api *AccountApi) togglePin(ctx *gin.Context) {

	shopid, err := strconv.Atoi(ctx.Param("shopid"))
	if err != nil {
		ctx.JSON(api.BaseResponseBadRequest(err))
		return
	}

	err = api.accountRepo.UpdateAccount(shopid, func(account *model.Account) {
		account.AccountData.Pinned = !account.AccountData.Pinned
	})
	if err != nil {
		ctx.JSON(api.BaseResponseInternalServerError(err))
		return
	}

	ctx.JSON(api.BaseResponseSuccess())
}

func (api *AccountApi) Register(group *v2_gots_sdk.SdkGroup) {

	group.Register(&pdc_api.Api{
		Method:       http.MethodGet,
		RelativePath: "",
		Payload:      repo.ListAccountFilter{},
		Response:     []model.Account{},
	}, api.list)

	group.Register(&pdc_api.Api{
		Method:       http.MethodGet,
		RelativePath: ":shopid",
		Response:     []AccountRes{},
	}, api.get)

	group.Register(&pdc_api.Api{
		Method:       http.MethodPost,
		RelativePath: "",
		Payload:      AddAccountPayload{},
		Response:     BaseResponse{},
	}, api.add)

	group.Register(&pdc_api.Api{
		Method:       http.MethodPut,
		RelativePath: "toggle_pinned/:shopid",
		Response:     BaseResponse{},
	}, api.togglePin)
}
