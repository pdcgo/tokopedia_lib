package service

import (
	"log"

	socketio "github.com/googollee/go-socket.io"
	"github.com/pdcgo/common_conf/common_concept"
	"github.com/pdcgo/common_conf/pdc_common"
	"github.com/pdcgo/tokopedia_lib"
	"github.com/pdcgo/tokopedia_lib/app/chat/config"
	"github.com/pdcgo/tokopedia_lib/app/chat/group"
	"github.com/pdcgo/tokopedia_lib/app/chat/model"
	"github.com/pdcgo/tokopedia_lib/app/chat/repo"
	"github.com/pdcgo/tokopedia_lib/app/chat/sio_event"
	"github.com/pdcgo/tokopedia_lib/lib/api"
	"github.com/rs/zerolog"
)

type NotificationService struct {
	event          *common_concept.CoreEvent
	initConfig     *config.InitConfig
	sio            *socketio.Server
	accountRepo    *repo.AccountRepo
	driverGroup    *group.DriverGroup
	accountService *AccountService
}

func NewNotificationService(
	event *common_concept.CoreEvent,
	initConfig *config.InitConfig,
	sio *socketio.Server,
	accountRepo *repo.AccountRepo,
	driverGroup *group.DriverGroup,
	accountService *AccountService,
) *NotificationService {

	notificationService := NotificationService{
		event:          event,
		initConfig:     initConfig,
		sio:            sio,
		accountRepo:    accountRepo,
		driverGroup:    driverGroup,
		accountService: accountService,
	}

	go notificationService.handleEvent()

	return &notificationService
}

func (s *NotificationService) SendSyncAccountNotification(account *model.Account) error {

	username := account.GetUsername()
	return s.driverGroup.WithDriverApi(username, func(driver *tokopedia_lib.DriverAccount, api *api.TokopediaApi) error {

		notif, err := api.NotificationCounter()
		if err != nil {
			return err
		}

		syncAccountEvent := sio_event.NewSyncAccountNotificationEvent(account.ID, notif)
		hash, err := syncAccountEvent.GetHash()
		if err != nil {
			return err
		}

		err = s.accountService.SyncAccount(account.ID, hash, notif)
		if err != nil {
			return err
		}

		log.Printf("[ %s ] send notification %s", username, hash)

		s.sio.BroadcastToNamespace("", "notification_event", &syncAccountEvent)
		s.sio.BroadcastToNamespace("", "notification", &sio_event.NotificationEvent{
			Shopid: account.ID,
			Event:  notif,
		})

		return nil
	})
}

func (s *NotificationService) handleEvent() {
	for event := range s.event.GetEvent() {
		switch ev := event.(type) {

		case *sio_event.SocketSyncEvent:

			account, err := s.accountRepo.GetChatAccount(s.initConfig.ActiveGroup, ev.Shopid)
			if err != nil {
				pdc_common.ReportErrorCustom(err, func(event *zerolog.Event) *zerolog.Event {
					return event.Str("event", "sync").Int("shopid", ev.Shopid)
				})
			}

			err = s.SendSyncAccountNotification(account)
			if err != nil {
				pdc_common.ReportErrorCustom(err, func(event *zerolog.Event) *zerolog.Event {
					return event.Str("event", "sync").Int("shopid", ev.Shopid)
				})
			}
		}
	}
}
