package api

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	socketio "github.com/googollee/go-socket.io"
	"github.com/pdcgo/tokopedia_lib/app/chat/config"
	"github.com/pdcgo/tokopedia_lib/app/chat/model"
	"github.com/pdcgo/tokopedia_lib/app/chat/repo"
	"github.com/pdcgo/tokopedia_lib/app/chat/service"
	apimodel "github.com/pdcgo/tokopedia_lib/lib/model"
	"github.com/pdcgo/v2_gots_sdk"
	"github.com/pdcgo/v2_gots_sdk/pdc_api"
)

type AccountApi struct {
	BaseApi
	sio            *socketio.Server
	accountService *service.AccountService
	initConfig     *config.InitConfig
}

func NewAccountApi(
	sio *socketio.Server,
	accountService *service.AccountService,
	initConfig *config.InitConfig,
) *AccountApi {

	return &AccountApi{
		sio:            sio,
		accountService: accountService,
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

	accounts, err := api.accountService.List(&query)
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

	res := AccountRes{}
	err = api.accountService.WithAccount(api.initConfig.ActiveGroup, shopid, func(account *model.Account) (err error) {
		res.Akun = account
		username := account.GetUsername()
		go api.accountService.OpenBrowser(username)

		res.Address, err = api.accountService.GetLocations(username)
		return
	})
	if err != nil {
		ctx.JSON(api.BaseResponseInternalServerError(err))
		return
	}

	ctx.JSON(http.StatusOK, res)
}

type AddAccountPayload struct {
	Akun    service.AccountPayload `json:"akun"`
	Name    string                 `json:"name"`
	Cookies []any                  `json:"cookies"`
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

func (api *AccountApi) edit(ctx *gin.Context) {

	payload := AddAccountPayload{}
	err := ctx.BindJSON(&payload)
	if err != nil {
		ctx.JSON(api.BaseResponseBadRequest(err))
		return
	}

	err = api.accountService.RemoveAccount(payload.Akun.Username)
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

func (api *AccountApi) remove(ctx *gin.Context) {

	username := ctx.Param("username")
	err := api.accountService.RemoveAccount(username)
	if err != nil {
		ctx.JSON(api.BaseResponseBadRequest(err))
		return
	}

	ctx.JSON(api.BaseResponseSuccess())
}

func (api *AccountApi) Register(group *v2_gots_sdk.SdkGroup) {

	group.Register(&pdc_api.Api{
		Method:       http.MethodGet,
		RelativePath: "",
		Query:        repo.ListAccountFilter{},
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
		RelativePath: "",
		Payload:      AddAccountPayload{},
		Response:     BaseResponse{},
	}, api.edit)

	group.Register(&pdc_api.Api{
		Method:       http.MethodDelete,
		RelativePath: "/:username",
		Response:     BaseResponse{},
	}, api.remove)

	group.Register(&pdc_api.Api{
		Method:       http.MethodPut,
		RelativePath: "toggle_pinned/:shopid",
		Response:     BaseResponse{},
	}, api.togglePin)

	group.Register(&pdc_api.Api{
		Method:       http.MethodPut,
		RelativePath: "/reconnect/:shopid",
		Response:     BaseResponse{},
	}, api.reconnect)

	group.Register(&pdc_api.Api{
		Method:       http.MethodPut,
		RelativePath: "/withdraw/:shopid",
		Response:     BaseResponse{},
	}, api.withdraw)

	group.Register(&pdc_api.Api{
		Method:       http.MethodPut,
		RelativePath: "/auto_withdraw",
		Response:     BaseResponse{},
	}, api.autoWithdraw)

	group.Register(&pdc_api.Api{
		Method:       http.MethodPut,
		RelativePath: "/set_pin",
		Query:        SetpinQuery{},
		Payload:      Setpinpayload{},
		Response:     BaseResponse{},
	}, api.setPin)
}
