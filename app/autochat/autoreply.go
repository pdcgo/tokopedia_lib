package autochat

import (
	"context"
	"log"

	"github.com/pdcgo/common_conf/pdc_common"
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

func (s *AutochatReply) Run(ctx context.Context) {

	name := s.sender.GetName()
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
			err := s.sender.SendReply(message.MsgID)
			if err != nil {
				pdc_common.ReportError(err)
			}
		}
	}
}
