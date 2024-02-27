package autochat

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"net/url"
	"strings"

	"github.com/pdcgo/common_conf/pdc_common"
	"github.com/pdcgo/tokopedia_lib/app/autochat/report"
	"github.com/pdcgo/tokopedia_lib/lib/api_public"
	"github.com/pdcgo/tokopedia_lib/lib/chat"
	"github.com/pdcgo/tokopedia_lib/lib/model_public"
)

type AutochatSend struct {
	sender       *AutochatSender
	pubapi       *api_public.TokopediaApiPublic
	reportUpdate report.EditAutosendReportItem
	config       *AutosendConfig
}

func NewAutochatSend(
	sender *AutochatSender,
	pubapi *api_public.TokopediaApiPublic,
	reportUpdate report.EditAutosendReportItem,
	config *AutosendConfig,
) *AutochatSend {

	return &AutochatSend{
		sender:       sender,
		pubapi:       pubapi,
		reportUpdate: reportUpdate,
		config:       config,
	}
}

func (s *AutochatSend) GetShopCore(shopname string) (*model_public.ShopCore, error) {
	shop, err := s.pubapi.ShopCoreInfo(&model_public.ShopCoreInfoVar{
		Domain: shopname,
	})
	if err != nil {
		return nil, err
	}

	if len(shop.Data.Result) == 0 {
		return nil, fmt.Errorf("no shop found")
	}

	return &shop.Data.Result[0].ShopCore, nil
}

func (s *AutochatSend) GetRandomProduct(shopId int) (*model_public.PdpGetLayout, error) {

	payload := model_public.NewShopProductVar(shopId)
	shopProduct, err := s.pubapi.ShopProducts(payload)
	if err != nil {
		return nil, err
	}

	products := shopProduct.Data.GetShopProduct.Data
	productlen := len(products)
	if productlen == 0 {
		return nil, fmt.Errorf("no product found")
	}

	product := products[rand.Intn(productlen-1)]

	layoutPayload, err := model_public.NewPdpGetlayoutQueryVar(product.ProductURL)
	if err != nil {
		return nil, err
	}

	layout, err := s.pubapi.PdpGetlayoutQuery(layoutPayload)
	if err != nil {
		return nil, err
	}

	return &layout.Data.PdpGetLayout, nil
}

func (s *AutochatSend) SendFirstMessage(shopId int, targetMsgID int64) error {

	product, err := s.GetRandomProduct(shopId)
	if err != nil {
		return err
	}

	s.reportUpdate(func(item *report.AutosendReportItem) error {
		item.ProductUrl = product.BasicInfo.URL
		return nil
	})

	err = s.sender.SendProductReply(targetMsgID, product)
	if err != nil {
		return err
	}

	return s.sender.SendReply(targetMsgID, func(chat *chat.SendChat) error {
		return s.reportUpdate(func(item *report.AutosendReportItem) error {
			item.SendChatCount++
			return nil
		})
	})
}

func (s *AutochatSend) WaitNReplies(
	ctx context.Context,
	messages chan *chat.RcvChat,
	targetMsgID int64,
	limit int,
) {

	name := s.sender.GetName()

Parent:
	for message := range messages {
		if message.MsgID == targetMsgID {
			select {

			case <-ctx.Done():
				break Parent

			default:
				s.reportUpdate(func(item *report.AutosendReportItem) error {
					item.SellerReplyCount++
					return nil
				})

				err := s.sender.SendReply(targetMsgID, func(chat *chat.SendChat) error {

					if limit == 1 && s.config.CustomLastMessage != "" {
						chat.Message = s.config.CustomLastMessage
					}

					return s.reportUpdate(func(item *report.AutosendReportItem) error {
						item.SendChatCount++
						return nil
					})
				})
				if err == nil {
					limit--
					log.Printf("[ %s ] %d | remaining %d replies", name, targetMsgID, limit)
				} else {
					pdc_common.ReportError(err)
				}
			}

			if limit <= 0 {
				break Parent
			}
		}
	}
}

func (s *AutochatSend) Run(
	ctx context.Context,
	limit int,
	shopname string,
) error {

	name := s.sender.GetName()
	rctx, cancel := context.WithCancel(ctx)
	defer cancel()

	uri, err := url.Parse(shopname)
	if err != nil {
		return err
	}

	fixshopname := strings.ReplaceAll(uri.Path, "/", "")
	shopCore, err := s.GetShopCore(fixshopname)
	if err != nil {
		return err
	}

	chatExisting, err := s.sender.socket.Api.GetChatExisting(shopCore.ShopID, 0)
	if err != nil {
		return err
	}

	messages := s.sender.GetMessages(rctx)
	targetMsgID := chatExisting.Data.ChatExistingChat.MessageID

	// send first message
	err = s.SendFirstMessage(shopCore.ShopID, targetMsgID)
	if err != nil {
		return err
	}

	// wait replies done
	log.Printf("[ %s ] %d | remaining %d replies", name, targetMsgID, limit-1)
	s.WaitNReplies(rctx, messages, targetMsgID, limit-1)

	return nil
}
