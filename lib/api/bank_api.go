package api

type BankListQueryVariable struct {
	IsAdmin bool `json:"isAdmin"`
}

type WalletAppData struct {
	State          int    `json:"state"`
	Message        string `json:"message"`
	CtaCopyWriting string `json:"cta_copy_writing"`
	CtaLink        string `json:"cta_link"`
	Typename       string `json:"__typename"`
}

type GetBankWDV2 struct {
	AccountNo          string         `json:"accountNo"`
	AccountName        string         `json:"accountName"`
	BankID             int            `json:"bankID"`
	BankName           string         `json:"bankName"`
	BankBranch         int            `json:"bankBranch"`
	BankAccountID      int            `json:"bankAccountID"`
	BankImageURL       string         `json:"bankImageUrl"`
	IsDefaultBank      int            `json:"isDefaultBank"`
	IsVerifiedAccount  int            `json:"isVerifiedAccount"`
	IsFraud            bool           `json:"is_fraud"`
	MinAmount          int            `json:"minAmount"`
	MaxAmount          int            `json:"maxAmount"`
	AdminFee           int            `json:"adminFee"`
	Status             int            `json:"status"`
	HaveRpProgram      bool           `json:"have_rp_program"`
	HaveSpecialOffer   bool           `json:"have_special_offer"`
	DefaultBankAccount bool           `json:"default_bank_account"`
	WarningMessage     string         `json:"warning_message"`
	WarningColor       int            `json:"warning_color"`
	Notes              string         `json:"notes"`
	WalletAppData      *WalletAppData `json:"wallet_app_data"`
	Typename           string         `json:"__typename"`
}

type GopayData struct {
	Limit            string `json:"limit"`
	LimitCopyWriting string `json:"limit_copy_writing"`
	ImageURL         string `json:"image_url"`
	WidgetNote       string `json:"widget_note"`
	BottomsheetData  struct {
		Title       string `json:"title"`
		Description string `json:"description"`
		Balance     string `json:"balance"`
		Typename    string `json:"__typename"`
	} `json:"bottomsheet_data"`
	Typename string `json:"__typename"`
}

type GetBankListWDV2Data []*GetBankWDV2

func (b GetBankListWDV2Data) GetDefaultBank() *GetBankWDV2 {
	for _, bank := range b {
		if bank.DefaultBankAccount {
			return bank
		}
	}

	return nil
}

type GetBankListWDV2 struct {
	Status    int                 `json:"status"`
	Message   string              `json:"message"`
	Data      GetBankListWDV2Data `json:"data"`
	GopayData *GopayData          `json:"gopay_data"`
	Typename  string              `json:"__typename"`
}

type BankListQueryResp struct {
	Data struct {
		GetBankListWDV2 *GetBankListWDV2 `json:"GetBankListWDV2"`
	} `json:"data"`
}

func (api *TokopediaApi) bankListQuery(payload *BankListQueryVariable) (*BankListQueryResp, error) {
	query := GraphqlPayload{
		OperationName: "bankListQuery",
		Variables:     payload,
		Query:         "query bankListQuery($isAdmin: Boolean) {\n  GetBankListWDV2(isAdmin: $isAdmin) {\n    status\n    message\n    data {\n      accountNo\n      accountName\n      bankID\n      bankName\n      bankBranch\n      bankAccountID\n      bankImageUrl\n      isDefaultBank\n      isVerifiedAccount\n      is_fraud\n      minAmount\n      maxAmount\n      adminFee\n      status\n      have_rp_program\n      have_special_offer\n      default_bank_account\n      warning_message\n      warning_color\n      notes\n      wallet_app_data {\n        state\n        message\n        cta_copy_writing\n        cta_link\n        __typename\n      }\n      __typename\n    }\n    gopay_data {\n      limit\n      limit_copy_writing\n      image_url\n      widget_note\n      bottomsheet_data {\n        title\n        description\n        balance\n        __typename\n      }\n      __typename\n    }\n    __typename\n  }\n}\n",
	}

	req := api.NewGraphqlReq(&query)

	var hasil BankListQueryResp
	err := api.SendRequest(req, &hasil)

	return &hasil, err
}

func (api *TokopediaApi) BankListQuery(isAdmin bool) (*BankListQueryResp, error) {
	payload := &BankListQueryVariable{
		IsAdmin: isAdmin,
	}

	return api.bankListQuery(payload)
}
