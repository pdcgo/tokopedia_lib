package api

import (
	"net/http"

	"github.com/pdcgo/tokopedia_lib/app/chat/config"
	"github.com/pdcgo/tokopedia_lib/app/chat/group"
	"github.com/pdcgo/tokopedia_lib/app/chat/model"
	"github.com/pdcgo/tokopedia_lib/app/chat/repo"
	tokpedapi "github.com/pdcgo/tokopedia_lib/lib/api"
)

type BaseApi struct{}

type BaseResponse struct {
	Code   int    `json:"code"`
	Detail string `json:"detail"`
}

func (a *BaseApi) BaseResponseSuccess() (int, BaseResponse) {
	return http.StatusOK, BaseResponse{
		Code:   http.StatusOK,
		Detail: "success",
	}
}

func (a *BaseApi) BaseResponseBadRequest(err error) (int, BaseResponse) {
	return http.StatusBadRequest, BaseResponse{
		Code:   http.StatusBadRequest,
		Detail: err.Error(),
	}
}

func (a *BaseApi) BaseResponseInternalServerError(err error) (int, BaseResponse) {
	return http.StatusInternalServerError, BaseResponse{
		Code:   http.StatusInternalServerError,
		Detail: err.Error(),
	}
}

type BaseDriverApi struct {
	BaseApi
	initConfig  *config.InitConfig
	accountRepo *repo.AccountRepo
	driverGroup *group.DriverGroup
}

func NewBaseDriverApi(
	initConfig *config.InitConfig,
	accountRepo *repo.AccountRepo,
	driverGroup *group.DriverGroup,
) *BaseDriverApi {

	return &BaseDriverApi{
		initConfig:  initConfig,
		accountRepo: accountRepo,
		driverGroup: driverGroup,
	}
}

type DriverApiHandler func(account *model.Account, driverApi *tokpedapi.TokopediaApi) error

func (api *BaseDriverApi) WithDriverApi(shopid int, handler DriverApiHandler) error {
	account, err := api.accountRepo.GetChatAccount(api.initConfig.ActiveGroup, shopid)
	if err != nil {
		return err
	}

	return api.driverGroup.WithDriverApi(account.GetUsername(), func(chatapi *tokpedapi.TokopediaApi) error {
		return handler(account, chatapi)
	})
}
