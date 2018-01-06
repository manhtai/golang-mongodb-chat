package models

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

// Message represents a single message which a client sent to a room
// (same meaning as a user send to a channel)
type Message struct {
	Name      string        `json:"name" bson:"name"`
	Body      string        `json:"body" bson:"body"`
	Channel   bson.ObjectId `json:"channel" bson:"channel"`
	User      bson.ObjectId `json:"user" bson:"user"`
	Timestamp time.Time     `json:"timestamp,omitempty" bson:"timestamp"`
}
