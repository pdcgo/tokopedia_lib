package model

import (
	"database/sql"
	"encoding/json"
)

type OrderItem struct {
	OrderID        int          `gorm:"index;constraint:OnDelete:CASCADE" json:"order_id"`
	ShopID         int          `gorm:"index;constraint:OnDelete:CASCADE"`
	ProductID      int          `gorm:"unique" json:"product_id"`
	Image          string       `json:"image"`
	ProductName    string       `json:"product_name"`
	Url            string       `json:"url"`
	Count          int          `json:"count"`
	Price          int          `json:"price"`
	NotFound       bool         `json:"not_found"`
	ProductCreated sql.NullTime `json:"product_created"`

	Order Order `json:"-"`
}

func (o *OrderItem) MarshalJSON() ([]byte, error) {

	type Alias OrderItem
	return json.Marshal(&struct {
		ProductCreated string `json:"product_created"`
		*Alias
	}{
		ProductCreated: dateStr(o.ProductCreated),
		Alias:          (*Alias)(o),
	})
}

type OrderSheet struct {
	OrderID int    `gorm:"primaryKey;autoIncrement:false;constraint:OnDelete:CASCADE" json:"order_id"`
	Status  string `json:"status"`

	Order Order `json:"order"`
}

type Order struct {
	ID        int    `gorm:"primaryKey"  json:"id"`
	ShopID    int    `gorm:"constraint:OnDelete:CASCADE" json:"shop_id"`
	StatusID  int    `json:"status_id"`
	Status    string `json:"status"`
	InvoiceID string `json:"invoice_id"`

	ItemCount   int `json:"item_count"`
	Fee         int `json:"fee"`
	ShippingFee int `json:"shipping_fee"`
	Total       int `json:"total"`

	BuyerName   string `json:"buyer_name"`
	BuyerID     string `json:"buyer_id"`
	CourierName string `json:"courier_name"`
	Resi        string `json:"resi"`

	DestinationReceiverName string `json:"destination_receiver_name"`
	DestinationPhone        string `json:"destination_phone"`
	DestinationProvince     string `json:"destination_province"`
	DestinationCity         string `json:"destination_city"`
	DestinationDistrict     string `json:"destination_district"`
	DestinationStreet       string `json:"destination_street"`
	DestinationPostalCode   string `json:"destination_postal_code"`

	ArrivedEstimation sql.NullTime `json:"arrived_estimation"`
	ProcessBefore     sql.NullTime `json:"process_before"`
	Created           sql.NullTime `json:"created"`
	PaymentDeadline   sql.NullTime `json:"payment_deadline"`

	Account    *Account     `gorm:"foreignKey:ShopID" json:"account"`
	OrderItems []*OrderItem `json:"order_items"`
	OrderSheet *OrderSheet  `json:"order_sheet"`
}

func dateStr(time sql.NullTime) string {

	if time.Valid {
		return time.Time.String()
	}
	return ""
}

func (o *Order) MarshalJSON() ([]byte, error) {

	type Alias Order
	return json.Marshal(&struct {
		ArrivedEstimation string `json:"arrived_estimation"`
		ProcessBefore     string `json:"process_before"`
		Created           string `json:"created"`
		PaymentDeadline   string `json:"payment_deadline"`
		*Alias
	}{
		ArrivedEstimation: dateStr(o.ArrivedEstimation),
		ProcessBefore:     dateStr(o.ProcessBefore),
		Created:           dateStr(o.Created),
		PaymentDeadline:   dateStr(o.PaymentDeadline),
		Alias:             (*Alias)(o),
	})
}
