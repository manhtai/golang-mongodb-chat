package config

import (
	"log"
	"os"

	"gopkg.in/mgo.v2"
)

// Mgo hold our Mongodb session
var Mgo *mgo.Session

func init() {
	mongoURI := os.Getenv("MONGODB_URI")
	if mongoURI == "" {
		mongoURI = "mongodb://localhost"
	}

	session, err := mgo.Dial(mongoURI)
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
