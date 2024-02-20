package model

import "strings"

type InvoicePaymentSummary struct {
	Label    string `json:"label"`
	Value    string `json:"value"`
	Key      string `json:"key"`
	Typename string `json:"__typename"`
}

type InvoiceShippingData struct {
	ShippingAddress        string `json:"shipping_address"`
	ShippingCourier        string `json:"shipping_courier"`
	ShippingWeight         string `json:"shipping_weight"`
	ReceiverName           string `json:"receiver_name"`
	ReceiverPhone          string `json:"receiver_phone"`
	IsBebasOngkir          bool   `json:"is_bebas_ongkir"`
	BebasOngkirImgURL      string `json:"bebas_ongkir_img_url"`
	IsShippingInsurance    bool   `json:"is_shipping_insurance"`
	ShippingInsuranceLabel string `json:"shipping_insurance_label"`
	Typename               string `json:"__typename"`
}

type InvoiceDropshipInfo struct {
	Name     string `json:"name"`
	Phone    string `json:"phone"`
	Typename string `json:"__typename"`
}

type TokopediaMerchandise struct {
	Name             string `json:"name"`
	Address          string `json:"address"`
	Npwp             string `json:"npwp"`
	WordingPpn       string `json:"wording_ppn"`
	IsStampDuty      bool   `json:"is_stamp_duty"`
	WordingStampDuty string `json:"wording_stamp_duty"`
	Typename         string `json:"__typename"`
}

type InvoiceFeatures struct {
	PartialFulfillment any    `json:"partial_fulfillment"`
	Typename           string `json:"__typename"`
}

type InvoiceDetailNonBundle struct {
	ProductID            int64  `json:"product_id"`
	ProductName          string `json:"product_name"`
	ProductPrice         string `json:"product_price"`
	ProductPriceTotal    string `json:"product_price_total"`
	ProductQty           int    `json:"product_qty"`
	ProductSkuID         string `json:"product_sku_id"`
	ProductWeight        string `json:"product_weight"`
	ProductNotes         string `json:"product_notes"`
	ProductCashbackLabel string `json:"product_cashback_label"`
	ProductSnapshotURL   string `json:"product_snapshot_url"`
	IsPpp                bool   `json:"is_ppp"`
	AddonSummary         any    `json:"addon_summary"`
	Typename             string `json:"__typename"`
}

type InvoiceOrderDataDetail struct {
	Bundles       []any                     `json:"bundles"`
	NonBundles    []*InvoiceDetailNonBundle `json:"non_bundles"`
	BundleIcon    string                    `json:"bundle_icon"`
	TotalProducts int                       `json:"total_products"`
	AddonIcon     string                    `json:"addon_icon"`
	AddonLabel    string                    `json:"addon_label"`
	BmgmIcon      string                    `json:"bmgm_icon"`
	Bmgms         []any                     `json:"bmgms"`
	Typename      string                    `json:"__typename"`
}

type InvoiceOrderData struct {
	InvoiceDate           string                   `json:"invoice_date"`
	InvoiceRefNum         string                   `json:"invoice_ref_num"`
	Status                string                   `json:"status"`
	UpdateTime            string                   `json:"update_time"`
	ShopName              string                   `json:"shop_name"`
	CustomerName          string                   `json:"customer_name"`
	IsTokoCabang          bool                     `json:"is_toko_cabang"`
	IsPreOrder            bool                     `json:"is_pre_order"`
	PreOrderDay           string                   `json:"pre_order_day"`
	AdditionalStatusLabel string                   `json:"additional_status_label"`
	HaveProductBundle     bool                     `json:"have_product_bundle"`
	PaymentSummaryInvoice []*InvoicePaymentSummary `json:"payment_summary_invoice"`
	ShippingData          *InvoiceShippingData     `json:"shipping_data"`
	DropshipInfo          *InvoiceDropshipInfo     `json:"dropship_info"`
	TokopediaMerchandise  *TokopediaMerchandise    `json:"tokopedia_merchandise"`
	AddonInfo             any                      `json:"addon_info"`
	Features              *InvoiceFeatures         `json:"features"`
	Details               *InvoiceOrderDataDetail  `json:"details"`
	Typename              string                   `json:"__typename"`
}

type Invoice struct {
	OrderData   []InvoiceOrderData `json:"order_data"`
	PaymentData struct {
		PaymentMethod []struct {
			Name              string `json:"name"`
			ImgURL            string `json:"img_url"`
			Amount            string `json:"amount"`
			InstallmentDetail string `json:"installment_detail"`
			Typename          string `json:"__typename"`
		} `json:"payment_method"`
		PaymentDetails               []any  `json:"payment_details"`
		OrderCount                   int    `json:"order_count"`
		Orders                       []any  `json:"orders"`
		VoucherDetails               []any  `json:"voucher_details"`
		VoucherDisclaimer            string `json:"voucher_disclaimer"`
		VoucherDisclaimerDisplayType int    `json:"voucher_disclaimer_display_type"`
		Typename                     string `json:"__typename"`
	} `json:"payment_data"`
	Typename string `json:"__typename"`
}

type GetInvoiceV3 struct {
	Invoices []*Invoice `json:"invoices"`
	Typename string     `json:"__typename"`
}

type InvoiceData struct {
	GetInvoiceV3 *GetInvoiceV3 `json:"get_invoice_v3"`
}

type InvoiceRes struct {
	Data *InvoiceData `json:"data"`
}

func (inv *InvoiceRes) IsCod() bool {

	invoices := inv.Data.GetInvoiceV3.Invoices
	if len(invoices) > 0 {
		for _, paymetd := range invoices[0].PaymentData.PaymentMethod {
			payname := strings.ToLower(paymetd.Name)
			isCod := strings.Contains(payname, "cod")
			if isCod {
				return true
			}
		}
	}

	return false
}
