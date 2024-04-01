package api

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/pdcgo/tokopedia_lib/app/chat/config"
	"github.com/pdcgo/tokopedia_lib/app/chat/group"
	"github.com/pdcgo/tokopedia_lib/app/chat/helper"
	"github.com/pdcgo/tokopedia_lib/app/chat/model"
	"github.com/pdcgo/tokopedia_lib/app/chat/repo"
	"github.com/pdcgo/tokopedia_lib/app/chat/service"
	tokpedapi "github.com/pdcgo/tokopedia_lib/lib/api"
	"github.com/pdcgo/tokopedia_lib/lib/chat"
	"github.com/pdcgo/v2_gots_sdk"
	"github.com/pdcgo/v2_gots_sdk/pdc_api"
)

type ChatApi struct {
	BaseApi
	initConfig          *config.InitConfig
	accountRepo         *repo.AccountRepo
	driverGroup         *group.DriverGroup
	sound               *helper.SoundPlayer
	chatService         *service.ChatService
	notificationService *service.NotificationService
}

func NewChatApi(
	initConfig *config.InitConfig,
	accountRepo *repo.AccountRepo,
	driverGroup *group.DriverGroup,
	sound *helper.SoundPlayer,
	chatService *service.ChatService,
	notificationService *service.NotificationService,
) *ChatApi {

	return &ChatApi{
		initConfig:          initConfig,
		accountRepo:         accountRepo,
		driverGroup:         driverGroup,
		sound:               sound,
		chatService:         chatService,
		notificationService: notificationService,
	}
}

func (api *ChatApi) users(ctx *gin.Context) {

	query := BaseQuery{}
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

	err = api.driverGroup.WithDriverApiByShopid(query.Shopid, func(username string, dapi *group.DriverApi) error {
		res, err := dapi.Api.GetChatList(payload)
		if err != nil {
			return err
		}

		ctx.JSON(http.StatusOK, res)
		return nil
	})
	if err != nil {
		ctx.JSON(api.BaseResponseInternalServerError(err))
	}
}

func (api *ChatApi) messages(ctx *gin.Context) {

	query := BaseQuery{}
	err := ctx.BindQuery(&query)
	if err != nil {
		ctx.JSON(api.BaseResponseBadRequest(err))
		return
	}

	payload := tokpedapi.ChatRoomVar{}
	err = ctx.BindJSON(&payload)
	if err != nil {
		ctx.JSON(api.BaseResponseBadRequest(err))
		return
	}

	err = api.driverGroup.WithDriverApiByShopid(query.Shopid, func(username string, dapi *group.DriverApi) error {
		res, err := dapi.Api.GetChatRoom(payload)
		if err != nil {
			return err
		}

		ctx.JSON(http.StatusOK, res)
		return nil
	})
	if err != nil {
		ctx.JSON(api.BaseResponseInternalServerError(err))
	}
}

type ChatReadQuery struct {
	*BaseQuery
	MessageId uint `json:"message_id" schema:"message_id" form:"message_id"`
}

func (api *ChatApi) read(ctx *gin.Context) {

	api.sound.Pause()

	query := ChatReadQuery{}
	err := ctx.BindQuery(&query)
	if err != nil {
		ctx.JSON(api.BaseResponseBadRequest(err))
		return
	}

	err = api.accountRepo.WithAccount(api.initConfig.ActiveGroup, query.Shopid, func(account *model.Account) error {
		err = api.chatService.ReadChat(account.GetUsername(), query.MessageId)
		if err != nil {
			return err
		}

		return api.notificationService.SendSyncAccountNotification(account)
	})
	if err != nil {
		ctx.JSON(api.BaseResponseInternalServerError(err))
		return
	}

	ctx.JSON(api.BaseResponseSuccess())
}

func (api *ChatApi) attachment(ctx *gin.Context) {

	query := BaseQuery{}
	err := ctx.BindQuery(&query)
	if err != nil {
		ctx.JSON(api.BaseResponseBadRequest(err))
		return
	}

	payload := tokpedapi.ChatAttachmentVar{}
	err = ctx.BindJSON(&payload)
	if err != nil {
		ctx.JSON(api.BaseResponseBadRequest(err))
		return
	}

	err = api.driverGroup.WithDriverApiByShopid(query.Shopid, func(username string, dapi *group.DriverApi) error {
		res, err := dapi.Api.GetChatAttachments(payload)
		if err != nil {
			return err
		}

		ctx.JSON(http.StatusOK, res)
		return nil
	})
	if err != nil {
		ctx.JSON(api.BaseResponseInternalServerError(err))
	}
}

type ChatPinQuery struct {
	*BaseQuery
	MessageId int64 `json:"msg_id" schema:"msg_id" form:"msg_id"`
}

func (api *ChatApi) pin(ctx *gin.Context) {

	query := ChatPinQuery{}
	err := ctx.BindQuery(&query)
	if err != nil {
		ctx.JSON(api.BaseResponseBadRequest(err))
		return
	}

	isUnpin := strings.Contains(ctx.Request.URL.Path, "unpin")
	err = api.driverGroup.WithDriverApiByShopid(query.Shopid, func(username string, dapi *group.DriverApi) error {

		if isUnpin {
			res, err := dapi.Api.ChatUnpin(query.MessageId)
			if err != nil {
				return err
			}
			ctx.JSON(http.StatusOK, res)

		} else {
			res, err := dapi.Api.ChatPin(query.MessageId)
			if err != nil {
				return err
			}
			ctx.JSON(http.StatusOK, res)
		}

		return nil
	})
	if err != nil {
		ctx.JSON(api.BaseResponseInternalServerError(err))
	}
}

func (api *ChatApi) userSearch(ctx *gin.Context) {

	query := BaseQuery{}
	err := ctx.BindQuery(&query)
	if err != nil {
		ctx.JSON(api.BaseResponseBadRequest(err))
		return
	}

	payload := tokpedapi.ChatSearchVar{
		IsSeller: 1,
		Status:   1,
	}
	err = ctx.BindJSON(&payload)
	if err != nil {
		ctx.JSON(api.BaseResponseBadRequest(err))
		return
	}

	err = api.driverGroup.WithDriverApiByShopid(query.Shopid, func(username string, dapi *group.DriverApi) error {
		res, err := dapi.Api.GetChatSearch(payload)
		if err != nil {
			return err
		}

		ctx.JSON(http.StatusOK, res)
		return nil
	})
	if err != nil {
		ctx.JSON(api.BaseResponseInternalServerError(err))
	}
}

func (api *ChatApi) send(ctx *gin.Context) {

	query := BaseQuery{}
	err := ctx.BindQuery(&query)
	if err != nil {
		ctx.JSON(api.BaseResponseBadRequest(err))
		return
	}

	payload := chat.SendChatPayload{}
	err = ctx.BindJSON(&payload)
	if err != nil {
		ctx.JSON(api.BaseResponseBadRequest(err))
		return
	}

	err = api.accountRepo.WithAccount(api.initConfig.ActiveGroup, query.Shopid, func(account *model.Account) error {
		data := payload.CreateEventData(account.ShopName)
		return api.chatService.SendChat(account.GetUsername(), data)
	})
	if err != nil {
		ctx.JSON(api.BaseResponseInternalServerError(err))
		return
	}

	ctx.JSON(api.BaseResponseSuccess())
}

func (api *ChatApi) Register(group *v2_gots_sdk.SdkGroup) {

	group.Register(&pdc_api.Api{
		Method:       http.MethodPost,
		RelativePath: "users",
		Query:        BaseQuery{},
		Payload:      tokpedapi.ChatListVar{},
		Response:     tokpedapi.ChatListRes{},
	}, api.users)

	group.Register(&pdc_api.Api{
		Method:       http.MethodPost,
		RelativePath: "messages",
		Query:        BaseQuery{},
		Payload:      tokpedapi.ChatRoomVar{},
		Response:     tokpedapi.ChatRoomRes{},
	}, api.messages)

	group.Register(&pdc_api.Api{
		Method:       http.MethodPost,
		RelativePath: "read",
		Query:        ChatReadQuery{},
		Response:     BaseResponse{},
	}, api.read)

	group.Register(&pdc_api.Api{
		Method:       http.MethodPost,
		RelativePath: "attachment",
		Query:        BaseQuery{},
		Payload:      tokpedapi.ChatAttachmentVar{},
		Response:     tokpedapi.ChatAttachmentRes{},
	}, api.attachment)

	group.Register(&pdc_api.Api{
		Method:       http.MethodPut,
		RelativePath: "pin",
		Query:        ChatPinQuery{},
		Payload:      tokpedapi.ChatPinVar{},
		Response:     tokpedapi.ChatPinRes{},
	}, api.pin)

	group.Register(&pdc_api.Api{
		Method:       http.MethodPut,
		RelativePath: "unpin",
		Query:        ChatPinQuery{},
		Payload:      tokpedapi.ChatPinVar{},
		Response:     tokpedapi.ChatUnpinRes{},
	}, api.pin)

	group.Register(&pdc_api.Api{
		Method:       http.MethodPost,
		RelativePath: "users/search",
		Query:        BaseQuery{},
		Payload:      tokpedapi.ChatSearchVar{},
		Response:     tokpedapi.ChatSearchRes{},
	}, api.userSearch)

	group.Register(&pdc_api.Api{
		Method:       http.MethodPost,
		RelativePath: "send",
		Query:        BaseQuery{},
		Payload:      chat.SendChatPayload{},
		Response:     tokpedapi.ChatSearchRes{},
	}, api.send)
}
