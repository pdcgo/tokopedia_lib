package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/pdcgo/common_conf/pdc_common"
)

type GraphqlPayload struct {
	OperationName string      `json:"operationName"`
	Variables     interface{} `json:"variables"`
	Query         string      `json:"query"`
}

func (api *TokopediaApi) graphqlDefaultHeader(req *http.Request) {

	headers := map[string]string{
		"User-Agent":   api.Session.UserAgent(),
		"Content-Type": "application/json",
		"Origin":       "https://seller.tokopedia.com",
		"Accept":       "*/*",
	}

	for key, value := range headers {
		req.Header.Set(key, value)
	}

}

func (api *TokopediaApi) NewGraphqlReq(payload *GraphqlPayload) *http.Request {
	ur := fmt.Sprintf("https://gql.tokopedia.com/graphql/%s", payload.OperationName)

	dataraw, err := json.Marshal([]*GraphqlPayload{payload})
	if err != nil {
		pdc_common.ReportError(err)
	}

	req, err := http.NewRequest(http.MethodPost, ur, bytes.NewReader(dataraw))
	if err != nil {
		pdc_common.ReportError(err)
	}
	api.graphqlDefaultHeader(req)
	api.Session.AddToHttpRequest(req)

	return req
}
