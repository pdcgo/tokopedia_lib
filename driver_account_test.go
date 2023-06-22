package tokopedia_lib_test

import (
	"testing"

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
		assert.Nil(t, err)

	})

}

func TestMitraLogin(t *testing.T) {
	driver, err := tokopedia_lib.NewDriverAccount("lamarbarton732@outlook.com", "vytTmjT8", "REQXXGY5NXJPMFHEKEKXFEBO46H2NMHU")
	assert.Nil(t, err)
	driver.Run(false, func(dctx *tokopedia_lib.DriverContext) error {
		return driver.MitraLogin(dctx.Ctx)
	})
}
