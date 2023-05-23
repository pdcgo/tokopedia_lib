package socket

import (
	"context"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/pdcgo/common_conf/pdc_common"
	"nhooyr.io/websocket"
)

type Session interface {
	Sync() error
	UserAgent() string
	GetCookies() []*http.Cookie
}

type SocketClient struct {
	url     string
	Session Session
	c       *websocket.Conn
}

func (sClient *SocketClient) cookiesString() string {
	cookies := sClient.Session.GetCookies()
	var results string
	for _, cookie := range cookies {
		results += cookie.Name + "=" + cookie.Value + ";"
	}
	return results
}

func (sClient *SocketClient) WsOption() *websocket.DialOptions {
	cookies := sClient.cookiesString()
	headers := http.Header{
		"User-Agent":               {sClient.Session.UserAgent()},
		"Origin":                   {"https://seller.tokopedia.com"},
		"Cookie":                   {cookies},
		"Pragma":                   {"no-cache"},
		"Cache-Control":            {"no-cache"},
		"Accept-Language":          {"en-US,en;q=0.9"},
		"Sec-WebSocket-Extensions": {"permessage-deflate; client_max_window_bits"},
	}
	opts := websocket.DialOptions{
		// HTTPClient: &http.Client{
		// 	Transport: &http.Transport{
		// 		Proxy: http.ProxyURL(&url.URL{
		// 			Scheme: "http",
		// 			Host:   "localhost:8888",
		// 		}),
		// 	},
		// },
		HTTPHeader: headers,
	}
	return &opts

}

func (sClient *SocketClient) NewClient(ctx context.Context, url string, opts *websocket.DialOptions) (*SocketClient, error) {
	c, _, err := websocket.Dial(ctx, url, opts)
	if err != nil {
		return nil, err
	}

	cl := &SocketClient{
		url: url,
		c:   c,
	}

	return cl, nil
}

func (sClient *SocketClient) ConnectWebsocket() {
	wsOpts := sClient.WsOption()
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	client, err := sClient.NewClient(ctx, sClient.url, wsOpts)
	if err != nil {
		pdc_common.ReportError(err)
	}
	defer client.c.Close(websocket.StatusNormalClosure, "connction was closed")

	go func() {
		for {
			msgType, msg, err := client.c.Reader(ctx)
			if err != nil {
				pdc_common.ReportError(err)
				return
			}
			data, _ := io.ReadAll(msg)
			log.Println(string(data))
			log.Println(msgType)
		}
	}()

	ticker := time.NewTicker(time.Second * 24)
	for {
		select {
		case <-ticker.C:
			err := client.c.Ping(ctx)
			if err != nil {
				pdc_common.ReportError(err)
			} else {
				log.Println("ping success")
			}
		}
	}
}

func (sClient *SocketClient) SendEvent(payload interface{}) error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	payloadText, _ := json.Marshal(payload)
	log.Println(string(payloadText))
	err := sClient.c.Write(ctx, websocket.MessageText, payloadText)
	if err != nil {
		pdc_common.ReportError(err)
		return err
	}
	return nil
}

func CreateSocketClient(session Session) *SocketClient {
	wsUrl := "wss://chat.tokopedia.com/connect"
	client := &SocketClient{
		Session: session,
		url:     wsUrl,
	}
	return client
}
