package api_public_test

import (
	"testing"

	"github.com/pdcgo/tokopedia_lib"
	"github.com/pdcgo/tokopedia_lib/lib/api_public"
	"github.com/pdcgo/tokopedia_lib/lib/model_public"
	"github.com/stretchr/testify/assert"
)

func TestShopCoreInfo(t *testing.T) {

	pSession := tokopedia_lib.NewSessionPublic()
	api := api_public.NewTokopediaApiPublic(pSession)

	t.Run("test shopCoreInfo", func(t *testing.T) {
		variable := model_public.ShopCoreInfoVar{
			ID:     0,
			Domain: "lenovo-tangerang",
		}
		hasil, err := api.ShopCoreInfo(&variable)
		assert.Nil(t, err)
		assert.NotEmpty(t, hasil)
	})
}

func TestShopStatisticQuery(t *testing.T) {

	pSession := tokopedia_lib.NewSessionPublic()
	api := api_public.NewTokopediaApiPublic(pSession)

	t.Run("test shopCoreInfo", func(t *testing.T) {
		variable := model_public.ShopStatisticQueryVar{
			ShopID:    11534215,
			ShopIDStr: "11534215",
		}
		hasil, err := api.ShopStatisticQuery(&variable)
		assert.Nil(t, err)
		assert.NotEmpty(t, hasil)
	})
}

func TestGetShopOperationalHourStatus(t *testing.T) {

	pSession := tokopedia_lib.NewSessionPublic()
	api := api_public.NewTokopediaApiPublic(pSession)

	t.Run("test shopCoreInfo", func(t *testing.T) {
		variable := model_public.ShopIdVar{
			ShopID: "11534215",
		}
		hasil, err := api.GetShopOperationalHourStatus(&variable)
		assert.Nil(t, err)
		assert.NotEmpty(t, hasil)
	})
}

func TestVoucherListQuery(t *testing.T) {

	pSession := tokopedia_lib.NewSessionPublic()
	api := api_public.NewTokopediaApiPublic(pSession)

	t.Run("test shopCoreInfo", func(t *testing.T) {
		variable := model_public.ShopIdVarInt{
			ShopID: 11534215,
		}
		hasil, err := api.VoucherListQuery(&variable)
		assert.Nil(t, err)
		assert.NotEmpty(t, hasil)
	})
}

func TestShopNote(t *testing.T) {

	pSession := tokopedia_lib.NewSessionPublic()
	api := api_public.NewTokopediaApiPublic(pSession)

	t.Run("test shopCoreInfo", func(t *testing.T) {
		variable := model_public.ShopNoteVar{
			ID:  "0",
			Sid: "11534215",
		}
		hasil, err := api.ShopNote(&variable)
		assert.Nil(t, err)
		assert.NotEmpty(t, hasil)
	})
}
