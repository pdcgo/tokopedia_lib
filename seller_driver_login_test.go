package tokopedia_lib_test

import (
	"testing"

	"github.com/pdcgo/tokopedia_lib"
	"github.com/stretchr/testify/assert"
)

func TestSellerLogin(t *testing.T) {
	t.Run("test login otp salah", func(t *testing.T) {
		driver, err := tokopedia_lib.NewDriverAccount("kedaiblanjadotcom@gmail.com", "pakdosen", "OF5LGY2FLIDPUZXK7GST5R3HT6QA6B5F")
		assert.Nil(t, err)

		err = driver.Run(false, func(dctx *tokopedia_lib.DriverContext) error {
			return driver.SellerLogin(dctx)
		})

		assert.Nil(t, err)

	})
}
