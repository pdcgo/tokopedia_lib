package api_test

import (
	"testing"

	"github.com/pdcgo/tokopedia_lib"
	"github.com/pdcgo/tokopedia_lib/lib/api"
	"github.com/stretchr/testify/assert"
)

func TestApiUser(t *testing.T) {
	driver, err := tokopedia_lib.NewDriverAccount("mantracode@yahoo.com", "Muhammad123!`", "3KPN2WN2LG42IMONAFRMMCNQJGXEEQGD")

	assert.Nil(t, err)

	t.Run("test api tanpa create api", func(t *testing.T) {
		driver.Run(false, func(dctx *tokopedia_lib.DriverContext) error {
			err := driver.SellerLogin(dctx)
			assert.Nil(t, err)
			return nil
		})

		sapi := api.NewTokopediaApi(driver.Session)
		_, err = sapi.IsAutheticated()
		assert.Nil(t, err)

	})

	// api, saveSession, err := driver.CreateApi()
	// assert.Nil(t, err)
	// defer saveSession()

}
