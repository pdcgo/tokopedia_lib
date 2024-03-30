package api_test

import (
	"testing"
	"time"

	"github.com/pdcgo/tokopedia_lib/lib/api"
	"github.com/pdcgo/tokopedia_lib/scenario"
	"github.com/stretchr/testify/assert"
)

func TestCloseShopApi(t *testing.T) {
	tapi, saveSession := scenario.GetTokopediaApiClient()
	defer saveSession()

	now := time.Now()

	t.Run("test close shop schedule libur", func(t *testing.T) {

		hasil, err := tapi.CloseShopSchedule(&api.CloseShopScheduleInput{
			Action:     0,
			CloseStart: int(now.Unix()),
			CloseEnd:   int(now.AddDate(0, 0, 7).Unix()),
		})
		assert.Nil(t, err)
		assert.NotEmpty(t, hasil)
		assert.True(t, hasil.Data.CloseShopSchedule.Success)
		assert.Contains(t, hasil.Data.CloseShopSchedule.Message, "toko libur disimpan")
	})

	t.Run("test close shop schedule buka", func(t *testing.T) {

		hasil, err := tapi.CloseShopSchedule(&api.CloseShopScheduleInput{
			Action: 1,
		})
		assert.Nil(t, err)
		assert.NotEmpty(t, hasil)
		assert.True(t, hasil.Data.CloseShopSchedule.Success)
		assert.Contains(t, hasil.Data.CloseShopSchedule.Message, "sudah buka kembali")
	})
}
