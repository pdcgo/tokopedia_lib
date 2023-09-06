package api_test

import (
	"testing"

	"github.com/pdcgo/tokopedia_lib/lib/model"
	"github.com/pdcgo/tokopedia_lib/scenario"
	"github.com/stretchr/testify/assert"
)

func TestStickerApi(t *testing.T) {
	apiSession, saveSession := scenario.GetTokopediaApiClient()
	defer saveSession()

	t.Run("test get sticker group api", func(t *testing.T) {
		hasil, err := apiSession.ChatGetGroupSticker(1)
		assert.Nil(t, err)
		assert.NotEmpty(t, hasil)
	})

	t.Run("test get sticker bundle api", func(t *testing.T) {
		hasil, err := apiSession.ChatGetBundleSticker(&model.ChatGetBundleStickerVar{
			ID:    "f0991c80-8e77-11ea-bb5e-000000000000",
			Limit: 8,
		})
		assert.Nil(t, err)
		assert.NotEmpty(t, hasil)
	})
}
