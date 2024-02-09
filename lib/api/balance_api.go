package api

type Balance struct {
	SellerAll    int    `json:"seller_all"`
	SellerUsable int    `json:"seller_usable"`
	BuyerUsable  int    `json:"buyer_usable"`
	Typename     string `json:"__typename"`
}

type SellerBalanceResp struct {
	Data struct {
		Balance *Balance `json:"balance"`
	} `json:"data"`
}

func (api *TokopediaApi) GetBalance() (*SellerBalanceResp, error) {
	query := GraphqlPayload{
		OperationName: "GetSellerBalance",
		Variables:     struct{}{},
		Query:         "query GetSellerBalance {\n  balance {\n    seller_all\n    seller_usable\n    buyer_usable\n    __typename\n  }\n}\n",
	}

	req := api.NewGraphqlReq(&query)

	var hasil SellerBalanceResp
	err := api.SendRequest(req, &hasil)

	return &hasil, err
}

type MidasGetAllDepositAmount struct {
	HaveError       bool   `json:"have_error"`
	Message         string `json:"message"`
	BuyerAll        int    `json:"buyer_all"`
	BuyerAllFmt     string `json:"buyer_all_fmt"`
	BuyerHold       int    `json:"buyer_hold"`
	BuyerHoldFmt    string `json:"buyer_hold_fmt"`
	BuyerUsable     int    `json:"buyer_usable"`
	BuyerUsableFmt  string `json:"buyer_usable_fmt"`
	SellerAll       int    `json:"seller_all"`
	SellerAllFmt    string `json:"seller_all_fmt"`
	SellerHold      int    `json:"seller_hold"`
	SellerHoldFmt   string `json:"seller_hold_fmt"`
	SellerUsable    int    `json:"seller_usable"`
	SellerUsableFmt string `json:"seller_usable_fmt"`
	Typename        string `json:"__typename"`
}

type WithdrawBalanceResp struct {
	Data struct {
		MidasGetAllDepositAmount *MidasGetAllDepositAmount `json:"MidasGetAllDepositAmount"`
	} `json:"data"`
}

func (api *TokopediaApi) WithDrawBalance() (*WithdrawBalanceResp, error) {
	query := GraphqlPayload{
		OperationName: "GetBalance",
		Variables:     struct{}{},
		Query:         "query GetBalance($isAdmin: Boolean) {\n  MidasGetAllDepositAmount(is_admin: $isAdmin) {\n    have_error\n    message\n    buyer_all\n    buyer_all_fmt\n    buyer_hold\n    buyer_hold_fmt\n    buyer_usable\n    buyer_usable_fmt\n    seller_all\n    seller_all_fmt\n    seller_hold\n    seller_hold_fmt\n    seller_usable\n    seller_usable_fmt\n    __typename\n  }\n}\n",
	}

	req := api.NewGraphqlReq(&query)

	var hasil WithdrawBalanceResp
	err := api.SendRequest(req, &hasil)

	return &hasil, err
}

type SaldoQueryVariable struct {
	IsAdmin bool `json:"isAdmin"`
}

type RichieBalanceWD struct {
	BuyerAll        int    `json:"buyer_all"`
	BuyerHold       int    `json:"buyer_hold"`
	BuyerUsable     int    `json:"buyer_usable"`
	BuyerAllFmt     string `json:"buyer_all_fmt"`
	BuyerHoldFmt    string `json:"buyer_hold_fmt"`
	BuyerUsableFmt  string `json:"buyer_usable_fmt"`
	SellerAll       int    `json:"seller_all"`
	SellerHold      int    `json:"seller_hold"`
	SellerUsable    int    `json:"seller_usable"`
	SellerAllFmt    string `json:"seller_all_fmt"`
	SellerHoldFmt   string `json:"seller_hold_fmt"`
	SellerUsableFmt string `json:"seller_usable_fmt"`
	HaveError       bool   `json:"have_error"`
	Typename        string `json:"__typename"`
}

type SaldoQueryResp struct {
	Data struct {
		RichieBalanceWD *RichieBalanceWD `json:"RichieBalanceWD"`
	} `json:"data"`
}

func (api *TokopediaApi) saldoQuery(payload *SaldoQueryVariable) (*RichieBalanceWD, error) {
	query := GraphqlPayload{
		OperationName: "saldoQuery",
		Variables:     payload,
		Query:         "",
	}

	req := api.NewGraphqlReq(&query)

	var hasil RichieBalanceWD
	err := api.SendRequest(req, &hasil)

	return &hasil, err
}

func (api *TokopediaApi) SaldoQuery(isAdmin bool) (*RichieBalanceWD, error) {
	payload := &SaldoQueryVariable{
		IsAdmin: isAdmin,
	}

	return api.saldoQuery(payload)
}
