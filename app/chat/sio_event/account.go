package sio_event

type WithdrawEvent struct {
	Name    string `json:"name"`
	Type    string `json:"type"`
	Message string `json:"message"`
}

type AccountWithdrawEvent struct {
	Shopid int            `json:"shopid,string"`
	Event  *WithdrawEvent `json:"event"`
}
