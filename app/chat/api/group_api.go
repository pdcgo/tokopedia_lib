package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pdcgo/tokopedia_lib/app/chat/group"
	"github.com/pdcgo/tokopedia_lib/app/chat/model"
	"github.com/pdcgo/v2_gots_sdk"
	"gorm.io/gorm"
)

type GroupApi struct {
	BaseApi
	db        *gorm.DB
	chatGroup *group.ChatGroup
}

func NewGroupApi(db *gorm.DB, chatGroup *group.ChatGroup) *GroupApi {
	return &GroupApi{
		db:        db,
		chatGroup: chatGroup,
	}
}

func (api *GroupApi) list(ctx *gin.Context) {

	groups := []model.Group{}
	tx := api.db.Model(model.Group{}).Find(&groups)

	if tx.Error != nil {
		ctx.JSON(api.BaseResponseInternalServerError(tx.Error))
		return
	}

	res := []string{}
	for _, group := range groups {
		res = append(res, group.Name)
	}
	ctx.JSON(http.StatusOK, res)
}

func (api *GroupApi) connect(ctx *gin.Context) {

	groupName := ctx.Param("group_name")
	api.chatGroup.Connect(groupName)

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
}
