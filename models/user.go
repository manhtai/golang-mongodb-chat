package models

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

// User hold information about an user
type User struct {
	ID        bson.ObjectId `json:"id" bson:"_id"`
	Name      string        `json:"name" bson:"name"`
	CreatedAt time.Time     `json:"created_at" bson:"created_at"`
	Active    bool          `json:"active" bson:"created_at"`
}
