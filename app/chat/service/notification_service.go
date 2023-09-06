package service

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"log"

	"github.com/pdcgo/common_conf/common_concept"
	"github.com/pdcgo/tokopedia_lib/app/chat/group"
	"github.com/pdcgo/tokopedia_lib/app/chat/model"
	"github.com/pdcgo/tokopedia_lib/app/chat/sio_event"
	"github.com/pdcgo/tokopedia_lib/lib/api"
)

type NotificationService struct {
	event          *common_concept.CoreEvent
	driverGroup    *group.DriverGroup
	accountService *AccountService
}

func NewNotificationService(
	event *common_concept.CoreEvent,
	driverGroup *group.DriverGroup,
	accountService *AccountService,
) *NotificationService {

	return &NotificationService{
		event:          event,
		driverGroup:    driverGroup,
		accountService: accountService,
	}
}

type AccountNotificationEvent struct {
	Shopid              int `json:"shopid"`
	ArriveAtDestination int `json:"arriveAtDestination"`
	NewOrder            int `json:"newOrder"`
	ReadyToShip         int `json:"readyToShip"`
	Shipped             int `json:"shipped"`
}

func (ev *AccountNotificationEvent) GetHash() (string, error) {

	b, err := json.Marshal(ev)
	if err != nil {
		return "", err
	}

	hash := md5.Sum(b)
	return hex.EncodeToString(hash[:]), nil
}

func (s *NotificationService) SendAccountNotifications(account *model.Account) error {

	username := account.GetUsername()
	return s.driverGroup.WithDriverApi(username, func(api *api.TokopediaApi) error {

		notif, err := api.NotificationCounter()
		if err != nil {
			return err
		}

		err = s.accountService.UpdateAccountNotifications(account.ID, notif)
		if err != nil {
			return err
		}

		accountNotifEv := AccountNotificationEvent{
			Shopid:              account.ID,
			ArriveAtDestination: notif.Data.Notifications.SellerOrderStatus.ArriveAtDestination,
			NewOrder:            notif.Data.Notifications.SellerOrderStatus.NewOrder,
			ReadyToShip:         notif.Data.Notifications.SellerOrderStatus.ReadyToShip,
			Shipped:             notif.Data.Notifications.SellerOrderStatus.Shipped,
		}
		hash, err := accountNotifEv.GetHash()
		if err != nil {
			return err
		}

		s.event.Emit(&accountNotifEv)
		log.Printf("[ %s ] send notification %s", username, hash)

		feNotifEv := sio_event.FrontendNotificationEvent{
			Shopid: account.ID,
			Event:  notif,
		}
		s.event.Emit(&feNotifEv)

		return nil
	})
}
