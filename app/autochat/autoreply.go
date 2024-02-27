package autochat

import (
	"context"
	"log"

	"github.com/pdcgo/common_conf/pdc_common"
	"github.com/pdcgo/tokopedia_lib/app/autochat/report"
	"github.com/pdcgo/tokopedia_lib/lib/chat"
)

type AutochatReply struct {
	sender *AutochatSender
}

func NewAutochatReply(sender *AutochatSender) *AutochatReply {
	return &AutochatReply{
		sender: sender,
	}
}

func (s *AutochatReply) Run(ctx context.Context, autoreport *report.AutoreplyReport) {

	name := s.sender.GetName()
	reportcache := map[string]report.EditAutoreplyReportItem{}
	messages := s.sender.GetMessages(ctx, func(msg *chat.RcvChat) bool {
		return msg.FromRole != "User"
	})

	log.Printf("[ %s ] wait message to reply", name)

Parent:
	for message := range messages {
		select {

		case <-ctx.Done():
			break Parent

		default:

			itemUpdate := reportcache[message.From]
			if itemUpdate == nil {
				_, itemUpdate = autoreport.CreateItem(name, message.FromUserName)
				reportcache[message.From] = itemUpdate
			}

			itemUpdate(func(item *report.AutoreplyReportItem) error {
				item.ReceiveChatCount++
				return nil
			})

			err := s.sender.SendReply(message.MsgID, func(chat *chat.SendChat) error {
				return itemUpdate(func(item *report.AutoreplyReportItem) error {
					item.SendChatCount++
					return nil
				})
			})
			if err != nil {
				pdc_common.ReportError(err)
			}
		}
	}
}
