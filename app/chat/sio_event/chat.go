package sio_event

import (
	"github.com/pdcgo/tokopedia_lib/lib/chat"
)

type ReadChatEvent struct {
	Shopid int                `json:"shopid,string"`
	Event  *chat.ReaduserChat `json:"event"`
}

type SendChatEvent struct {
	Shopid int           `json:"shopid,string"`
	Event  *chat.RcvChat `json:"event"`
}

type TypingStartChatEvent struct {
	Shopid int                  `json:"shopid,string"`
	Event  *chat.RcvStartTyping `json:"event"`
}

type TypingEndChatEvent struct {
	Shopid int                `json:"shopid,string"`
	Event  *chat.RcvEndTyping `json:"event"`
}
