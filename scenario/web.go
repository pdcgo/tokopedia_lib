package scenario

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/schema"
	"github.com/pdcgo/v2_gots_sdk"
	"github.com/pdcgo/v2_gots_sdk/pdc_api"
	"github.com/stretchr/testify/assert"
)

func CreateSdk() *v2_gots_sdk.ApiSdk {
	r := gin.Default()

	return v2_gots_sdk.NewApiSdk(r)
}

type SendRequest func(req *pdc_api.Api) *httptest.ResponseRecorder
type RunHandler func(sdk *v2_gots_sdk.ApiSdk, sendApi SendRequest) error

var encoder = schema.NewEncoder()

func RunWebSdk(t *testing.T, handler RunHandler) {
	sdk := CreateSdk()

	var sendApi SendRequest = func(api *pdc_api.Api) *httptest.ResponseRecorder {
		w := httptest.NewRecorder()

		data := bytes.NewBuffer(nil)
		if api.Payload != nil {
			err := json.NewEncoder(data).Encode(api.Payload)
			assert.Nil(t, err)
		}

		req, err := http.NewRequest(api.Method, api.RelativePath, data)
		assert.Nil(t, err)

		if api.Query != nil {
			q := req.URL.Query()
			encoder.Encode(api.Query, q)
			req.URL.RawQuery = q.Encode()
		}

		sdk.R.ServeHTTP(w, req)
		return w
	}

	err := handler(sdk, sendApi)
	assert.Nil(t, err)

}
