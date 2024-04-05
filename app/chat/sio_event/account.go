package sio_event

type WithdrawEvent struct {
	Name    string `json:"name"`
	Type    string `json:"type"`
	Message string `json:"message"`
}

func NewWithdrawEvent(name string) *WithdrawEvent {
	return &WithdrawEvent{
		Name:    name,
		Type:    "success",
		Message: "success",
	}
}

func (e *WithdrawEvent) SetError(err error) {
	e.Type = "error"
	e.Message = err.Error()
}

type AccountWithdrawEvent struct {
	Shopid int            `json:"shopid,string"`
	Event  *WithdrawEvent `json:"event"`
}

type AccountActiveEvent struct {
	Shopid int `json:"shopid,string"`
}

func NewAccountActiveEvent(shopid int) *AccountActiveEvent {
	return &AccountActiveEvent{
		Shopid: shopid,
	}
}
