package service

import (
	"database/sql"
	"log"

	"github.com/pdcgo/common_conf/pdc_common"
	"github.com/pdcgo/tokopedia_lib/app/chat/group"
	"github.com/pdcgo/tokopedia_lib/app/chat/model"
	"github.com/pdcgo/tokopedia_lib/app/chat/sio_event"
	apimodel "github.com/pdcgo/tokopedia_lib/lib/model"
	"github.com/pdcgo/tokopedia_lib/lib/query"
)

func (s *OrderService) syncUpdateUnpaidOrderItem(shopid, orderid int, item *apimodel.OrderPaymentItemProduct) func(upitem *model.OrderItem) error {
	return func(upitem *model.OrderItem) (err error) {
		upitem.OrderID = orderid
		upitem.ShopID = shopid
		upitem.Image = item.Picture
		upitem.ProductName = item.ProductName
		upitem.Count = item.ProductQty
		upitem.Price, err = parsePrice(item.ProductPrice)
		return
	}
}

func (s *OrderService) syncUpdateUnpaidOrder(shopid int, order *apimodel.OrderPaymentItem) func(uporder *model.Order) error {
	return func(uporder *model.Order) error {

		uporder.ShopID = shopid
		uporder.BuyerName = order.BuyerName

		processBefore, err := order.GetDiprosesSebelum()
		uporder.ProcessBefore = sql.NullTime{
			Time:  processBefore,
			Valid: !processBefore.IsZero(),
		}
		if err != nil {
			return err
		}

		var itemCount, total int
		for _, product := range order.Products {

			price, _ := parsePrice(product.ProductPrice)
			itemCount += product.ProductQty
			total += product.ProductQty * price

			syncItem := s.syncUpdateUnpaidOrderItem(shopid, order.OrderId, product)
			if err := s.orderRepo.CreateOrUpdateOrderItem(product.ProductId, syncItem); err != nil {
				return err
			}
		}

		uporder.ItemCount = itemCount
		uporder.Total = total
		return nil
	}
}

func (s *OrderService) syncUnpaidOrder(ev *sio_event.SocketConnectEvent) {
	err := s.driverGroup.WithDriverApi(ev.Shopid, func(dapi *group.DriverApi) error {
		username := dapi.GetUsername()
		log.Printf("[ %s ] syncronize unpaid order", username)

		payload := query.NewOrderPendingListQuery()
		res, err := dapi.Api.OrderPendingList(payload)
		if err != nil {
			return err
		}

		if res.Data.OrderListWaitingPayment != nil {
			for _, order := range res.Data.OrderListWaitingPayment.List {
				log.Printf("[ %s ] updating unpaid order %d", username, order.OrderId)
				err := s.orderRepo.CreateOrUpdateOrder(order.OrderId, s.syncUpdateUnpaidOrder(ev.Shopid, order))
				if err != nil {
					return err
				}
			}
		}
		return nil
	})
	if err != nil {
		pdc_common.ReportError(err)
	}
}

func (s *OrderService) GetUnpaid(shopid int, payload *query.OrderPendingListQuery) (res *apimodel.OrderListWaitingPaymentRes, err error) {
	err = s.driverGroup.WithDriverApi(shopid, func(dapi *group.DriverApi) error {
		res, err = dapi.Api.OrderPendingList(payload)
		return err
	})
	return
}
