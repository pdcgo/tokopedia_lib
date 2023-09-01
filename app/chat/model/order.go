package model

import (
	"database/sql"
)

type OrderItem struct {
	OrderID        int          `gorm:"constraint:OnDelete:CASCADE" json:"order_id"`
	ShopID         int          `gorm:"constraint:OnDelete:CASCADE"`
	ProductID      int          `json:"product_id"`
	Image          string       `json:"image"`
	ProductName    string       `json:"product_name"`
	Url            string       `json:"url"`
	Count          int          `json:"count"`
	Price          int          `json:"price"`
	NotFound       bool         `json:"not_found"`
	ProductCreated sql.NullTime `json:"product_created"`

	Order Order `json:"order"`
}

type OrderSheet struct {
	OrderID int    `gorm:"primaryKey;autoIncrement:false;constraint:OnDelete:CASCADE" json:"order_id"`
	Status  string `json:"status"`

	Order Order `json:"order"`
}

type Order struct {
	ID        int    `gorm:"primaryKey;autoIncrement"  json:"id"`
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

	Account    Account     `gorm:"foreignKey:ShopID" json:"account"`
	OrderItems []OrderItem `json:"order_items"`
	OrderSheet *OrderSheet `json:"order_sheet"`
}
