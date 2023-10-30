package tokopedia_lib_test

import (
	"context"
	"testing"
	"time"

	"github.com/pdcgo/tokopedia_lib"
	"github.com/stretchr/testify/assert"
)

func TestDriverAccount(t *testing.T) {

	t.Run("testing create api", func(t *testing.T) {
		driver, err := tokopedia_lib.NewDriverAccount("pdcthoni@gmail.com", "SilentIsMyMantra", "IULIWGH6TIK3CZBKHGE27DBRLQ5LR5WQ")
		assert.Nil(t, err)

		// t.Run("test login", func(t *testing.T) {

		// 	driver.Run(false, func(dctx *tokopedia_lib.DriverContext) error {

		// 		return driver.MitraLogin(dctx.Ctx)
		// 	})
		// })

		api, saveSession, err := driver.CreateApi()
		defer saveSession()

		assert.NotEmpty(t, api)
		assert.NotEmpty(t, api.AuthenticatedData.UserShopInfo)
		assert.Nil(t, err)

	})

	t.Run("test create api otp salah", func(t *testing.T) {

		driver, err := tokopedia_lib.NewDriverAccount("kedaiblanjadotcom@gmail.com", "pakdosen", "OF5LGY2FLIDPUZXK7GST5R3HT6QA6B5F")
		assert.Nil(t, err)

		api, saveSession, err := driver.CreateApi()
		defer saveSession()

		assert.NotEmpty(t, api)
		assert.NotEmpty(t, api.AuthenticatedData.UserShopInfo)
		assert.Nil(t, err)
	})

}

func TestWithParentContext(t *testing.T) {
	driver, err := tokopedia_lib.NewDriverAccount("pdcthoni@gmail.com", "SilentIsMyMantra", "IULIWGH6TIK3CZBKHGE27DBRLQ5LR5WQ")
	assert.Nil(t, err)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	driver.ParentCtx = ctx

	driver.Run(false, func(dctx *tokopedia_lib.DriverContext) error {
		tc := time.After(time.Minute)

		select {
		case <-dctx.Ctx.Done():
			return nil
		case <-tc:
			t.Error("context parent tidak berguna")
			return nil
		}

	})

}

func TestMitraLogin(t *testing.T) {
	driver, err := tokopedia_lib.NewDriverAccount("lamarbarton732@outlook.com", "vytTmjT8", "REQXXGY5NXJPMFHEKEKXFEBO46H2NMHU")
	assert.Nil(t, err)
	driver.Run(false, func(dctx *tokopedia_lib.DriverContext) error {
		return driver.MitraLogin(dctx.Ctx)
	})
}
