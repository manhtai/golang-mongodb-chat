package models

import (
	"strings"
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
	room *Room
	user *User
}

func (c *Client) read() {
	defer c.socket.Close()
	for {
		var msg *Message
		err := c.socket.ReadJSON(&msg)
		if err != nil {
			return
		}
		msg.Timestamp = time.Now()

		// Default nick name to Anonymous
		if len(c.user.Name) == 0 {
			c.user.Name = "User #" + c.user.ID.Hex()
		}

		// Allow to change nick name
		nick := "/nick "
		if strings.HasPrefix(msg.Body, nick) && len(msg.Body[len(nick):]) > 0 {
			c.user.Name = msg.Body[len(nick):]
			msg.Name = c.user.Name
			msg.Body = "Your nick now is " + c.user.Name
			c.send <- msg
		} else {
			msg.Name = c.user.Name
			c.room.forward <- msg
		}
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
