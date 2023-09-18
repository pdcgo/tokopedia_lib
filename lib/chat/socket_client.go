package chat

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"sync"

	"github.com/pdcgo/common_conf/pdc_common"
	"github.com/pdcgo/tokopedia_lib/lib/api"
	"github.com/rs/zerolog"
	"nhooyr.io/websocket"
)

type Session interface {
	Sync() error
	UserAgent() string
	GetCookies() []*http.Cookie
}

type SocketClient struct {
	sync.Mutex
	Api     *api.TokopediaApi
	Ctx     context.Context
	Session Session
	Con     *websocket.Conn
}

func NewSocketClient(api *api.TokopediaApi) *SocketClient {

	client := &SocketClient{
		Session: api.Session,
		Api:     api,
	}
	return client
}

func (socket *SocketClient) logError(err error, handlers ...func(event *zerolog.Event) *zerolog.Event) error {
	apidata := socket.Api.AuthenticatedData
	return pdc_common.ReportErrorCustom(err, func(event *zerolog.Event) *zerolog.Event {

		nevent := event.Str("username", apidata.User.Email)
		for _, handle := range handlers {
			nevent = handle(nevent)
		}

		return nevent
	})
}

type SocketEventhandler func(socket *SocketClient, event *RcvEventSocket) error
type SocketErrorhandler func(socket *SocketClient, err error) bool

func (socket *SocketClient) Connect(ctx context.Context, eventhandler SocketEventhandler, errhandler SocketErrorhandler) {
	socket.Lock()
	defer socket.Unlock()

	uri := "wss://chat.tokopedia.com/connect"
	cookies := socket.cookiesString()
	headers := http.Header{
		"User-Agent":               {socket.Session.UserAgent()},
		"Origin":                   {"https://seller.tokopedia.com"},
		"Cookie":                   {cookies},
		"Pragma":                   {"no-cache"},
		"Cache-Control":            {"no-cache"},
		"Accept-Language":          {"en-US,en;q=0.9"},
		"Sec-WebSocket-Extensions": {"permessage-deflate; client_max_window_bits"},
	}
	opts := websocket.DialOptions{
		HTTPHeader: headers,
	}

	con, _, err := websocket.Dial(ctx, uri, &opts)
	if err != nil {
		pdc_common.ReportError(err)
	}

	socket.Con = con
	socket.Ctx = ctx
	socket.ListenData(eventhandler, errhandler)

}

func (socket *SocketClient) ListenData(eventhandler SocketEventhandler, errhandler SocketErrorhandler) {
Parent:
	for {
		select {
		case <-socket.Ctx.Done():
			break Parent

		default:
			tipe, msg, err := socket.Con.Read(socket.Ctx)
			if err != nil {
				isBreak := errhandler(socket, err)
				if isBreak {
					break Parent
				}

				pdc_common.ReportError(err)
				continue
			}

			if tipe == websocket.MessageText {
				event := RcvEventSocket{}
				err := json.Unmarshal(msg, &event)

				if err != nil {
					socket.logError(err, func(event *zerolog.Event) *zerolog.Event {
						return event.Str("eventdata", string(msg))
					})

					continue
				}

				err = eventhandler(socket, &event)
				if err != nil {
					socket.logError(err)
					continue
				}

			} else {
				log.Println(tipe, string(msg))
			}

		}
	}
}

func (socket *SocketClient) cookiesString() string {
	cookies := socket.Session.GetCookies()
	var results string
	for _, cookie := range cookies {
		results += cookie.Name + "=" + cookie.Value + ";"
	}
	return results
}

func (socket *SocketClient) EmitEvent(event *EmitEventSocket) error {
	w, err := socket.Con.Writer(socket.Ctx, websocket.MessageText)
	if err != nil {
		return socket.logError(err)
	}
	return json.NewEncoder(w).Encode(event)

}

func (sClient *SocketClient) SendEvent(payload interface{}) error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	payloadText, _ := json.Marshal(payload)
	log.Println(string(payloadText))
	err := sClient.Con.Write(ctx, websocket.MessageText, payloadText)
	if err != nil {
		pdc_common.ReportError(err)
		return err
	}
	return nil
}
