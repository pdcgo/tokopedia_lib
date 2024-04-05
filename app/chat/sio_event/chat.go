package sio_event

import (
	"github.com/pdcgo/tokopedia_lib/lib/chat"
)

type ReadChatEvent struct {
	Shopid int                `json:"shopid,string"`
	Event  *chat.ReaduserChat `json:"event"`
}

func NewReadChatEvent(shopid int, event *chat.ReaduserChat) *ReadChatEvent {
	return &ReadChatEvent{
		Shopid: shopid,
		Event:  event,
	}
}

type SendChatEvent struct {
	Shopid int           `json:"shopid,string"`
	Event  *chat.RcvChat `json:"event"`
}

func NewSendChatEvent(shopid int, event *chat.RcvChat) *SendChatEvent {
	return &SendChatEvent{
		Shopid: shopid,
		Event:  event,
	}
}

type TypingStartChatEvent struct {
	Shopid int                  `json:"shopid,string"`
	Event  *chat.RcvStartTyping `json:"event"`
}

func NewTypingStartChatEvent(shopid int, event *chat.RcvStartTyping) *TypingStartChatEvent {
	return &TypingStartChatEvent{
		Shopid: shopid,
		Event:  event,
	}
}

type TypingEndChatEvent struct {
	Shopid int                `json:"shopid,string"`
	Event  *chat.RcvEndTyping `json:"event"`
}

func NewTypingEndChatEvent(shopid int, event *chat.RcvEndTyping) *TypingEndChatEvent {
	return &TypingEndChatEvent{
		Shopid: shopid,
		Event:  event,
	}
}
