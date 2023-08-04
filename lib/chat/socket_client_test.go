package chat_test

import (
	"context"
	"testing"
	"time"

	"github.com/pdcgo/tokopedia_lib"
	"github.com/pdcgo/tokopedia_lib/lib/chat"
	"github.com/stretchr/testify/assert"
)

func TestRunningSocket(t *testing.T) {
	driver, err := tokopedia_lib.NewDriverAccount("pdcthoni@gmail.com", "SilentIsMyMantra", "IULIWGH6TIK3CZBKHGE27DBRLQ5LR5WQ")
	assert.Nil(t, err)

	api, saveSession, err := driver.CreateApi()
	assert.Nil(t, err)

	defer saveSession()

	cha := chat.NewSocketClient(api)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	cha.Connect(ctx, func(socket *chat.SocketClient, event *chat.RcvEventSocket) error {
		return nil
	})

}

// import (
// 	"context"
// 	"testing"

// 	"github.com/pdcgo/tokopedia_lib/lib/socket"
// 	"github.com/pdcgo/tokopedia_lib/scenario"
// 	"github.com/stretchr/testify/assert"
// )

// func TestNewClientWebsocket(t *testing.T) {
// 	apiSession, _ := scenario.GetTokopediaApiClient()
// 	sClient := socket.CreateSocketClient(apiSession.Session)
// 	ctx, cancel := context.WithCancel(context.Background())
// 	defer cancel()

// 	err := sClient.NewClient(ctx)
// 	assert.Nil(t, err)
// }

// func TestSendEvent(t *testing.T) {
// 	apiSession, _ := scenario.GetTokopediaApiClient()
// 	sClient := socket.CreateSocketClient(apiSession.Session)
// 	ctx, cancel := context.WithCancel(context.Background())
// 	defer cancel()

// 	err := sClient.NewClient(ctx)
// 	assert.Nil(t, err)

// 	payload := socket.BaseSocketType{
// 		Code: 301,
// 		Data: socket.MessageId{
// 			MsgId: 2484220963,
// 		},
// 	}
// 	err = sClient.SendEvent(payload)
// 	assert.Nil(t, err)
// }

// func TestSendMessage(t *testing.T) {
// 	apiSession, _ := scenario.GetTokopediaApiClient()
// 	sClient := socket.CreateSocketClient(apiSession.Session)
// 	ctx, cancel := context.WithCancel(context.Background())
// 	defer cancel()

// 	err := sClient.NewClient(ctx)
// 	assert.Nil(t, err)

// 	startTime := socket.StartTime()
// 	t.Logf("%s startTime", startTime)

// 	localId := socket.LocalId()
// 	t.Logf("%s localId", localId)
// 	t.Logf("%s Reply Time String", socket.ReplyTimeStr())

// 	payload := socket.BaseSocketType{
// 		Code: 103,
// 		Data: socket.DataSend{
// 			MessageId:    2484220963,
// 			Message:      "halo dek",
// 			From:         "Imam",
// 			FromUserName: "Imam",
// 			Source:       "inbox",
// 			ParentReply:  nil,
// 			LocalId:      localId,
// 			StartTime:    startTime,
// 		},
// 	}

// 	err = sClient.SendEvent(payload)
// 	assert.Nil(t, err)
// }
