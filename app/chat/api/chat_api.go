package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pdcgo/tokopedia_lib/app/chat/config"
	"github.com/pdcgo/tokopedia_lib/app/chat/group"
	"github.com/pdcgo/tokopedia_lib/app/chat/repo"
	tokpedapi "github.com/pdcgo/tokopedia_lib/lib/api"
	"github.com/pdcgo/v2_gots_sdk"
)

type ChatApi struct {
	BaseApi
	initConfig  *config.InitConfig
	accountRepo *repo.AccountRepo
	driverGroup *group.DriverGroup
}

func NewChatApi(initConfig *config.InitConfig, accountRepo *repo.AccountRepo, driverGroup *group.DriverGroup) *ChatApi {
	return &ChatApi{
		initConfig:  initConfig,
		accountRepo: accountRepo,
		driverGroup: driverGroup,
	}
}

type ChatUsersQuery struct {
	Shopid int `json:"shopid" schema:"shopid" form:"shopid"`
}

func (api *ChatApi) users(ctx *gin.Context) {

	query := ChatUsersQuery{}
	err := ctx.BindQuery(&query)
	if err != nil {
		ctx.JSON(api.BaseResponseBadRequest(err))
		return
	}

	payload := tokpedapi.ChatListVar{}
	err = ctx.BindJSON(&payload)
	if err != nil {
		ctx.JSON(api.BaseResponseBadRequest(err))
		return
	}

	accountData, err := api.accountRepo.GetChatAccountData(api.initConfig.ActiveGroup, query.Shopid)
	if err != nil {
		ctx.JSON(api.BaseResponseInternalServerError(err))
		return
	}

	api.driverGroup.WithDriverApi(accountData.Username, func(tokpedapi *tokpedapi.TokopediaApi) error {

		res, err := tokpedapi.GetChatList(payload)
		if err != nil {
			ctx.JSON(api.BaseResponseInternalServerError(err))
			return err
		}

		ctx.JSON(http.StatusOK, res)
		return nil
	})
}

func (api *ChatApi) Register(group *v2_gots_sdk.SdkGroup) {

	group.Register(&v2_gots_sdk.Api{
		Method:       http.MethodPost,
		RelativePath: "users",
		Query:        ChatUsersQuery{},
		Payload:      tokpedapi.ChatListVar{},
		Response:     tokpedapi.ChatListRes{},
	}, api.users)
}
