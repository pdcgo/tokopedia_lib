package api_test

import (
	"net/http"
	"testing"

	"github.com/pdcgo/tokopedia_lib/app/web/api"
	"github.com/pdcgo/tokopedia_lib/scenario"
	"github.com/pdcgo/v2_gots_sdk"
	"github.com/pdcgo/v2_gots_sdk/pdc_api"
	"github.com/stretchr/testify/assert"
)

func TestDataFilterApi(t *testing.T) {
	scenario.RunWebSdk(t, func(sdk *v2_gots_sdk.ApiSdk, sendApi scenario.SendRequest) error {
		scen := scenario.NewScenario(t)

		dapi := api.NewDataFilterApi(scen)
		dapi.RegisterApi(sdk.Group("filter"))

		t.Run("test getting filter kota", func(t *testing.T) {
			res := sendApi(&pdc_api.Api{
				Method:       http.MethodGet,
				RelativePath: "/filter/fcity",
			})

			assert.Equal(t, http.StatusOK, res.Result().StatusCode)

		})

		return nil
	})
}
