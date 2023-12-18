package api_test

import (
	"testing"

	"github.com/pdcgo/tokopedia_lib/scenario"
	"github.com/stretchr/testify/assert"
)

func TestInvoiceApi(t *testing.T) {

	api, saveSession := scenario.GetTokopediaApiClient()
	defer saveSession()

	hasil, err := api.GetInvoiceV3("INV/20231004/MPL/3491982476")
	assert.Nil(t, err)
	assert.NotEmpty(t, hasil)
	assert.NotEmpty(t, hasil.Data.GetInvoiceV3.Invoices)
	assert.NotEmpty(t, hasil.Data.GetInvoiceV3.Invoices[0].PaymentData)
	assert.NotEmpty(t, hasil.Data.GetInvoiceV3.Invoices[0].OrderData)
	assert.False(t, hasil.IsCod())
}
