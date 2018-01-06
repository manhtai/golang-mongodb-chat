package controllers

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/manhtai/cusbot/models"
)

const (
	socketBufferSize  = 1024
	messageBufferSize = 256
)

var upgrader = &websocket.Upgrader{ReadBufferSize: socketBufferSize,
	WriteBufferSize: socketBufferSize}

// RoomChat take a room, return a HandlerFunc,
// responsible for send & receive websocket data for all channels
func RoomChat(r *models.Room) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {

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
}
