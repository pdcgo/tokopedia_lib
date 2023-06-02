package api_test

import (
	"testing"

	"github.com/pdcgo/tokopedia_lib/scenario"
	"github.com/stretchr/testify/assert"
)

func TestNotificationApi(t *testing.T) {

	api, saveSession := scenario.GetTokopediaApiClient()
	defer saveSession()

	hasil, err := api.NotificationCounter()
	assert.NotEmpty(t, hasil)
	assert.Nil(t, err)
}
