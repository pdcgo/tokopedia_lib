package api_test

import (
	"testing"
	"time"

	"github.com/pdcgo/tokopedia_lib"
	"github.com/pdcgo/tokopedia_lib/lib/api"
	"github.com/stretchr/testify/assert"
)

func TestMidasGetDepositHistory(t *testing.T) {
	driver, err := tokopedia_lib.NewDriverAccount("SuratiKioss@outlook.com", "@Srengat123", "3LP7SEYTNCQOMREZXBUYPGZL5JRG52BZ")
	assert.Nil(t, err)
	driver.SetPIN("778899")

	tApi, saveSession, err := driver.CreateApi()
	assert.Nil(t, err)
	defer saveSession()

	now := time.Now().UTC()
	lastMonth := now.AddDate(0, 0, -31)

	variable := api.NewDepositHistoryVariable()
	variable.DateFrom = lastMonth.Format(api.YYYYMMDD)
	res, err := tApi.MidasGetDepositHistory(variable)
	assert.Nil(t, err)
	assert.NotNil(t, res)
}
