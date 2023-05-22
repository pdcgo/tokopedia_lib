package api_test

import (
	"testing"

	"github.com/pdcgo/tokopedia_lib"
	"github.com/stretchr/testify/assert"
)

func TestNotificationApi(t *testing.T) {
	driver, _ := tokopedia_lib.NewDriverAccount("bethdunn892@outlook.com", "MZT2Zk8U", "FSR3 CTR2 5ZJX XIL5 TVK6 E72R HSRA U5GW")
	driver.DevMode = true
	api, saveSession, _ := driver.CreateApi()
	defer saveSession()

	hasil, err := api.NotificationCounter()
	assert.NotEmpty(t, hasil)
	assert.Nil(t, err)
}
