package api_public

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

func (api *TokopediaApiPublic) graphqlDefaultHeader(req *http.Request) {

	headers := map[string]string{
		"User-Agent":   api.Session.Ua,
		"Content-Type": "application/json",
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
