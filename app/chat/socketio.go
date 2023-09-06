package chat

import (
	"log"

	socketio "github.com/googollee/go-socket.io"
	"github.com/pdcgo/common_conf/common_concept"
	"github.com/pdcgo/tokopedia_lib/app/chat/sio_event"
)

func CreateSocketIO(event *common_concept.CoreEvent) *socketio.Server {
	server := socketio.NewServer(nil)

	// handle socket events
	go func() {
		for event := range event.GetEvent() {
			switch ev := event.(type) {

			case *sio_event.FrontendNotificationEvent:
				go server.BroadcastToNamespace("", "notification", ev)
			}
		}
	}()

	// on socket connect
	server.OnConnect("", func(c socketio.Conn) error {
		log.Println("[ socket ] connected", c.ID())
		return nil
	})

	// run socketio
	go func() {
		log.Println("running socketio...")
		if err := server.Serve(); err != nil {
			log.Panicln(err)
		}
	}()

	return server
}
