package model

import "time"

type EmptyStateMetadata struct {
	Title    string `json:"title"`
	Subtitle string `json:"subtitle"`
	ImageURL string `json:"image_url"`
	Typename string `json:"__typename"`
}

type AvailableCourierAgency struct {
	Value    int    `json:"value"`
	Text     string `json:"text"`
	Typename string `json:"__typename"`
}

type AvailableCourierProduct struct {
	Value    int    `json:"value"`
	Text     string `json:"text"`
	Typename string `json:"__typename"`
}

type AvailableCourier struct {
	Agency   *AvailableCourierAgency    `json:"agency"`
	Product  []*AvailableCourierProduct `json:"product"`
	Typename string                     `json:"__typename"`
}

type Paging struct {
	ShowBackButton     bool   `json:"show_back_button"`
	ShowNextButton     bool   `json:"show_next_button"`
	PagesShowValueList []int  `json:"pages_show_value_list"`
	PagesRealValueList []int  `json:"pages_real_value_list"`
	CurrentBatchPage   int    `json:"current_batch_page"`
	CurrentPage        int    `json:"current_page"`
	NextChangerValue   int    `json:"next_changer_value"`
	PrevChangerValue   int    `json:"prev_changer_value"`
	Typename           string `json:"__typename"`
}

type ProductSingle struct {
	ProductID    string `json:"productId"`
	SnapshotURL  string `json:"snapshotUrl"`
	ProductName  string `json:"productName"`
	OrderNote    string `json:"orderNote"`
	ProductPrice string `json:"productPrice"`
	ProductQty   int    `json:"productQty"`
	SubTotal     string `json:"subTotal"`
	TotalWeight  string `json:"totalWeight"`
	AddonSummary any    `json:"addonSummary"`
	Returnable   int    `json:"returnable"`
	Sku          string `json:"sku"`
	Picture      string `json:"picture"`
	Typename     string `json:"__typename"`
}

type OrderDetails struct {
	TotalProducts  int              `json:"totalProducts"`
	AddonLabel     string           `json:"addonLabel"`
	IconAddon      string           `json:"iconAddon"`
	IconBundle     string           `json:"iconBundle"`
	IconBMGM       string           `json:"iconBMGM"`
	ProductsBundle []any            `json:"productsBundle"`
	ProductsSingle []*ProductSingle `json:"productsSingle"`
	ProductsBMGM   []any            `json:"productsBMGM"`
	Typename       string           `json:"__typename"`
}

type Ticker struct {
	Text           string `json:"text"`
	ActionText     string `json:"action_text"`
	ActionKey      string `json:"action_key"`
	Type           string `json:"type"`
	CtaActionType  string `json:"cta_action_type"`
	CtaActionValue string `json:"cta_action_value"`
	CtaText        string `json:"cta_text"`
	Typename       string `json:"__typename"`
}

type ExpiredFinishNotifLabel struct {
	IsExpiredFinishLabel bool   `json:"is_expired_finish_label"`
	UnixExpiredLabelTime int    `json:"unix_expired_label_time"`
	UnixFinishOrderTime  int    `json:"unix_finish_order_time"`
	FinishOrderTime      string `json:"finish_order_time"`
	ExpiredLabelTime     string `json:"expired_label_time"`
	Typename             string `json:"__typename"`
}

type OnlineBooking struct {
	IsHideInputAwb   bool   `json:"is_hide_input_awb"`
	IsRemoveInputAwb bool   `json:"is_remove_input_awb"`
	IsShowInfo       bool   `json:"is_show_info"`
	InfoText         string `json:"info_text"`
	Typename         string `json:"__typename"`
}

type OrderProduct struct {
	ProductID    string `json:"productId"`
	SnapshotURL  string `json:"snapshotUrl"`
	ProductName  string `json:"productName"`
	OrderNote    string `json:"orderNote"`
	ProductPrice string `json:"productPrice"`
	ProductQty   int    `json:"productQty"`
	SubTotal     string `json:"subTotal"`
	TotalWeight  string `json:"totalWeight"`
	AddonSummary any    `json:"addonSummary"`
	Returnable   int    `json:"returnable"`
	Sku          string `json:"sku"`
	Picture      string `json:"picture"`
	Typename     string `json:"__typename"`
}

type LogisticInfoAll struct {
	ID            string `json:"id"`
	Priority      int    `json:"priority"`
	Description   string `json:"description"`
	InfoTextShort string `json:"info_text_short"`
	InfoTextLong  string `json:"info_text_long"`
	Typename      string `json:"__typename"`
}

type LogisticInfo struct {
	All      []*LogisticInfoAll `json:"all"`
	Typename string             `json:"__typename"`
}

type OrderTag struct {
	IsAffiliate bool   `json:"isAffiliate"`
	Typename    string `json:"__typename"`
}

type OrderItem struct {
	OrderDetails                *OrderDetails            `json:"orderDetails"`
	PofData                     any                      `json:"pofData"`
	IsMitra                     bool                     `json:"isMitra"`
	Ticker                      *Ticker                  `json:"ticker"`
	AddonInfo                   any                      `json:"addonInfo"`
	SellerNotesText             string                   `json:"seller_notes_text"`
	IsFlaggedOrder              bool                     `json:"is_flagged_order"`
	FulfillBy                   int                      `json:"fulfill_by"`
	KeroCode                    int                      `json:"kero_code"`
	CourierType                 int                      `json:"courier_type"`
	CourierProductID            int                      `json:"courier_product_id"`
	ID                          int                      `json:"id,string"`
	IsAdditionalCost            bool                     `json:"is_additional_cost"`
	HasBookingInfo              int                      `json:"has_booking_info"`
	BookingInfoURL              string                   `json:"booking_info_url"`
	CourierProductName          string                   `json:"courier_product_name"`
	CourierInfo                 string                   `json:"courier_info"`
	CancelRequest               int                      `json:"cancel_request"`
	CancelRequestOriginNote     string                   `json:"cancel_request_origin_note"`
	CancelRequestNote           string                   `json:"cancel_request_note"`
	CancelRequestTime           string                   `json:"cancel_request_time"`
	DropshipName                string                   `json:"dropship_name"`
	DropshipPhone               string                   `json:"dropship_phone"`
	InsuranceType               int                      `json:"insurance_type"`
	Status                      string                   `json:"status"`
	OrderStatusID               int                      `json:"order_status_id"`
	OrderResi                   string                   `json:"order_resi"`
	OrderResiURL                string                   `json:"order_resi_url"`
	OriginAddress               string                   `json:"origin_address"`
	OriginDistrict              int                      `json:"origin_district"`
	OriginGeo                   string                   `json:"origin_geo"`
	OriginPostalCode            string                   `json:"origin_postal_code"`
	IsPurchaseProtection        bool                     `json:"is_purchase_protection"`
	OrderTotalPrice             string                   `json:"order_total_price"`
	Labels                      []any                    `json:"labels"`
	PreorderProcessTimeDaysLeft int                      `json:"preorder_process_time_days_left"`
	BuyerName                   string                   `json:"buyer_name"`
	OrderDate                   string                   `json:"order_date"`
	BuyerID                     string                   `json:"buyer_id"`
	CourierID                   int                      `json:"courier_id"`
	Cashback                    string                   `json:"cashback"`
	CourierName                 string                   `json:"courier_name"`
	IsChecked                   bool                     `json:"isChecked"`
	IsTopads                    bool                     `json:"is_topads"`
	IsBroadcastChat             bool                     `json:"is_broadcast_chat"`
	IsShippingPrinted           bool                     `json:"is_shipping_printed"`
	IsReplacementTaken          int                      `json:"is_replacement_taken"`
	TradeInFee                  string                   `json:"trade_in_fee"`
	DeadlineColor               string                   `json:"deadline_color"`
	DeadlineText                string                   `json:"deadline_text"`
	DeadlineTimeLeft            int                      `json:"deadline_time_left"`
	DeadlineStyle               int                      `json:"deadline_style"`
	DestinationStreet           string                   `json:"destination_street"`
	DestinationDistrict         string                   `json:"destination_district"`
	DestinationCity             string                   `json:"destination_city"`
	DestinationProvince         string                   `json:"destination_province"`
	DestinationPostalCode       string                   `json:"destination_postal_code"`
	DestinationPhone            string                   `json:"destination_phone"`
	DestinationReceiverName     string                   `json:"destination_receiver_name"`
	FreeReturnOrder             bool                     `json:"free_return_order"`
	CourierRef                  string                   `json:"courier_ref"`
	IsFreeShipping              bool                     `json:"is_free_shipping"`
	IsPenaltyReject             bool                     `json:"is_penalty_reject"`
	PenaltyRejectWording        string                   `json:"penalty_reject_wording"`
	IsShowPrintLabel            bool                     `json:"is_show_print_label"`
	IsTokocabang                bool                     `json:"is_tokocabang"`
	ExpiredFinishNotifLabel     *ExpiredFinishNotifLabel `json:"expired_finish_notif_label"`
	OnlineBooking               *OnlineBooking           `json:"online_booking"`
	OrderProduct                []*OrderProduct          `json:"order_product"`
	LogisticInfo                *LogisticInfo            `json:"logistic_info"`
	ButtonList                  []any                    `json:"buttonList"`
	WarehouseName               string                   `json:"warehouse_name"`
	CancelRequestStatus         int                      `json:"cancelRequestStatus"`
	IsShowSellerNotes           bool                     `json:"isShowSellerNotes"`
	IsShowChatButton            bool                     `json:"isShowChatButton"`
	PlusData                    any                      `json:"plus_data"`
	ShipmentLogo                string                   `json:"shipment_logo"`
	HasResoStatus               bool                     `json:"has_reso_status"`
	TxID                        string                   `json:"tx_id"`
	GroupType                   int                      `json:"group_type"`
	Tag                         *OrderTag                `json:"tag"`
	Typename                    string                   `json:"__typename"`
}

func (s *OrderItem) GetTanggalPemesanan() (dtime time.Time, err error) {
	if s.OrderDate != "" {
		dtime, err = time.Parse("02 Jan 2006, 15:04 WIB", s.OrderDate)
		if err != nil {
			return
		}
	}
	return
}

func (s *OrderItem) GetDiprosesSebelum() (dtime time.Time, err error) {
	if s.DeadlineText != "" {

		ordtime, oerr := s.GetTanggalPemesanan()
		if oerr != nil {
			return
		}

		dtime, err = time.Parse("2 Jan; 15:04", s.DeadlineText)
		if err != nil {
			return time.Time{}, err
		}

		// handle tahun baru
		_, dmonth, _ := dtime.Date()
		year, ordmonth, _ := ordtime.Date()
		if ordmonth > dmonth {
			year += 1
		}

		dtime = dtime.AddDate(year, 0, 0)
		return
	}

	return
}

type OrderList struct {
	FirstOrderID         int                 `json:"firstOrderId"`
	LastOrderID          int                 `json:"lastOrderId"`
	FirstDate            int                 `json:"firstDate"`
	LastDate             int                 `json:"lastDate"`
	TotalDataPerBatch    int                 `json:"total_data_per_batch"`
	EmptyStateMetadata   *EmptyStateMetadata `json:"emptyStateMetadata"`
	UserID               string              `json:"user_id"`
	ShopID               string              `json:"shop_id"`
	UserToken            string              `json:"user_token"`
	IsWarehouseAdmin     bool                `json:"is_warehouse_admin"`
	IsShowBulkAction     bool                `json:"isShowBulkAction"`
	IsShowDownloadReport bool                `json:"isShowDownloadReport"`
	AvailableCourier     []*AvailableCourier `json:"available_courier"`
	Paging               Paging              `json:"paging"`
	List                 []*OrderItem        `json:"list"`
	Typename             string              `json:"__typename"`
}

type OrderListResp struct {
	Data struct {
		OrderList OrderList `json:"orderList"`
	} `json:"data"`
}
