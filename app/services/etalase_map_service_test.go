package services_test

import (
	"testing"

	"github.com/pdcgo/tokopedia_lib/app/services"
	"github.com/pdcgo/tokopedia_lib/scenario"
	"github.com/stretchr/testify/assert"
)

func TestEtalase(t *testing.T) {
	db := scenario.GetDb()

	service := services.NewEtalaseMapService(db)

	t.Run("testing add etalase", func(t *testing.T) {
		err := service.AddMap(&services.EtalasePayload{
			Etalase: "test_etalase",
			CatIDs:  []int{1, 2, 3, 4},
		})

		assert.Nil(t, err)

		err = service.AddMap(&services.EtalasePayload{
			Etalase: "test_etalase",
			CatIDs:  []int{1, 2},
		})

		assert.Nil(t, err)
	})

	t.Run("list etalase", func(t *testing.T) {
		hasil, err := service.ListEtalase()

		assert.NotEmpty(t, hasil)
		assert.Nil(t, err)
	})

	t.Run("get etalase from cat id", func(t *testing.T) {
		item, err := service.GetEtalase(1)
		assert.NotEmpty(t, item)
		assert.Nil(t, err)
	})

	t.Run("delete etalase", func(t *testing.T) {
		err := service.DeleteEtalase("test_etalase")
		assert.Nil(t, err)

		_, err = service.GetEtalase(1)
		assert.NotNil(t, err)
	})
}
