package websocket

import (
	log "github.com/sirupsen/logrus"
)

type Pool struct {
	Register   chan *Client
	Unregister chan *Client
	Clients    map[*Client]bool
	Broadcast  chan Message
}

func NewPool() *Pool {
	return &Pool{
		Register:   make(chan *Client),
		Unregister: make(chan *Client),
		Clients:    make(map[*Client]bool),
		Broadcast:  make(chan Message),
	}
}

func (pool *Pool) Start() {
	for {
		select {
		case client := <-pool.Register:
			pool.Clients[client] = true
			log.Debug("Size of Connection Pool", len(pool.Clients))
			for client := range pool.Clients {
				log.Debug(client)
				err := client.Conn.WriteJSON(Message{Type: 1, Body: "New User Joined..."})
				if err != nil {
					log.Error("Can't write to socket: ", err)
				}
			}
			break
		case client := <-pool.Unregister:
			delete(pool.Clients, client)
			log.Debug("Size of Connection Pool", len(pool.Clients))
			for client := range pool.Clients {
				err := client.Conn.WriteJSON(Message{Type: 1, Body: "User Disconnected..."})
				if err != nil {
					log.Error("Can't write to socket: ", err)
				}
			}
			break
		case message := <-pool.Broadcast:
			log.Debug("Sending messages to all clients in Pool")
			for client:= range pool.Clients {
				if err := client.Conn.WriteJSON(message); err != nil {
					log.Error(err)
					return
				}
			}
		}
	}
}
