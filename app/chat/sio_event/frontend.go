package sio_event

import "github.com/pdcgo/tokopedia_lib/lib/api"

type FrontendNotificationEvent struct {
	Shopid int                         `json:"shopid"`
	Event  *api.NotificationCounterRes `json:"event"`
}
