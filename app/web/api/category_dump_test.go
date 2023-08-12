package api_test

import (
	"net/http"
	"testing"

	"github.com/pdcgo/tokopedia_lib/app/web/api"
	"github.com/pdcgo/tokopedia_lib/lib/api_public"
	"github.com/pdcgo/tokopedia_lib/scenario"
	"github.com/pdcgo/v2_gots_sdk"
	"github.com/stretchr/testify/assert"
)

func TestApiCategoryDump(t *testing.T) {
	scen := scenario.NewScenario(t)
	pubapi, err := api_public.NewTokopediaApiPublic()
	assert.Nil(t, err)

	scenario.RunWebSdk(t, func(sdk *v2_gots_sdk.ApiSdk, sendApi scenario.SendRequest) error {
		scen.WithBase(func(dirbase string, scen *scenario.Scenario) {

			dapi := api.NewCategoryDumpApi(scen, pubapi)
			dapi.RegisterApi(sdk.Group("category"))

			t.Run("test category dump", func(t *testing.T) {
				res := sendApi(&v2_gots_sdk.Api{
					Method:       http.MethodGet,
					RelativePath: "/category/category_dump",
				})

				assert.Equal(t, http.StatusOK, res.Result().StatusCode)

			})

		})

		return nil
	})
}
