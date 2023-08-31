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
		"Content-Type":        "application/json",
		"Origin":              "https://www.tokopedia.com",
		"Sec-Ch-Ua":           `"Chromium";v="116", "Not)A;Brand";v="24", "Google Chrome";v="116"`,
		"Sec-Ch-Ua-Mobile":    "?0",
		"Sec-Ch-Ua-Platform":  `"Windows"`,
		"User-Agent":          "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/116.0.0.0 Safari/537.36",
		"X-Device":            "desktop",
		"X-Source":            "tokopedia-lite",
		"X-Tkpd-Lite-Service": "zeus",
		"X-Version":           "b078a1e",
		"Pragma":              "no-cache",
		"Cache-Control":       "no-cache",
		"Sec-Fetch-Site":      "same-site",
		"Sec-Fetch-Mode":      "cors",
		"Sec-Fetch-Dest":      "empty",
		"Referer":             "https://www.tokopedia.com/",
		// "Accept-Encoding":     "gzip, deflate, br",
		"Accept-Language": "en,en-US;q=0.9,id-ID;q=0.8,id;q=0.7",
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

	return req, nil
}
