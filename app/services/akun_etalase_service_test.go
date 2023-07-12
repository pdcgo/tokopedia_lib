package services_test

import (
	"testing"

	"github.com/pdcgo/tokopedia_lib/app/services"
	"github.com/pdcgo/tokopedia_lib/scenario"
	"github.com/stretchr/testify/assert"
)

func TestAkunEtalase(t *testing.T) {
	db := scenario.GetDb()
	mapsrv := services.NewEtalaseMapService(db)
	apiclient, saveSession := scenario.GetTokopediaApiClient()
	defer saveSession()

	akunsrv := services.NewAkunEtalaseService(apiclient, mapsrv)

	mapsrv.AddMap(&services.EtalasePayload{
		Etalase: "Gamis",
		CatIDs:  []int{20},
	})

	err := akunsrv.RefreshShowCase()
	assert.Nil(t, err)

	t.Run("test create etalase", func(t *testing.T) {
		showcase, err := akunsrv.GetEtalase(20)

		assert.NotEmpty(t, showcase)
		assert.Nil(t, err)
	})

	t.Run("test create etalase map not found", func(t *testing.T) {
		_, err := akunsrv.GetEtalase(21)

		assert.NotNil(t, err)
	})

}
