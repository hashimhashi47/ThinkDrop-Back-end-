package websocket

import (
	"github.com/gofiber/websocket/v2"
	"sync"
)

type Hub struct {
	clients   map[*websocket.Conn]bool
	broadcast chan []byte
	mutex     sync.Mutex
}

func NewHub() *Hub {
	return &Hub{
		clients:   make(map[*websocket.Conn]bool),
		broadcast: make(chan []byte),
	}
}

func (h *Hub) Run() {
	for {
		message := <-h.broadcast

		h.mutex.Lock()
		for client := range h.clients {
			err := client.WriteMessage(websocket.TextMessage, message)
			if err != nil {
				client.Close()
				delete(h.clients, client)
			}
		}
		h.mutex.Unlock()
	}
}

func (h *Hub) AddClient(c *websocket.Conn) {
	h.mutex.Lock()
	h.clients[c] = true
	h.mutex.Unlock()
}

func (h *Hub) RemoveClient(c *websocket.Conn) {
	h.mutex.Lock()
	delete(h.clients, c)
	h.mutex.Unlock()
}

func (h *Hub) Broadcast(message []byte) {
	h.broadcast <- message
}
