package group_test

import (
	"testing"

	socketio "github.com/googollee/go-socket.io"
	"github.com/pdcgo/tokopedia_lib/app/chat/group"
	"github.com/pdcgo/tokopedia_lib/app/chat/model"
	"github.com/pdcgo/tokopedia_lib/lib/chat"
	"github.com/stretchr/testify/assert"
)

func TestSocketData(t *testing.T) {

	server := socketio.NewServer(nil)
	data := group.NewSocketData(server)

	account := model.AccountData{
		Username: "test",
		ShopID:   6699,
	}

	t.Run("test add socket", func(t *testing.T) {

		called := false
		cancel := func() {
			called = true
		}

		addclient := &chat.SocketClient{}
		data.Add(&account, addclient, cancel)

		ssocket, err := data.Get(account.ShopID)
		assert.Nil(t, err)
		assert.NotEmpty(t, ssocket)
		assert.False(t, called)

		t.Run("test cancel called", func(t *testing.T) {
			data.Add(&account, addclient, cancel)

			ssocket, err := data.Get(account.ShopID)
			assert.Nil(t, err)
			assert.NotEmpty(t, ssocket)
			assert.True(t, called)
		})
	})

	t.Run("test get socket", func(t *testing.T) {

		gdapi, err := data.Get(6699)
		assert.Nil(t, err)
		assert.NotEmpty(t, gdapi)
		assert.Equal(t, "test", gdapi.Account.Username)
	})
}
