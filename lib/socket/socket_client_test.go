package socket_test

import (
	"context"
	"testing"

	"github.com/pdcgo/tokopedia_lib"
	"github.com/pdcgo/tokopedia_lib/lib/socket"
	"github.com/stretchr/testify/assert"
)

func TestNewClientWebsocket(t *testing.T) {
	driver, _ := tokopedia_lib.NewDriverAccount("bethdunn892@outlook.com", "MZT2Zk8U", "FSR3 CTR2 5ZJX XIL5 TVK6 E72R HSRA U5GW")
	sClient := socket.CreateSocketClient(driver.Session)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	err := sClient.NewClient(ctx)
	assert.Nil(t, err)
}

func TestSendEvent(t *testing.T) {
	driver, _ := tokopedia_lib.NewDriverAccount("bethdunn892@outlook.com", "MZT2Zk8U", "FSR3 CTR2 5ZJX XIL5 TVK6 E72R HSRA U5GW")
	sClient := socket.CreateSocketClient(driver.Session)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	err := sClient.NewClient(ctx)
	assert.Nil(t, err)

	payload := socket.BaseSocketType{
		Code: 301,
		Data: socket.MessageId{
			MsgId: 2484220963,
		},
	}
	err = sClient.SendEvent(payload)
	assert.Nil(t, err)
}

func TestSendMessage(t *testing.T) {
	driver, _ := tokopedia_lib.NewDriverAccount("bethdunn892@outlook.com", "MZT2Zk8U", "FSR3 CTR2 5ZJX XIL5 TVK6 E72R HSRA U5GW")
	sClient := socket.CreateSocketClient(driver.Session)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	err := sClient.NewClient(ctx)
	assert.Nil(t, err)

	startTime := socket.StartTime()
	t.Logf("%s startTime", startTime)

	localId := socket.LocalId()
	t.Logf("%s localId", localId)
	t.Logf("%s Reply Time String", socket.ReplyTimeStr())

	payload := socket.BaseSocketType{
		Code: 103,
		Data: socket.DataSend{
			MessageId:    2484220963,
			Message:      "halo dek",
			From:         "Imam",
			FromUserName: "Imam",
			Source:       "inbox",
			ParentReply:  nil,
			LocalId:      localId,
			StartTime:    startTime,
		},
	}

	err = sClient.SendEvent(payload)
	assert.Nil(t, err)
}
