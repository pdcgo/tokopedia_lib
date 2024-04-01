package model

type OrderStatus string

const (
	UnpaidOrder        OrderStatus = "unpaid_order"
	NewOrder           OrderStatus = "new_order"
	ConfirmShipping    OrderStatus = "confirm_shipping"
	InShipping         OrderStatus = "in_shipping"
	WaitingPickup      OrderStatus = "waiting_pickup"
	Shipped            OrderStatus = "shipped"
	WaitingAwbStatus   OrderStatus = "waiting_awb_status"
	AwbChanged         OrderStatus = "awb_changed"
	ReadyToCollect     OrderStatus = "ready_to_collect"
	OrderReturned      OrderStatus = "order_returned"
	Delivered          OrderStatus = "delivered"
	Complaint          OrderStatus = "complaint"
	Done               OrderStatus = "done"
	OrderCanceled      OrderStatus = "order_canceled"
	BuyerInstantCancel OrderStatus = "buyer_instant_cancel"
	SellerReject       OrderStatus = "seller_reject"
	SystemReject       OrderStatus = "system_reject"
)

type OrderStatusesMap map[OrderStatus][]int

func (s OrderStatusesMap) Batch(statuses ...OrderStatus) (res []int) {

	statusUsed := map[int]bool{}
	for _, status := range statuses {
		sts := s[status]
		for _, snum := range sts {
			if !statusUsed[snum] {
				statusUsed[snum] = true
				res = append(res, snum)
			}
		}
	}
	return
}

var OrderStatuses = OrderStatusesMap{
	UnpaidOrder:     {},
	NewOrder:        {220},
	ConfirmShipping: {400, 520},

	InShipping:       {450, 500, 501, 530, 540, 550, 600},
	WaitingPickup:    {450},
	Shipped:          {500},
	WaitingAwbStatus: {501},
	AwbChanged:       {530},
	ReadyToCollect:   {540},
	OrderReturned:    {550},
	Delivered:        {600},

	Complaint: {601},
	Done:      {690, 691, 695, 698, 699, 700, 701},

	OrderCanceled:      {0, 4, 6, 10, 11, 15},
	BuyerInstantCancel: {15},
	SellerReject:       {10},
	SystemReject:       {0, 4, 6, 11},
}
