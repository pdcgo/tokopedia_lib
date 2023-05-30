package upload_app_test

import (
	"testing"

	"github.com/pdcgo/tokopedia_lib/app/upload_app"
	"github.com/pdcgo/tokopedia_lib/scenario"
	"github.com/stretchr/testify/assert"
)

func TestAkunIterate(t *testing.T) {
	db := scenario.GetDb()

	status := upload_app.AkunUploadStatus{
		LimitUpload: 200,
		Active:      true,
	}
	db.Save(&upload_app.AkunItem{
		Username:         "pucung@gmail.com",
		AkunUploadStatus: status,
	})

	db.Save(&upload_app.AkunItem{
		Username:         "kinanthi@gmail.com",
		AkunUploadStatus: status,
	})
	db.Save(&upload_app.AkunItem{
		Username:         "mijil@gmail.com",
		AkunUploadStatus: status,
	})
	db.Save(&upload_app.AkunItem{
		Username:         "gambuh@gmail.com",
		AkunUploadStatus: status,
	})

	iter := upload_app.NewAkunUploadIterator(db)
	akun, _, _, err := iter.Get()
	assert.NotEmpty(t, akun)
	assert.Nil(t, err)

	t.Run("akun kedua harus beda", func(t *testing.T) {

		akun2, _, _, err := iter.Get()
		assert.NotEmpty(t, akun)
		assert.Nil(t, err)
		assert.NotEqual(t, akun2.Username, akun.Username)
	})

	t.Run("test merata bergiliran", func(t *testing.T) {
		akunt := akun
		for c := 0; c < 10; c += 1 {

			akunn, update, _, err := iter.Get()
			errup := update(1, nil)
			assert.Nil(t, errup)
			t.Log(akunn.Username)

			assert.NotEqual(t, akunt.Username, akunn.Username)

			akunt = akunn
			assert.NotEmpty(t, akun)
			assert.Nil(t, err)

		}
	})

	t.Run("getting upload report", func(t *testing.T) {
		data, err := iter.GetStatus()
		t.Log("data status", data)
		assert.NotEmpty(t, data)
		assert.Nil(t, err)
	})

	t.Run("test loop akun tidak ada akun lagi tpi ada task", func(t *testing.T) {
		t.Error("not implemented")
	})
}
