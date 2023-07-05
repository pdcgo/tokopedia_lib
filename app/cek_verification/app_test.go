package cek_verification_test

import (
	"testing"

	"github.com/pdcgo/tokopedia_lib"
	"github.com/pdcgo/tokopedia_lib/app/cek_verification"
	"github.com/stretchr/testify/assert"
)

func TestCheckVerif(t *testing.T) {
	dd, err := tokopedia_lib.NewDriverAccount("jimmiebowman418@yahoo.com", "jmupzYcn", "DHOUPLOFYIKL6AYAAJKRMDV5X7HQ5PBZ")
	driver := cek_verification.VerifDriverAccount{
		DriverAccount: dd,
	}
	assert.Nil(t, err)
	driver.Run(false, func(dctx *tokopedia_lib.DriverContext) error {
		err := driver.MitraLogin(dctx.Ctx)
		assert.Nil(t, err)

		err = driver.CheckVerif(dctx)
		assert.Nil(t, err)

		return nil
	})
}
