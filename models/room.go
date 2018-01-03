package models

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
	"gopkg.in/mgo.v2/bson"
)

// Room represents a room to chat
type Room struct {
	Id bson.ObjectId `json:"id" bson:"_id"`

	// forward is a channel that holds incoming messages
	// that should be forwarded to the other clients.
	forward chan *Message
	// join is a channel for clients wishing to join the room.
	join chan *Client
	// leave is a channel for clients wishing to leave the room.
	leave chan *Client
	// clients holds all current clients in this room.
	clients map[*Client]bool
}

// Run start a room and run it forever
func (r *Room) Run() {
	for {
		select {
		case client := <-r.join:
			// joining
			r.clients[client] = true
		case client := <-r.leave:
			// leaving
			delete(r.clients, client)
			close(client.send)
		case msg := <-r.forward:
			// forward message to all clients
			for client := range r.clients {
				client.send <- msg
			}
		}
	}
}

// NewRoom creates a new room for clients to join
func NewRoom() *Room {
	return &Room{
		forward: make(chan *Message),
		join:    make(chan *Client),
		leave:   make(chan *Client),
		clients: make(map[*Client]bool),
	}
}

const (
	socketBufferSize  = 1024
	messageBufferSize = 256
)

var upgrader = &websocket.Upgrader{ReadBufferSize: socketBufferSize,
	WriteBufferSize: socketBufferSize}

func (r *Room) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	socket, err := upgrader.Upgrade(w, req, nil)
	if err != nil {
		log.Fatal("ServeHTTP:", err)
		return
	}

	user := &User{
		Id: len(r.clients) + 1,
	}

	client := &Client{
		socket: socket,
		send:   make(chan *Message, messageBufferSize),
		room:   r,
		user:   user,
	}

	r.join <- client
	defer func() { r.leave <- client }()
	go client.write()
	client.read()
}
