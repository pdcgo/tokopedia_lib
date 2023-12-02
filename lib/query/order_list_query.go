package query

import "time"

type OrderListInput struct {
	StatusKey         string `json:"status_key"`
	Deadline          int    `json:"deadline"`
	IsShowFilter      any    `json:"is_show_filter"`
	IsPage            any    `json:"is_page"`
	SortBy            int    `json:"sort_by"`
	ShippingID        any    `json:"shipping_id"`
	Search            string `json:"search"`
	StartDate         string `json:"start_date"`
	EndDate           string `json:"end_date"`
	FilterStatus      int    `json:"filter_status"`
	Lang              string `json:"lang"`
	IsResiOnly        any    `json:"is_resi_only"`
	OrderTypeList     []any  `json:"order_type_list"`
	StatusList        []any  `json:"status_list"`
	Page              int    `json:"page"`
	BatchPage         int    `json:"batch_page"`
	ShowPage          int    `json:"show_page"`
	ShopID            int    `json:"shop_id"`
	WarehouseID       []any  `json:"warehouse_id"`
	ShippingList      []any  `json:"shipping_list"`
	IsShippingPrinted int    `json:"is_shipping_printed"`
	FirstDate         int    `json:"first_date"`
	LastDate          int    `json:"last_date"`
	FirstOrderID      int    `json:"first_order_id"`
	LastOrderID       int    `json:"last_order_id"`
	Source            string `json:"source"`
}

type OrderListQuery struct {
	Input *OrderListInput `json:"input"`
}

var now = time.Now()
var OrderDateFormat = "02/01/2006"

func NewOrderListQuery() *OrderListQuery {
	query := OrderListQuery{
		Input: &OrderListInput{
			FilterStatus:  999,
			Lang:          "id",
			SortBy:        3,
			Source:        "som-desktop",
			StatusKey:     "all_order",
			OrderTypeList: []any{},
			StatusList:    []any{},
			WarehouseID:   []any{},
			ShippingList:  []any{},
		},
	}
	query.SetPage(1)
	query.SetYear(now.Year())
	return &query
}

func (q *OrderListQuery) SetPage(page int) {
	q.Input.BatchPage = 1
	q.Input.Page = 1
	q.Input.ShowPage = 1
}

func (q *OrderListQuery) SetYear(year int) {
	startDate := time.Date(year, 1, 1, 0, 0, 0, 0, time.UTC)
	endDate := time.Date(year, 12, 31, 0, 0, 0, 0, time.UTC)

	q.Input.StartDate = startDate.Format(OrderDateFormat)
	q.Input.EndDate = endDate.Format(OrderDateFormat)
}

type OrderIncomeDetailInput struct {
	OrderID int `json:"order_id"`
}

type OrderIncomeDetailQuery struct {
	Input *OrderIncomeDetailInput `json:"input"`
}
