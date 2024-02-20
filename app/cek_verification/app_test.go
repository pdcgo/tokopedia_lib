package cek_verification_test

import (
	"testing"

	"github.com/pdcgo/tokopedia_lib"
	"github.com/pdcgo/tokopedia_lib/app/cek_verification"
	"github.com/pdcgo/tokopedia_lib/lib/report"
	"github.com/stretchr/testify/assert"
)

func TestCheckVerif(t *testing.T) {
	dd, err := tokopedia_lib.NewDriverAccount("jimmiebowman418@yahoo.com", "jmupzYcn", "DHOUPLOFYIKL6AYAAJKRMDV5X7HQ5PBZ")
	assert.Nil(t, err)

	t.Run("test check verif", func(t *testing.T) {
		driver := report.CekVerifReport{
			DriverAccount: dd,
		}
		err = cek_verification.CheckVerif(&driver)
		assert.Nil(t, err)
	})

	t.Run("test check verif v2", func(t *testing.T) {
		driver := report.CekVerifReport{
			DriverAccount: dd,
		}
		err = cek_verification.CheckVerifV2(&driver)
		assert.Nil(t, err)
		assert.Equal(t, driver.Status, "success")
	})
}
