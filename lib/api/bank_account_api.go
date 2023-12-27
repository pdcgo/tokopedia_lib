package api

import "github.com/pdcgo/tokopedia_lib/lib/model"

type BankAccount struct {
	AccID        int    `json:"accID"`
	AccName      string `json:"accName"`
	AccNumber    string `json:"accNumber"`
	BankID       int    `json:"bankID"`
	BankName     string `json:"bankName"`
	BankImageURL string `json:"bankImageUrl"`
	Fsp          int    `json:"fsp"`
	StatusFraud  int    `json:"statusFraud"`
	CopyWriting  string `json:"copyWriting"`
	Typename     string `json:"__typename"`
}

type BankAccountUserInfo struct {
	Message    string `json:"message"`
	IsVerified bool   `json:"isVerified"`
	Typename   string `json:"__typename"`
}

type GetBankAccountDataInfo struct {
	BankAccounts []*BankAccount      `json:"bankAccounts"`
	UserInfo     BankAccountUserInfo `json:"userInfo"`
	Typename     string              `json:"__typename"`
}

type GetBankAccount struct {
	Status   string                  `json:"status"`
	Header   *model.Header           `json:"header"`
	Data     *GetBankAccountDataInfo `json:"data"`
	Typename string                  `json:"__typename"`
}

type GetBankAccountData struct {
	GetBankAccount *GetBankAccount `json:"GetBankAccount"`
}

type GetBankAccountRes struct {
	Data *GetBankAccountData `json:"data"`
}

func (api *TokopediaApi) GetBankAccount() (*GetBankAccountRes, error) {
	query := GraphqlPayload{
		OperationName: "GetBankAccount",
		Variables:     map[string]any{},
		Query: `query GetBankAccount {
			GetBankAccount {
			  status
			  header {
				processTime
				message
				reason
				errorCode
				__typename
			  }
			  data {
				bankAccounts {
				  accID
				  accName
				  accNumber
				  bankID
				  bankName
				  bankImageUrl
				  fsp
				  statusFraud
				  copyWriting
				  __typename
				}
				userInfo {
				  message
				  isVerified
				  __typename
				}
				__typename
			  }
			  __typename
			}
		  }
		  `,
	}

	req := api.NewGraphqlReq(&query)

	var hasil GetBankAccountRes
	err := api.SendRequest(req, &hasil)
	if err != nil {
		return nil, err
	}

	return &hasil, nil

}
