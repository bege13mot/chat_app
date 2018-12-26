package main

import (
	"/backend/pkg/websocket"
	"fmt"
	"net/http"
)

// WebSocket endpoint
func serveWs(w http.ResponseWriter, r *http.Request) {

	// upgrade this connection to WebSocket
	ws, err := websocket.Upgrade(w, r)

	if err != nil {
		fmt.Fprintf(w, "%+V\n", err)
	}

	//listen indefinitely for new messages coming
	go websocket.Writer(ws)
	websocket.Reader(ws)
}

func setupRoutes() {
	http.HandleFunc("/ws", serveWs)
}

func main() {
	fmt.Println("Distributed Chat App v0.01")
	setupRoutes()
	http.ListenAndServe(":8080", nil)
}
