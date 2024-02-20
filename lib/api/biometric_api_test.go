package api_test

import (
	"testing"

	"github.com/pdcgo/tokopedia_lib/scenario"
	"github.com/stretchr/testify/assert"
)

func TestBiometricBeginRegister(t *testing.T) {
	api, saveSession := scenario.GetTokopediaApiClient()
	defer saveSession()

	hasil, err := api.BiometricBeginRegister()
	assert.NotEmpty(t, hasil)
	assert.Nil(t, err)
}
