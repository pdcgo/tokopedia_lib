package service

import (
	"database/sql"
	"strconv"
	"strings"

	"github.com/pdcgo/common_conf/common_concept"
	"github.com/pdcgo/common_conf/pdc_common"
	"github.com/pdcgo/tokopedia_lib"
	"github.com/pdcgo/tokopedia_lib/app/chat/config"
	"github.com/pdcgo/tokopedia_lib/app/chat/group"
	"github.com/pdcgo/tokopedia_lib/app/chat/model"
	"github.com/pdcgo/tokopedia_lib/app/chat/repo"
	"github.com/pdcgo/tokopedia_lib/app/chat/sio_event"
	"github.com/pdcgo/tokopedia_lib/lib/api"
	apimodel "github.com/pdcgo/tokopedia_lib/lib/model"
	"github.com/pdcgo/tokopedia_lib/lib/query"
)

type OrderService struct {
	event       *common_concept.CoreEvent
	initConfig  *config.InitConfig
	accountRepo *repo.AccountRepo
	orderRepo   *repo.OrderRepo
	driverGroup *group.DriverGroup
}

func NewOrderService(
	event *common_concept.CoreEvent,
	initConfig *config.InitConfig,
	accountRepo *repo.AccountRepo,
	orderRepo *repo.OrderRepo,
	driverGroup *group.DriverGroup,
) *OrderService {

	orderservice := OrderService{
		event:       event,
		initConfig:  initConfig,
		accountRepo: accountRepo,
		orderRepo:   orderRepo,
		driverGroup: driverGroup,
	}

	go orderservice.handleEvent()
	return &orderservice
}

func (s *OrderService) updateNotifHash(ev *sio_event.SyncAccountNotificationEvent) error {
	return s.accountRepo.UpdateAccount(ev.Shopid, func(account *model.Account) error {
		hash, err := ev.GetHash()
		if err != nil {
			return err
		}

		account.NotifHash = hash
		return nil
	})
}

func parsePrice(price string) (int, error) {
	price = strings.ReplaceAll(price, "Rp ", "")
	price = strings.ReplaceAll(price, ".", "")
	return strconv.Atoi(price)
}

func syncUpdateOrder(shopid int, order *apimodel.OrderItem) func(uporder *model.Order) error {
	return func(uporder *model.Order) (err error) {
		uporder.ShopID = shopid
		uporder.StatusID = order.OrderStatusID
		uporder.Status = order.Status
		uporder.Resi = order.CourierRef
		uporder.BuyerName = order.BuyerName
		uporder.InvoiceID = order.OrderResi
		uporder.BuyerID = order.BuyerID
		uporder.DestinationCity = order.DestinationCity
		uporder.DestinationDistrict = order.DestinationDistrict
		uporder.DestinationPhone = order.DestinationPhone
		uporder.DestinationPostalCode = order.DestinationPostalCode
		uporder.DestinationProvince = order.DestinationProvince
		uporder.DestinationReceiverName = order.DestinationReceiverName
		uporder.DestinationStreet = order.DestinationStreet
		uporder.CourierName = order.CourierName
		uporder.Fee, _ = parsePrice(order.TradeInFee)

		created, err := order.GetTanggalPemesanan()
		if !created.IsZero() {
			uporder.Created = sql.NullTime{
				Time:  created,
				Valid: true,
			}
		}
		if err != nil {
			return
		}

		processBefore, err := order.GetDiprosesSebelum()
		if !processBefore.IsZero() {
			uporder.ProcessBefore = sql.NullTime{
				Time:  processBefore,
				Valid: true,
			}
		}
		return err
	}
}

func (s *OrderService) syncOrder(ev *sio_event.SyncAccountNotificationEvent) {
	err := s.accountRepo.WithAccount(s.initConfig.ActiveGroup, ev.Shopid, func(account *model.Account) error {

		err := s.updateNotifHash(ev)
		if err != nil {
			return err
		}

		return s.driverGroup.WithDriverApi(account.GetUsername(), func(driver *tokopedia_lib.DriverAccount, api *api.TokopediaApi) error {
			payload := query.NewOrderListQuery()
			return api.IterateOrder(payload, func(order *apimodel.OrderItem) error {
				return s.orderRepo.CreateOrUpdateOrder(order.ID, syncUpdateOrder(ev.Shopid, order))
			})
		})
	})

	if err != nil {
		pdc_common.ReportError(err)
	}
}

func (s *OrderService) syncUnpaidOrder(ev *sio_event.SocketConnectEvent) {
	// err := s.accountRepo.WithAccount(s.initConfig.ActiveGroup, ev.Shopid, func(account *model.Account) error {
	// 	return s.driverGroup.WithDriverApi(account.GetUsername(), func(driver *tokopedia_lib.DriverAccount, api *api.TokopediaApi) error {
	// 		payload := query.NewOrderListQuery()
	// 		return api.IterateOrder(payload, func(order *apimodel.OrderItem) error {
	// 			return s.orderRepo.CreateOrUpdateOrder(order.ID, syncUpdateOrder(ev.Shopid, order))
	// 		})
	// 	})
	// })

	// if err != nil {
	// 	pdc_common.ReportError(err)
	// }
}

func (s *OrderService) handleEvent() {
	for event := range s.event.GetEvent() {
		switch ev := event.(type) {

		case *sio_event.SyncAccountNotificationEvent:
			s.syncOrder(ev)

		case *sio_event.SocketConnectEvent:
			s.syncUnpaidOrder(ev)
		}
	}
}
