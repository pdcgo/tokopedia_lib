package cek_bot_test

import (
	"testing"

	"github.com/pdcgo/tokopedia_lib"
	"github.com/pdcgo/tokopedia_lib/app/cek_bot"
	"github.com/pdcgo/tokopedia_lib/lib/report"
	"github.com/stretchr/testify/assert"
)

func TestCheckbot(t *testing.T) {
	acdriver, err := tokopedia_lib.NewDriverAccount("pdcthoni@gmail.com", "SilentIsMyMantra", "IULIWGH6TIK3CZBKHGE27DBRLQ5LR5WQ")
	assert.Nil(t, err)
	driver := report.CekReport{
		DriverAccount: acdriver,
	}
	cek_bot.Waitallakun.Add(1)
	cek_bot.Cekbot(&driver)

	cek_bot.Waitallakun.Wait()

	t.Log(driver)
}
