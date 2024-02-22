package service

import (
	"log"

	"github.com/pdcgo/common_conf/common_concept"
	"github.com/pdcgo/tokopedia_lib/app/chat/group"
	"github.com/pdcgo/tokopedia_lib/app/chat/helper"
	"github.com/pdcgo/tokopedia_lib/app/chat/sio_event"
	"github.com/pdcgo/tokopedia_lib/lib/chat"
)

type ChatService struct {
	event       *common_concept.CoreEvent
	socketGroup *group.SocketGroup
	sound       *helper.SoundPlayer
}

func NewChatService(
	event *common_concept.CoreEvent,
	socketGroup *group.SocketGroup,
	sound *helper.SoundPlayer,
) *ChatService {

	chatService := ChatService{
		event:       event,
		socketGroup: socketGroup,
		sound:       sound,
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
		err := sc.SendEvent(readEvent)
		return err
	})
}

type SendChat struct {
	Message        string               `json:"message"`
	MessageId      int64                `json:"message_id"`
	Sticker        *chat.Payload        `json:"sticker,omitempty"`
	ParentReply    *chat.ParentReply    `json:"parent_reply,omitempty"`
	ProductId      int64                `json:"product_id,omitempty"`
	ProductProfile *chat.ProductProfile `json:"product_profile,omitempty"`
	Voucher        *chat.Voucher        `json:"voucher,omitempty"`
	Invoice        *chat.InvoiceLink    `json:"invoice,omitempty"`
}

func (c *SendChat) CreateSendChatEventData(name string) *chat.SendChat {

	data := chat.SendChat{
		From:         name,
		FromUserName: name,
		MessageID:    c.MessageId,
		Message:      c.Message,
		ParentReply:  c.ParentReply,
		Source:       "inbox",
	}

	if c.Sticker != nil {
		data.Payload = c.Sticker
		data.Message = c.Sticker.Intention
		data.AttachmentType = 21

	} else if c.ProductId > 0 {
		data.ProductId = c.ProductId
		data.ProductProfile = c.ProductProfile
		data.AttachmentType = 3

	} else if c.Voucher != nil {
		data.Payload = c.Voucher
		data.AttachmentType = 11

	} else if c.Invoice != nil {
		data.Payload = c.Invoice
		data.InvoiceLink = c.Invoice
		data.AttachmentType = 7
	}

	return &data
}

func (s *ChatService) SendChat(username, name string, sendChat *SendChat) error {
	err := s.socketGroup.WithSocket(username, func(sc *chat.SocketClient) error {

		data := sendChat.CreateSendChatEventData(name)
		event := chat.EmitEventSocket{
			EventCode: &chat.EventCode{
				Code: chat.ChatEvent,
			},
			Data: data,
		}

		log.Printf("[ %s ] send message attach:%d", username, data.AttachmentType)
		err := sc.SendEvent(&event)

		return err
	})
	return err
}

func (s *ChatService) handleEvent() {
	for event := range s.event.GetEvent() {
		switch ev := event.(type) {

		case *sio_event.SendChatEvent:
			if ev.Event.IsOpposite {
				s.sound.Play()
			}
		}
	}
}
