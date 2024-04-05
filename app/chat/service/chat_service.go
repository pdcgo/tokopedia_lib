package service

import (
	"log"
	"time"

	"github.com/pdcgo/common_conf/common_concept"
	"github.com/pdcgo/common_conf/pdc_common"
	"github.com/pdcgo/tokopedia_lib/app/chat/config"
	"github.com/pdcgo/tokopedia_lib/app/chat/group"
	"github.com/pdcgo/tokopedia_lib/app/chat/helper"
	"github.com/pdcgo/tokopedia_lib/app/chat/model"
	"github.com/pdcgo/tokopedia_lib/app/chat/repo"
	"github.com/pdcgo/tokopedia_lib/app/chat/sio_event"
	"github.com/pdcgo/tokopedia_lib/lib/api"
	"github.com/pdcgo/tokopedia_lib/lib/chat"
)

type ChatService struct {
	event        *common_concept.CoreEvent
	initConfig   *config.InitConfig
	areplyConfig *config.AutoReplyConfig
	accountRepo  *repo.AccountRepo
	socketGroup  *group.SocketGroup
	driverGroup  *group.DriverGroup
	sound        *helper.SoundPlayer
}

func NewChatService(
	event *common_concept.CoreEvent,
	initConfig *config.InitConfig,
	areplyConfig *config.AutoReplyConfig,
	accountRepo *repo.AccountRepo,
	socketGroup *group.SocketGroup,
	driverGroup *group.DriverGroup,
	sound *helper.SoundPlayer,
) *ChatService {

	chatService := ChatService{
		event:        event,
		initConfig:   initConfig,
		areplyConfig: areplyConfig,
		accountRepo:  accountRepo,
		socketGroup:  socketGroup,
		driverGroup:  driverGroup,
		sound:        sound,
	}

	go chatService.handleEvent()
	return &chatService
}

func (s *ChatService) ReadChat(shopid int, msgId uint) error {

	readEvent := chat.BaseSocketType{
		Code: chat.ReadUserChatEvent,
		Data: chat.MessageId{
			MsgId: msgId,
		},
	}

	return s.socketGroup.WithSocketByShopid(shopid, func(username string, sc *chat.SocketClient) error {
		return sc.SendEvent(readEvent)
	})
}

func (s *ChatService) SendChat(shopid int, payload chat.SendChatPayload) error {
	return s.socketGroup.WithSocketByShopid(shopid, func(username string, sc *chat.SocketClient) error {

		data := payload.CreateEventData(username)

		log.Printf("[ %s ] send message attach:%d", username, data.AttachmentType)
		return sc.SendEvent(&chat.EmitEventSocket{
			EventCode: &chat.EventCode{
				Code: chat.ChatEvent,
			},
			Data: data,
		})
	})
}

func (s *ChatService) Pin(shopid int, pin bool, msgid int64) (any, error) {
	var res any
	err := s.driverGroup.WithDriverApiByShopid(shopid, func(username string, dapi *group.DriverApi) (err error) {
		if pin {
			res, err = dapi.Api.ChatPin(msgid)
		} else {
			res, err = dapi.Api.ChatUnpin(msgid)
		}
		return
	})

	return res, err
}

func (s *ChatService) GetChatSearch(shopid int, payload api.ChatSearchVar) (res *api.ChatSearchRes, err error) {
	err = s.driverGroup.WithDriverApiByShopid(shopid, func(username string, dapi *group.DriverApi) error {
		res, err = dapi.Api.GetChatSearch(payload)
		return err
	})

	return
}

func (s *ChatService) GetChatList(shopid int, payload api.ChatListVar) (res *api.ChatListRes, err error) {
	err = s.driverGroup.WithDriverApiByShopid(shopid, func(username string, dapi *group.DriverApi) error {
		res, err = dapi.Api.GetChatList(payload)
		return err
	})

	return
}

func (s *ChatService) GetChatRoom(shopid int, payload api.ChatRoomVar) (res *api.ChatRoomRes, err error) {
	err = s.driverGroup.WithDriverApiByShopid(shopid, func(username string, dapi *group.DriverApi) error {
		res, err = dapi.Api.GetChatRoom(payload)
		return err
	})

	return
}

func (s *ChatService) GetChatAttachments(shopid int, payload api.ChatAttachmentVar) (res *api.ChatAttachmentRes, err error) {
	err = s.driverGroup.WithDriverApiByShopid(shopid, func(username string, dapi *group.DriverApi) error {
		res, err = dapi.Api.GetChatAttachments(payload)
		return err
	})

	return
}

func (s *ChatService) autoReply(shopid int, data *chat.RcvChat) {
	if s.areplyConfig.Data.Active {

		reply := s.areplyConfig.Find(data.Message.OriginalReply)
		if reply != nil {
			err := s.accountRepo.WithAccount(s.initConfig.ActiveGroup, shopid, func(account *model.Account) error {
				delay := reply.GetDelay()
				log.Printf("[ %s ] auto replying for %ds...", account.AccountData.Username, int(delay.Seconds()))
				time.Sleep(delay)

				sendpayload := chat.SendChatPayload{
					MessageId: data.MsgID,
					Message:   reply.Reply,
				}
				return s.SendChat(shopid, sendpayload)
			})
			if err != nil {
				pdc_common.ReportError(err)
				return
			}
		}
	}
}

func (s *ChatService) handleEvent() {
	for event := range s.event.GetEvent() {
		switch ev := event.(type) {

		case *sio_event.SendChatEvent:
			if ev.Event.IsOpposite {
				s.sound.Play()
				s.autoReply(ev.Shopid, ev.Event)
			}

		}
	}
}
