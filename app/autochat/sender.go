package autochat

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/pdcgo/common_conf/pdc_common"
	"github.com/pdcgo/tokopedia_lib/lib/api"
	"github.com/pdcgo/tokopedia_lib/lib/chat"
	"github.com/pdcgo/tokopedia_lib/lib/model_public"
)

type AutochatSender struct {
	socket  *chat.SocketClient
	message *AutochatMessage
	config  *AutochatConfig
}

func NewAutochatSender(api *api.TokopediaApi, message *AutochatMessage, config *AutochatConfig) *AutochatSender {

	socket := chat.NewSocketClient(api)
	return &AutochatSender{
		socket:  socket,
		message: message,
		config:  config,
	}
}

func (s *AutochatSender) GetName() string {
	return s.socket.Api.AuthenticatedData.UserShopInfo.Info.ShopName
}

func (s *AutochatSender) ReadMessage(msgId int64) error {
	return s.socket.SendEvent(chat.BaseSocketType{
		Code: chat.ReadUserChatEvent,
		Data: chat.MessageId{
			MsgId: uint(msgId),
		},
	})
}

func (s *AutochatSender) TypingMessage(msgId int64) error {

	err := s.socket.SendEvent(chat.BaseSocketType{
		Code: chat.StartTypingEvent,
		Data: chat.MessageId{
			MsgId: uint(msgId),
		},
	})
	if err != nil {
		return err
	}

	name := s.GetName()
	interval := s.config.SendInterval.Get()
	log.Printf("[ %s ] %d | wait typing interval %ds", name, msgId, interval)
	time.Sleep(time.Duration(interval) * time.Second)

	return s.socket.SendEvent(chat.BaseSocketType{
		Code: chat.EndTypingEvent,
		Data: chat.MessageId{
			MsgId: uint(msgId),
		},
	})
}

type SendReplyHandler func(chat *chat.SendChat) error

func (s *AutochatSender) SendReply(msgId int64, handlers ...SendReplyHandler) error {

	err := s.ReadMessage(msgId)
	if err != nil {
		return err
	}

	err = s.TypingMessage(msgId)
	if err != nil {
		return err
	}

	name := s.GetName()
	message := s.message.GetMessage()

	log.Printf("[ %s ] %d | send message '%s'", name, msgId, message)
	data := chat.SendChat{
		From:         name,
		FromUserName: name,
		MessageID:    msgId,
		Message:      message,
		Source:       "inbox",
		StartTime:    time.Now(),
	}

	for _, handler := range handlers {
		err := handler(&data)
		if err != nil {
			return err
		}
	}

	err = s.socket.SendEvent(chat.EmitEventSocket{
		EventCode: &chat.EventCode{
			Code: chat.ChatEvent,
		},
		Data: data,
	})
	if err != nil {
		return err
	}

	interval := s.config.SendInterval.Get()
	log.Printf("[ %s ] %d | wait send interval %ds", name, msgId, interval)
	time.Sleep(time.Duration(interval) * time.Second)
	return nil
}

func (s *AutochatSender) SendProductReply(msgId int64, product *model_public.PdpGetLayout) error {

	name := s.GetName()
	images, err := product.GetImages()
	if err != nil {
		return err
	}

	stock, err := product.GetStock()
	if err != nil {
		return err
	}

	laycontent, err := model_public.GetComponent[model_public.ProductContentComponent](product)
	if err != nil {
		return err
	}
	content := laycontent.Data[0]
	androidUrl := fmt.Sprintf("tokopedia://product/%d?src=chat", product.BasicInfo.ID)

	log.Printf("[ %s ] %d | send product '%s'", name, msgId, content.Name)
	return s.socket.SendEvent(chat.EmitEventSocket{
		EventCode: &chat.EventCode{
			Code: chat.ChatEvent,
		},
		Data: chat.SendChat{
			AttachmentType: 3,
			From:           name,
			FromUserName:   name,
			MessageID:      msgId,
			ProductId:      product.BasicInfo.ID,
			ProductProfile: &chat.ProductProfile{
				ImageUrl:             images[0],
				Name:                 content.Name,
				Price:                content.Price.PriceFmt,
				PriceInt:             int64(content.Price.Value),
				Url:                  product.BasicInfo.URL + "?src=chat",
				AndroidUrl:           androidUrl,
				IosUrl:               androidUrl,
				PlayStoreProductData: map[string]string{},
				ShopId:               product.BasicInfo.ShopID,
				Status:               1,
				MinOrder:             product.BasicInfo.MinOrder,
				Categoryid:           product.BasicInfo.Category.ID,
				RemainingStock:       stock,
				ListImageUrl:         images,
				Type:                 "product",
				DefaultChild:         product.BasicInfo.ID,
				Id:                   product.BasicInfo.ID,
			},
			Source:    "pdp",
			StartTime: time.Now(),
		},
	})
}

func (s *AutochatSender) GetMessages(ctx context.Context, filters ...func(msg *chat.RcvChat) bool) chan *chat.RcvChat {

	messages := make(chan *chat.RcvChat, s.config.Concurrent)
	name := s.GetName()

	go func() {
		<-ctx.Done()
		close(messages)
	}()

	go s.socket.Connect(ctx,
		func(socket *chat.SocketClient, event *chat.RcvEventSocket) error {
			switch data := event.Data.(type) {
			case *chat.RcvChat:

				inPattern := s.message.InPattern(data.Message.OriginalReply)
				if data.IsOpposite && inPattern {

					select {
					case <-ctx.Done():
						return nil

					default:
						for _, filter := range filters {
							if filter(data) {
								return nil
							}
						}

						log.Printf("[ %s ] %d | get message in pattern '%s'", name, data.MsgID, data.Message.OriginalReply)
						messages <- data
					}
				}
			}

			return nil

		}, func(socket *chat.SocketClient, err error) bool {
			pdc_common.ReportError(err)
			return false
		})

	return messages
}
