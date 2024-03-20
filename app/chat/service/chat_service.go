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
	"github.com/pdcgo/tokopedia_lib/lib/chat"
)

type ChatService struct {
	event        *common_concept.CoreEvent
	initConfig   *config.InitConfig
	areplyConfig *config.AutoReplyConfig
	accountRepo  *repo.AccountRepo
	socketGroup  *group.SocketGroup
	sound        *helper.SoundPlayer
}

func NewChatService(
	event *common_concept.CoreEvent,
	initConfig *config.InitConfig,
	areplyConfig *config.AutoReplyConfig,
	accountRepo *repo.AccountRepo,
	socketGroup *group.SocketGroup,
	sound *helper.SoundPlayer,
) *ChatService {

	chatService := ChatService{
		event:        event,
		initConfig:   initConfig,
		areplyConfig: areplyConfig,
		accountRepo:  accountRepo,
		socketGroup:  socketGroup,
		sound:        sound,
	}

	go chatService.handleEvent()
	return &chatService
}

func (s *ChatService) ReadChat(username string, msgId uint) error {

	readEvent := chat.BaseSocketType{
		Code: chat.ReadUserChatEvent,
		Data: chat.MessageId{
			MsgId: msgId,
		},
	}

	return s.socketGroup.WithSocket(username, func(sc *chat.SocketClient) error {
		return sc.SendEvent(readEvent)
	})
}

func (s *ChatService) SendChat(username string, data *chat.SendChat) error {
	err := s.socketGroup.WithSocket(username, func(sc *chat.SocketClient) error {

		log.Printf("[ %s ] send message attach:%d", username, data.AttachmentType)
		return sc.SendEvent(&chat.EmitEventSocket{
			EventCode: &chat.EventCode{
				Code: chat.ChatEvent,
			},
			Data: data,
		})
	})
	return err
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
				return s.SendChat(account.GetUsername(), sendpayload.CreateEventData(account.ShopName))
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
