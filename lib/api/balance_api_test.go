package api_test

import (
	"testing"

	"github.com/pdcgo/tokopedia_lib/scenario"
	"github.com/stretchr/testify/assert"
)

func TestGetBalance(t *testing.T) {
	api, saveSession := scenario.GetTokopediaApiClient()
	defer saveSession()

	hasil, err := api.GetBalance()
	assert.NotEmpty(t, hasil)
	assert.Nil(t, err)
}

func TestWithdrawBalance(t *testing.T) {
	api, saveSession := scenario.GetTokopediaApiClient()
	defer saveSession()

	hasil, err := api.WithDrawBalance()
	assert.NotEmpty(t, hasil)
	assert.Nil(t, err)
}

// func TestSaldoQuery(t *testing.T) {
// 	api, saveSession := scenario.GetTokopediaApiClient()
// 	defer saveSession()

// 	hasil, err := api.SaldoQuery(false)
// 	assert.NotEmpty(t, hasil)
// 	assert.Nil(t, err)
// }
