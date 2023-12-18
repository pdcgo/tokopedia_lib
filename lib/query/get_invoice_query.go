package query

type GetInvoiceInput struct {
	IsSellerView bool     `json:"is_seller_view"`
	PaymentIds   []any    `json:"payment_ids"`
	Invoices     []string `json:"invoices"`
	Signature    []any    `json:"signature"`
	Time         []any    `json:"time"`
	Token        []any    `json:"token"`
	Source       string   `json:"source"`
}

type GetInvoiceVariable struct {
	Input *GetInvoiceInput `json:"input"`
}

func NewGetInvoiceVariable(invoice string) *GetInvoiceVariable {

	input := GetInvoiceInput{
		IsSellerView: true,
		PaymentIds:   []any{},
		Invoices:     []string{invoice},
		Signature:    []any{},
		Time:         []any{},
		Token:        []any{},
		Source:       "frontend",
	}

	return &GetInvoiceVariable{
		Input: &input,
	}
}
