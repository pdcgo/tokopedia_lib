package api_test

import (
	"testing"

	"github.com/pdcgo/tokopedia_lib"
	"github.com/pdcgo/tokopedia_lib/lib/api"
	"github.com/stretchr/testify/assert"
)

func TestApiUser(t *testing.T) {
	driver, err := tokopedia_lib.NewDriverAccount("dirondap@piekiih.com", "Balikpapan1*", "PN72GYGA62UFFJVOKURIA6WVR43FUBSG")

	assert.Nil(t, err)

	t.Run("test api tanpa create api", func(t *testing.T) {
		driver.Run(false, func(dctx *tokopedia_lib.DriverContext) error {
			err := driver.SellerLogin(dctx)
			assert.Nil(t, err)
			return nil
		})

		driver.Run(false, func(dctx *tokopedia_lib.DriverContext) error {

			sapi := api.NewTokopediaApi(driver.Session)
			_, err = sapi.IsAutheticated()
			assert.Nil(t, err)

			return err
		})

	})

	// api, saveSession, err := driver.CreateApi()
	// assert.Nil(t, err)
	// defer saveSession()

}
