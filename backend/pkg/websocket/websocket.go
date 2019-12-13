package websocket

import (
	log "github.com/sirupsen/logrus"
	"net/http"

	"github.com/gorilla/websocket"
)

// We will need to define an Upgrader
// it requres a Read and Write buffer size
var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func Upgrade(w http.ResponseWriter, r *http.Request) (*websocket.Conn, error) {
	upgrader.CheckOrigin = func(r *http.Request) bool { return true }

	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Error(err)
		return ws, err
	}
	return ws, nil
}
