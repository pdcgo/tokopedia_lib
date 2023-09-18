package sio_event

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"

	"github.com/pdcgo/tokopedia_lib/lib/api"
)

type NotificationEvent struct {
	Shopid int                         `json:"shopid"`
	Event  *api.NotificationCounterRes `json:"event"`
}

type SyncAccountNotificationEvent struct {
	Shopid              int `json:"shopid"`
	ArriveAtDestination int `json:"arriveAtDestination"`
	NewOrder            int `json:"newOrder"`
	ReadyToShip         int `json:"readyToShip"`
	Shipped             int `json:"shipped"`
}

func NewSyncAccountNotificationEvent(shopid int, notif *api.NotificationCounterRes) *SyncAccountNotificationEvent {
	return &SyncAccountNotificationEvent{
		Shopid:              shopid,
		ArriveAtDestination: notif.Data.Notifications.SellerOrderStatus.ArriveAtDestination,
		NewOrder:            notif.Data.Notifications.SellerOrderStatus.NewOrder,
		ReadyToShip:         notif.Data.Notifications.SellerOrderStatus.ReadyToShip,
		Shipped:             notif.Data.Notifications.SellerOrderStatus.Shipped,
	}
}

func (ev *SyncAccountNotificationEvent) GetHash() (string, error) {

	b, err := json.Marshal(ev)
	if err != nil {
		return "", err
	}

	hash := md5.Sum(b)
	return hex.EncodeToString(hash[:]), nil
}
