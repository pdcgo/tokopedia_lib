package group_test

import (
	"testing"

	socketio "github.com/googollee/go-socket.io"
	"github.com/pdcgo/tokopedia_lib"
	"github.com/pdcgo/tokopedia_lib/app/chat/group"
	"github.com/stretchr/testify/assert"
)

func TestDriverData(t *testing.T) {

	server := socketio.NewServer(nil)
	data := group.NewDriverApiData(server)

	driver, err := tokopedia_lib.NewDriverAccount("pdcthoni@gmail.com", "SilentIsMyMantra", "IULIWGH6TIK3CZBKHGE27DBRLQ5LR5WQ")
	assert.Nil(t, err)

	tapi, saveSession, err := driver.CreateApi()
	defer saveSession()
	assert.Nil(t, err)

	t.Run("test add driver", func(t *testing.T) {

		shopid := tapi.AuthenticatedData.UserShopInfo.Info.ShopID
		dapi := data.Add(driver, tapi)

		gdapi, err := data.Get(int(shopid))
		assert.Nil(t, err)
		assert.NotEmpty(t, gdapi)
		assert.Equal(t, dapi.GetUsername(), gdapi.GetUsername())
	})

	t.Run("test get driver", func(t *testing.T) {

		gdapi, err := data.Get(7125740)
		assert.Nil(t, err)
		assert.NotEmpty(t, gdapi)
		assert.Equal(t, "schmart", gdapi.GetUsername())
	})
}
