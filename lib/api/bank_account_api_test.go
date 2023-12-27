package api_test

import (
	"testing"

	"github.com/pdcgo/tokopedia_lib/scenario"
	"github.com/stretchr/testify/assert"
)

func TestBankAccountApi(t *testing.T) {
	api, saveSession := scenario.GetTokopediaApiClient()
	defer saveSession()

	hasil, err := api.GetBankAccount()
	assert.Nil(t, err)
	assert.NotEmpty(t, hasil)
	assert.Equal(t, len(hasil.Data.GetBankAccount.Data.BankAccounts), 2)
}
