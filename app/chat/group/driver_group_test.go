package group_test

import (
	"testing"

	socketio "github.com/googollee/go-socket.io"
	"github.com/pdcgo/tokopedia_lib/app/chat/group"
	"github.com/stretchr/testify/assert"
)

func TestDriverGroup(t *testing.T) {

	sio := socketio.NewServer(nil)
	driverData := group.NewDriverApiData(sio)
	driverGroup := group.NewDriverGroup(driverData)

	t.Run("test add driver api", func(t *testing.T) {

		dapi, err := driverGroup.AddDriverApi("pdcthoni@gmail.com", "SilentIsMyMantra", "IULIWGH6TIK3CZBKHGE27DBRLQ5LR5WQ")
		assert.Nil(t, err)
		assert.NotEmpty(t, dapi)
		assert.Equal(t, "schmart", dapi.GetUsername())

		ndapi, err := driverData.Get(7125740)
		assert.Nil(t, err)
		assert.NotEmpty(t, ndapi)
		assert.Equal(t, "schmart", ndapi.GetUsername())
	})

	t.Run("test with driver api", func(t *testing.T) {

		exist := false
		err := driverGroup.WithDriverApi(7125740, func(dapi *group.DriverApi) error {
			exist = true
			assert.NotEmpty(t, dapi)
			assert.Equal(t, "schmart", dapi.GetUsername())
			return nil
		})
		assert.Nil(t, err)
		assert.True(t, exist)
	})

	t.Run("test open driver", func(t *testing.T) {

		cancel, err := driverGroup.OpenDriver(7125740)
		assert.Nil(t, err)
		cancel()
	})
}
