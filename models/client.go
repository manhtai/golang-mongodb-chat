package models

import (
	"log"
	"time"

	"github.com/gorilla/websocket"
)

// Client represents a user connect to a room, one user may have many devices to chat,
// so it should not be the same as user
type Client struct {
	// socket is the web socket for this client.
	socket *websocket.Conn
	// send is a channel on which messages are sent.
	send chan *Message
	// room is the room this client is chatting in.
	room    *Room
	user    *User
	channel string
	save    *chan SaveMessage
}

func (c *Client) read() {
	defer c.socket.Close()
	for {
		var msg *Message
		err := c.socket.ReadJSON(&msg)
		if err != nil {
			log.Fatal(err)
			return
		}
		msg.Timestamp = time.Now()

		// Default nick name to Anonymous
		if len(c.user.Name) == 0 {
			c.user.Name = c.user.ID.Hex()
		}

		msg.Name = c.user.Name
		msg.User = c.user.ID.Hex()
		msg.Channel = c.channel

		c.room.forward <- msg

		sm := &SaveMessage{
			channel: c.channel,
			message: msg,
		}

		// send message to save in another channel
		*c.save <- *sm
	}
}

func (c *Client) write() {
	defer c.socket.Close()
	for msg := range c.send {
		err := c.socket.WriteJSON(msg)
		if err != nil {
			return
		}
	}
}
