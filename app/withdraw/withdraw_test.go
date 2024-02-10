package withdraw_test

import (
	"testing"

	"github.com/pdcgo/tokopedia_lib"
	"github.com/pdcgo/tokopedia_lib/app/withdraw"
	"github.com/stretchr/testify/assert"
)

func TestWithdraw(t *testing.T) {
	driver, err := tokopedia_lib.NewDriverAccount("SuratiKioss@outlook.com", "@Srengat123", "3LP7SEYTNCQOMREZXBUYPGZL5JRG52BZ")
	assert.Nil(t, err)
	driver.SetPIN("778899")

	_, err = withdraw.RunWithdraw([]*tokopedia_lib.DriverAccount{driver})
	if err != nil {
		assert.ErrorIs(t, err, withdraw.ErrSaldoKosong)
		return
	}
	assert.Nil(t, err)
}
