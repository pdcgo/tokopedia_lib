package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/pdcgo/common_conf/pdc_common"
)

type GraphqlPayload struct {
	OperationName string      `json:"operationName"`
	Variables     interface{} `json:"variables"`
	Query         string      `json:"query"`
}

func (api *TokopediaApi) graphqlDefaultHeader(req *http.Request) {

	headers := map[string]string{
		"Origin": "https://seller.tokopedia.com",

		"User-Agent":   api.Session.UserAgent(),
		"Content-Type": "application/json",
	}

	for key, value := range headers {
		req.Header.Set(key, value)
	}

}

func (api *TokopediaApi) NewGraphqlReq(payload *GraphqlPayload) *http.Request {
	ur := fmt.Sprintf("https://gql.tokopedia.com/graphql/%s", payload.OperationName)

	dataraw, err := json.Marshal(payload)
	datastring := strings.ReplaceAll(string(dataraw), "\t", "")
	if err != nil {
		pdc_common.ReportError(err)
	}
	req, err := http.NewRequest(http.MethodPost, ur, strings.NewReader(datastring))
	if err != nil {
		pdc_common.ReportError(err)
	}
	api.graphqlDefaultHeader(req)
	api.Session.AddToHttpRequest(req)

	return req
}
