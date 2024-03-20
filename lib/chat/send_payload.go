package chat

import "time"

type SendChatPayload struct {
	Message        string          `json:"message"`
	MessageId      int64           `json:"message_id"`
	Sticker        *Payload        `json:"sticker,omitempty"`
	ParentReply    *ParentReply    `json:"parent_reply,omitempty"`
	From           string          `json:"from,omitempty"`
	FromUserName   string          `json:"from_user_name,omitempty"`
	ProductId      int64           `json:"product_id,omitempty"`
	ProductProfile *ProductProfile `json:"product_profile,omitempty"`
	Voucher        *Voucher        `json:"voucher,omitempty"`
	Invoice        *InvoiceLink    `json:"invoice,omitempty"`
}

func (c *SendChatPayload) CreateEventData(name string) *SendChat {

	data := SendChat{
		From:         name,
		FromUserName: name,
		StartTime:    time.Now(),
		MessageID:    c.MessageId,
		Message:      c.Message,
		ParentReply:  c.ParentReply,
		Source:       "inbox",
	}

	if c.Sticker != nil {
		data.Payload = c.Sticker
		data.Message = c.Sticker.Intention
		data.AttachmentType = 21

	} else if c.ProductId > 0 {
		data.ProductId = c.ProductId
		data.ProductProfile = c.ProductProfile
		data.AttachmentType = 3

	} else if c.Voucher != nil {
		data.Payload = c.Voucher
		data.AttachmentType = 11

	} else if c.Invoice != nil {
		data.Payload = c.Invoice
		data.InvoiceLink = c.Invoice
		data.AttachmentType = 7
	}

	return &data
}
