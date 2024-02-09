package api

import (
	"encoding/json"
	"net/http"
	"strings"
	"time"
)

type DepositHistoryVariable struct {
	Page      int    `json:"page"`
	MaxRows   int    `json:"maxRows"`
	DateTo    string `json:"dateTo"`
	DateFrom  string `json:"dateFrom"`
	SaldoType int    `json:"saldoType"`
	IsAdmin   bool   `json:"isAdmin"`
}

var YYYYMMDD = "2006-01-02"

func NewDepositHistoryVariable() *DepositHistoryVariable {
	now := time.Now().UTC()
	lastweek := now.AddDate(0, 0, -7)

	return &DepositHistoryVariable{
		Page:      1,
		MaxRows:   20,
		DateTo:    now.Format(YYYYMMDD),
		DateFrom:  lastweek.Format(YYYYMMDD),
		SaldoType: 2,
		IsAdmin:   false,
	}
}

type DepositContent struct {
	DepositID              int64  `json:"deposit_id"`
	TypeDescription        string `json:"type_description"`
	Type                   int    `json:"type"`
	Class                  string `json:"class"`
	Amount                 int    `json:"amount"`
	AmountFmt              string `json:"amount_fmt"`
	Note                   string `json:"note"`
	CreateTime             string `json:"create_time"`
	WithdrawalDate         string `json:"withdrawal_date"`
	WithdrawalStatus       int    `json:"withdrawal_status"`
	Saldo                  int    `json:"saldo"`
	SaldoFmt               string `json:"saldo_fmt"`
	Image                  string `json:"image"`
	WithdrawalStatusString string `json:"withdrawal_status_string"`
	WithdrawalStatusColor  int    `json:"withdrawal_status_color"`
	WithdrawalID           int    `json:"withdrawal_id"`
	HaveDetail             bool   `json:"have_detail"`
	DetailType             int    `json:"detail_type"`
	Typename               string `json:"__typename"`
}

type MidasGetDepositHistory struct {
	HaveError     bool              `json:"have_error"`
	Status        int               `json:"status"`
	MessageStatus string            `json:"message_status"`
	HaveNextPage  bool              `json:"have_next_page"`
	Content       []*DepositContent `json:"content"`
	Typename      string            `json:"__typename"`
}

type MidasGetDepositHistoryData struct {
	MidasGetDepositHistory *MidasGetDepositHistory `json:"MidasGetDepositHistory"`
}

type MidasGetDepositHistoryResp struct {
	Data *MidasGetDepositHistoryData `json:"data"`
}

func (m *MidasGetDepositHistoryResp) GetContent() []*DepositContent {
	return m.Data.MidasGetDepositHistory.Content
}

// Api usage for getting data deposit hitories.
//
// maximum date range is 31 days
func (api *TokopediaApi) MidasGetDepositHistory(payload *DepositHistoryVariable) (*MidasGetDepositHistoryResp, error) {
	query := &GraphqlPayload{
		OperationName: "GetHistoryList",
		Variables:     payload,
		Query:         "query GetHistoryList($page: Int!, $maxRows: Int!, $dateFrom: String!, $dateTo: String!, $saldoType: Int!, $isAdmin: Boolean) {\n  MidasGetDepositHistory(params: {date_from: $dateFrom, date_to: $dateTo, saldo_type: $saldoType, page: $page, max_rows: $maxRows, is_admin: $isAdmin}) {\n    have_error\n    status\n    message_status\n    have_next_page\n    content {\n      deposit_id\n      type_description\n      type\n      class\n      amount\n      amount_fmt\n      note\n      create_time\n      withdrawal_date\n      withdrawal_status\n      saldo\n      saldo_fmt\n      image\n      withdrawal_status_string\n      withdrawal_status_color\n      withdrawal_id\n      have_detail\n      detail_type\n      __typename\n    }\n    __typename\n  }\n}\n",
	}

	ur := "https://gql.tokopedia.com/graphql/midas/MidasGetDepositHistory"

	dataraw, err := json.Marshal(query)
	if err != nil {
		return nil, err
	}
	datastring := strings.ReplaceAll(string(dataraw), "\t", "")
	req, err := http.NewRequest(http.MethodPost, ur, strings.NewReader(datastring))
	if err != nil {
		return nil, err
	}
	api.graphqlDefaultHeader(req)
	api.Session.AddToHttpRequest(req)

	var hasil MidasGetDepositHistoryResp
	err = api.SendRequest(req, &hasil)

	return &hasil, err
}
