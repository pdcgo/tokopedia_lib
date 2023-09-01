package chat

import (
	"log"

	socketio "github.com/googollee/go-socket.io"
	"github.com/pdcgo/common_conf/common_concept"
)

type TestSocketEvent struct{}

func CreateSocketIO(event *common_concept.CoreEvent) *socketio.Server {
	server := socketio.NewServer(nil)

	// handle socket events
	go func() {
		for event := range event.GetEvent() {

			log.Println(event)

			switch ev := event.(type) {

			case *TestSocketEvent:
				log.Println(ev)

			}
		}
	}()

	// on socket connect
	server.OnConnect("", func(c socketio.Conn) error {
		log.Println("[ socket ] connected", c.ID())
		return nil
	})

	server.OnDisconnect("", func(c socketio.Conn, msg string) {
		log.Println("[ socket ] disconnected", c.ID(), "-", msg)
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
