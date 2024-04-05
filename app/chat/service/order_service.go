package service

import (
	"database/sql"
	"log"
	"strconv"
	"strings"

	"github.com/pdcgo/common_conf/common_concept"
	"github.com/pdcgo/common_conf/pdc_common"
	"github.com/pdcgo/tokopedia_lib/app/chat/group"
	"github.com/pdcgo/tokopedia_lib/app/chat/model"
	"github.com/pdcgo/tokopedia_lib/app/chat/repo"
	"github.com/pdcgo/tokopedia_lib/app/chat/sio_event"
	apimodel "github.com/pdcgo/tokopedia_lib/lib/model"
	"github.com/pdcgo/tokopedia_lib/lib/query"
)

type OrderService struct {
	event       *common_concept.CoreEvent
	accountRepo *repo.AccountRepo
	orderRepo   *repo.OrderRepo
	driverGroup *group.DriverGroup
}

func NewOrderService(
	event *common_concept.CoreEvent,
	accountRepo *repo.AccountRepo,
	orderRepo *repo.OrderRepo,
	driverGroup *group.DriverGroup,
) *OrderService {

	orderservice := OrderService{
		event:       event,
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

func (s *OrderService) syncUpdateOrderItem(shopid, orderid int, item *apimodel.OrderProduct) func(upitem *model.OrderItem) error {
	return func(upitem *model.OrderItem) (err error) {
		upitem.OrderID = orderid
		upitem.ShopID = shopid
		upitem.Image = item.Picture
		upitem.ProductName = item.ProductName
		upitem.Count = item.ProductQty
		upitem.Price, err = parsePrice(item.ProductPrice)
		upitem.Url = item.SnapshotURL
		return
	}
}

func (s *OrderService) syncUpdateOrder(shopid int, order *apimodel.OrderItem) func(uporder *model.Order) error {
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
		uporder.Created = sql.NullTime{
			Time:  created,
			Valid: !created.IsZero(),
		}
		if err != nil {
			return
		}

		processBefore, err := order.GetDiprosesSebelum()
		uporder.ProcessBefore = sql.NullTime{
			Time:  processBefore,
			Valid: !processBefore.IsZero(),
		}
		if err != nil {
			return
		}

		for _, product := range order.OrderProduct {
			syncItem := s.syncUpdateOrderItem(shopid, order.ID, product)
			if err := s.orderRepo.CreateOrUpdateOrderItem(product.ProductID, syncItem); err != nil {
				return err
			}
		}

		return nil
	}
}

func (s *OrderService) syncOrder(ev *sio_event.SyncAccountNotificationEvent) {
	err := s.driverGroup.WithDriverApiByShopid(ev.Shopid, func(username string, dapi *group.DriverApi) error {
		log.Printf("[ %s ] syncronize order", username)

		err := s.updateNotifHash(ev)
		if err != nil {
			return err
		}

		payload := query.NewOrderListQuery()
		return dapi.Api.IterateOrder(payload, func(order *apimodel.OrderItem) error {
			log.Printf("[ %s ] updating order %s", username, order.OrderResi)
			return s.orderRepo.CreateOrUpdateOrder(order.ID, s.syncUpdateOrder(ev.Shopid, order))
		})
	})

	if err != nil {
		pdc_common.ReportError(err)
	}
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
