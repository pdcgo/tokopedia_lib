package api

import (
	"encoding/json"
	"time"
)

type ShopScorePenaltyDetailVar struct {
	ShopID    string    `json:"shopID"`
	Page      int       `json:"page"`
	Total     int       `json:"total"`
	StartDate time.Time `json:"startDate"`
	EndDate   time.Time `json:"endDate"`
	Sort      int       `json:"sort"`
	Source    string    `json:"source"`
}

func (u *ShopScorePenaltyDetailVar) MarshalJSON() ([]byte, error) {
	type Alias ShopScorePenaltyDetailVar
	return json.Marshal(&struct {
		StartDate string `json:"startDate"`
		EndDate   string `json:"endDate"`
		*Alias
	}{
		StartDate: u.StartDate.Format("2006-01-02"),
		EndDate:   u.EndDate.Format("2006-01-02"),
		Alias:     (*Alias)(u),
	})
}

type ShopScorePenaltyDetailRes struct {
	Data *ShopScorePenaltyDetailData `json:"data"`
}
type ProductDetail struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Typename string `json:"__typename"`
}
type ShopScorePenaltyDetailResult struct {
	ShopPenaltyID         string        `json:"shopPenaltyID"`
	InvoiceNumber         string        `json:"invoiceNumber"`
	Reason                string        `json:"reason"`
	Score                 int           `json:"score"`
	CreateTime            string        `json:"createTime"`
	TypeID                int           `json:"typeID"`
	TypeName              string        `json:"typeName"`
	PenaltyStartDate      string        `json:"penaltyStartDate"`
	PenaltyExpirationDate string        `json:"penaltyExpirationDate"`
	Status                string        `json:"status"`
	ProductDetail         ProductDetail `json:"productDetail"`
	PenaltyTypeGroup      int           `json:"penaltyTypeGroup"`
	Typename              string        `json:"__typename"`
}
type Error struct {
	Message  string `json:"message"`
	Typename string `json:"__typename"`
}
type ShopScorePenaltyDetail struct {
	Result   []*ShopScorePenaltyDetailResult `json:"result"`
	HasNext  bool                            `json:"hasNext"`
	HasPrev  bool                            `json:"hasPrev"`
	Error    *Error                          `json:"error"`
	Typename string                          `json:"__typename"`
}
type ShopScorePenaltyDetailData struct {
	ShopScorePenaltyDetail *ShopScorePenaltyDetail `json:"shopScorePenaltyDetail"`
}

func (api *TokopediaApi) ShopScorePenaltyDetail(payload *ShopScorePenaltyDetailVar) (*ShopScorePenaltyDetailRes, error) {
	query := GraphqlPayload{
		OperationName: "shopScorePenaltyDetail",
		Variables:     payload,
		Query: `query shopScorePenaltyDetail($page: Int!, $total: Int!, $startDate: String!, $endDate: String!, $source: String!, $sort: Int, $shopID: String!) {
			shopScorePenaltyDetail(input: {page: $page, total: $total, startDate: $startDate, endDate: $endDate, source: $source, sort: $sort, shopID: $shopID}) {
			  result {
				shopPenaltyID
				invoiceNumber
				reason
				score
				createTime
				typeID
				typeName
				penaltyStartDate
				penaltyExpirationDate
				status
				productDetail {
				  id
				  name
				  __typename
				}
				penaltyTypeGroup
				__typename
			  }
			  hasNext
			  hasPrev
			  error {
				message
				__typename
			  }
			  __typename
			}
		  }`,
	}

	req := api.NewGraphqlReq(&query)

	var hasil ShopScorePenaltyDetailRes
	err := api.SendRequest(req, &hasil)

	return &hasil, err
}

type ShopScorePenaltySummaryVar struct {
	StartDate time.Time `json:"startDate"` // "2023-07-14"
	EndDate   time.Time `json:"endDate"`
	ShopID    string    `json:"shopID"`
	Source    string    `json:"source"`
}

func (u *ShopScorePenaltySummaryVar) MarshalJSON() ([]byte, error) {
	type Alias ShopScorePenaltySummaryVar
	data, err := json.Marshal(&struct {
		StartDate string `json:"startDate"`
		EndDate   string `json:"endDate"`
		*Alias
	}{
		StartDate: u.StartDate.Format("2006-01-02"),
		EndDate:   u.EndDate.Format("2006-01-02"),
		Alias:     (*Alias)(u),
	})

	return data, err
}

type ShopScorePenaltySummaryRes struct {
	Data *ShopScorePenaltySummaryData `json:"data"`
}
type ShopScorePenaltySummaryResult struct {
	Penalty       int    `json:"penalty"`
	PenaltyAmount int    `json:"penaltyAmount"`
	Typename      string `json:"__typename"`
}
type ShopScorePenaltySummaryError struct {
	Message  string `json:"message"`
	Typename string `json:"__typename"`
}

func (er *ShopScorePenaltySummaryError) Error() string {
	return er.Message + er.Typename
}

type ShopScorePenaltySummary struct {
	Result   *ShopScorePenaltySummaryResult `json:"result"`
	Error    *ShopScorePenaltySummaryError  `json:"error"`
	Typename string                         `json:"__typename"`
}
type ShopScorePenaltySummaryData struct {
	ShopScorePenaltySummary *ShopScorePenaltySummary `json:"shopScorePenaltySummary"`
}

func (api *TokopediaApi) ShopScorePenaltySummary(payload *ShopScorePenaltySummaryVar) (*ShopScorePenaltySummaryRes, error) {
	query := GraphqlPayload{
		OperationName: "shopScorePenaltySummary",
		Variables:     payload,
		Query: `query shopScorePenaltySummary($startDate: String!, $endDate: String!, $source: String!, $shopID: String!) {
			shopScorePenaltySummary(input: {startDate: $startDate, endDate: $endDate, source: $source, shopID: $shopID}) {
			  result {
				penalty
				penaltyAmount
				__typename
			  }
			  error {
				message
				__typename
			  }
			  __typename
			}
		  }`,
	}

	req := api.NewGraphqlReq(&query)

	var hasil ShopScorePenaltySummaryRes
	err := api.SendRequest(req, &hasil)
	if err != nil {
		return nil, err
	}

	if hasil.Data.ShopScorePenaltySummary.Error.Message != "" {
		return nil, hasil.Data.ShopScorePenaltySummary.Error
	}

	return &hasil, nil

}
