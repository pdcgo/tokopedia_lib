package api

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/pdcgo/tokopedia_lib/app/chat/config"
	"github.com/pdcgo/tokopedia_lib/app/chat/group"
	"github.com/pdcgo/tokopedia_lib/app/chat/repo"
	"github.com/pdcgo/v2_gots_sdk"
	"gorm.io/gorm"
)

type GroupApi struct {
	BaseApi
	db          *gorm.DB
	initConfig  *config.InitConfig
	accountRepo *repo.AccountRepo
	groupRepo   *repo.GroupRepo
	chatGroup   *group.ChatGroup
}

func NewGroupApi(
	db *gorm.DB,
	initConfig *config.InitConfig,
	accountRepo *repo.AccountRepo,
	groupRepo *repo.GroupRepo,
	chatGroup *group.ChatGroup,
) *GroupApi {

	return &GroupApi{
		db:          db,
		initConfig:  initConfig,
		accountRepo: accountRepo,
		groupRepo:   groupRepo,
		chatGroup:   chatGroup,
	}
}

func (api *GroupApi) list(ctx *gin.Context) {

	list, err := api.groupRepo.GetList()
	if err != nil {
		ctx.JSON(api.BaseResponseInternalServerError(err))
		return
	}

	ctx.JSON(http.StatusOK, list)
}

func (api *GroupApi) connect(ctx *gin.Context) {

	groupName := ctx.Param("group_name")
	api.chatGroup.Connect(groupName)

	ctx.JSON(api.BaseResponseSuccess())
}

func (api *GroupApi) reconnect(ctx *gin.Context) {

	shopid, err := strconv.Atoi(ctx.Param("shopid"))
	if err != nil {
		ctx.JSON(api.BaseResponseBadRequest(err))
		return
	}

	err = api.chatGroup.Reconnect(shopid)
	if err != nil {
		ctx.JSON(api.BaseResponseInternalServerError(err))
		return
	}

	ctx.JSON(api.BaseResponseSuccess())
}

func (api *GroupApi) delete(ctx *gin.Context) {

	groupName := ctx.Param("group_name")
	err := api.groupRepo.Delete(groupName)
	if err != nil {
		ctx.JSON(api.BaseResponseInternalServerError(err))
		return
	}

	ctx.JSON(api.BaseResponseSuccess())
}

func (api *GroupApi) Register(group *v2_gots_sdk.SdkGroup) {

	group.Register(&v2_gots_sdk.Api{
		Method:       http.MethodGet,
		RelativePath: "",
		Response:     []string{},
	}, api.list)

	group.Register(&v2_gots_sdk.Api{
		Method:       http.MethodPut,
		RelativePath: ":group_name",
		Response:     BaseResponse{},
	}, api.connect)

	group.Register(&v2_gots_sdk.Api{
		Method:       http.MethodPut,
		RelativePath: "reconnect/:shopid",
		Response:     BaseResponse{},
	}, api.reconnect)

	group.Register(&v2_gots_sdk.Api{
		Method:       http.MethodDelete,
		RelativePath: ":group_name",
		Response:     BaseResponse{},
	}, api.delete)
}
