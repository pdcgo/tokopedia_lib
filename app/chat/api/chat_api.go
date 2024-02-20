package api

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/pdcgo/tokopedia_lib/app/chat/helper"
	"github.com/pdcgo/tokopedia_lib/app/chat/model"
	"github.com/pdcgo/tokopedia_lib/app/chat/service"
	tokpedapi "github.com/pdcgo/tokopedia_lib/lib/api"
	"github.com/pdcgo/v2_gots_sdk"
	"github.com/pdcgo/v2_gots_sdk/pdc_api"
)

type ChatApi struct {
	*BaseDriverApi
	sound               *helper.SoundPlayer
	chatService         *service.ChatService
	notificationService *service.NotificationService
}

func NewChatApi(
	driverApi *BaseDriverApi,
	sound *helper.SoundPlayer,
	chatService *service.ChatService,
	notificationService *service.NotificationService,
) *ChatApi {

	return &ChatApi{
		BaseDriverApi:       driverApi,
		sound:               sound,
		chatService:         chatService,
		notificationService: notificationService,
	}
}

type ChatQuery struct {
	Shopid int `json:"shopid" schema:"shopid" form:"shopid"`
}

func (api *ChatApi) users(ctx *gin.Context) {

	query := ChatQuery{}
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

	err = api.WithDriverApi(query.Shopid, func(account *model.Account, driverApi *tokpedapi.TokopediaApi) error {
		res, err := driverApi.GetChatList(payload)
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

	query := ChatQuery{}
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

	err = api.WithDriverApi(query.Shopid, func(account *model.Account, driverApi *tokpedapi.TokopediaApi) error {
		res, err := driverApi.GetChatRoom(payload)
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
	ChatQuery
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

	account, err := api.accountRepo.GetChatAccount(api.initConfig.ActiveGroup, query.Shopid)
	if err != nil {
		ctx.JSON(api.BaseResponseInternalServerError(err))
		return
	}

	err = api.chatService.ReadChat(account.GetUsername(), query.MessageId)
	if err != nil {
		ctx.JSON(api.BaseResponseInternalServerError(err))
		return
	}

	err = api.notificationService.SendSyncAccountNotification(account)
	if err != nil {
		ctx.JSON(api.BaseResponseInternalServerError(err))
		return
	}

	ctx.JSON(api.BaseResponseSuccess())
}

func (api *ChatApi) attachment(ctx *gin.Context) {

	query := ChatQuery{}
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

	err = api.WithDriverApi(query.Shopid, func(account *model.Account, driverApi *tokpedapi.TokopediaApi) error {
		res, err := driverApi.GetChatAttachments(payload)
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
	ChatQuery
	MessageId int64 `json:"msg_id" schema:"msg_id" form:"msg_id"`
}

func (api *ChatApi) pin(ctx *gin.Context) {

	query := ChatPinQuery{}
	err := ctx.BindQuery(&query)
	if err != nil {
		ctx.JSON(api.BaseResponseBadRequest(err))
		return
	}

	err = api.WithDriverApi(query.Shopid, func(account *model.Account, driverApi *tokpedapi.TokopediaApi) error {
		isUnpin := strings.Contains(ctx.Request.URL.Path, "unpin")

		if isUnpin {
			res, err := driverApi.ChatUnpin(query.MessageId)
			if err != nil {
				return err
			}
			ctx.JSON(http.StatusOK, res)

		} else {
			res, err := driverApi.ChatPin(query.MessageId)
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

	query := ChatQuery{}
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

	err = api.WithDriverApi(query.Shopid, func(account *model.Account, driverApi *tokpedapi.TokopediaApi) error {
		res, err := driverApi.GetChatSearch(payload)
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

	query := ChatQuery{}
	err := ctx.BindQuery(&query)
	if err != nil {
		ctx.JSON(api.BaseResponseBadRequest(err))
		return
	}

	payload := service.SendChat{}
	err = ctx.BindJSON(&payload)
	if err != nil {
		ctx.JSON(api.BaseResponseBadRequest(err))
		return
	}

	account, err := api.accountRepo.GetChatAccount(api.initConfig.ActiveGroup, query.Shopid)
	if err != nil {
		ctx.JSON(api.BaseResponseInternalServerError(err))
		return
	}

	err = api.chatService.SendChat(account.GetUsername(), account.ShopName, &payload)
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
		Query:        ChatQuery{},
		Payload:      tokpedapi.ChatListVar{},
		Response:     tokpedapi.ChatListRes{},
	}, api.users)

	group.Register(&pdc_api.Api{
		Method:       http.MethodPost,
		RelativePath: "messages",
		Query:        ChatQuery{},
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
		Query:        ChatQuery{},
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
		Query:        ChatQuery{},
		Payload:      tokpedapi.ChatSearchVar{},
		Response:     tokpedapi.ChatSearchRes{},
	}, api.userSearch)

	group.Register(&pdc_api.Api{
		Method:       http.MethodPost,
		RelativePath: "send",
		Query:        ChatQuery{},
		Payload:      service.SendChat{},
		Response:     tokpedapi.ChatSearchRes{},
	}, api.send)
}
