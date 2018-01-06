package config

import (
	"log"

	"gopkg.in/mgo.v2"
)

// Mgo hold our Mongodb session
var Mgo *mgo.Session

func init() {
	session, err := mgo.Dial("mongodb://localhost")
	if err != nil {
		panic(err)
	}
	Mgo = session

	// Ensure some Index
	err = session.DB("cusbot").C("messages").EnsureIndexKey("channel", "timestamp")
	if err != nil {
		log.Fatal(err)
	}
}
