package models

import "time"

// User hold information about an user
type User struct {
	ID        string    `json:"id" bson:"_id"`
	Name      string    `json:"name" bson:"name"`
	CreatedAt time.Time `json:"created_at" bson:"created_at"`
	Active    bool      `json:"active" bson:"created_at"`
}
