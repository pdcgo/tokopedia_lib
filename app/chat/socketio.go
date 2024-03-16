package chat

import (
	"log"

	socketio "github.com/googollee/go-socket.io"
	"github.com/pdcgo/common_conf/common_concept"
)

func CreateSocketIO(event *common_concept.CoreEvent) *socketio.Server {
	server := socketio.NewServer(nil)

	// on socket connect
	server.OnConnect("/", func(c socketio.Conn) error {
		c.SetContext("")
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
