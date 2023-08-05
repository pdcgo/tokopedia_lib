package api_public

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/pdcgo/common_conf/pdc_common"
)

type GraphqlPayload struct {
	OperationName string      `json:"operationName"`
	Variables     interface{} `json:"variables"`
	Query         string      `json:"query"`
}

func (api *TokopediaApiPublic) graphqlDefaultHeader(req *http.Request) {

	headers := map[string]string{
		"User-Agent":   api.Session.Ua,
		"Content-Type": "application/json",
		"Origin":       "https://www.tokopedia.com",
	}

	for key, value := range headers {
		req.Header.Set(key, value)
	}

}

func (api *TokopediaApiPublic) NewGraphqlReq(payload *GraphqlPayload) *http.Request {
	ur := fmt.Sprintf("https://gql.tokopedia.com/graphql/%s", payload.OperationName)

	dataraw, err := json.Marshal(payload)
	if err != nil {
		pdc_common.ReportError(err)
	}

	req, err := http.NewRequest(http.MethodPost, ur, bytes.NewReader(dataraw))
	if err != nil {
		pdc_common.ReportError(err)
	}
	api.graphqlDefaultHeader(req)

	for _, cookie := range api.Session.Cookies {
		req.AddCookie(cookie)
	}

	return req
}

var ErrGraphqlBatchNoOperationName = errors.New("graphql batch no operation name")
var ErrGraphqlBatchNoQuery = errors.New("graphql batch no query")

func (api *TokopediaApiPublic) NewGraphqlReqBatch(operationName string, payloads []*GraphqlPayload) (*http.Request, error) {

	if operationName == "" {
		return nil, ErrGraphqlBatchNoOperationName
	}

	if len(payloads) == 0 {
		return nil, ErrGraphqlBatchNoQuery
	}

	for _, payload := range payloads {
		payload.OperationName = operationName
	}

	ur := fmt.Sprintf("https://gql.tokopedia.com/graphql/%s", operationName)

	dataraw, err := json.Marshal(payloads)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(http.MethodPost, ur, bytes.NewReader(dataraw))
	if err != nil {
		return req, err
	}
	api.graphqlDefaultHeader(req)

	for _, cookie := range api.Session.Cookies {
		req.AddCookie(cookie)
	}

	return req, nil
}
