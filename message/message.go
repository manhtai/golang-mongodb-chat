package message

import "time"

// Message represents a single message which a client sent to a room
type Message struct {
	Name string
	Body string
	When time.Time
}
