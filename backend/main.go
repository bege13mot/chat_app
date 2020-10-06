package main

import (
	"net/http"

	log "github.com/sirupsen/logrus"

	"github.com/bege13mot/chat_app/pkg/websocket"
)

// WebSocket endpoint
func serveWs(pool *websocket.Pool, w http.ResponseWriter, r *http.Request) {
	log.Debug("WebSocket Endpoint Hit")

	// upgrade this connection to WebSocket
	conn, err := websocket.Upgrade(w, r)

	if err != nil {
		log.Error(w, "%+v\n", err)
	}

	client := &websocket.Client{
		Conn: conn,
		Pool: pool,
	}

	pool.Register <- client
	client.Read()
}

func setupRoutes() {
	pool := websocket.NewPool()
	go pool.Start()

	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		serveWs(pool, w, r)
	})
}

func main() {
	log.Warning("Distributed Chat App v0.01")
	setupRoutes()
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Error("Can't start server: ", err)
	}
}
