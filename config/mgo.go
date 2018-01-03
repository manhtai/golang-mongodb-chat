package config

import "gopkg.in/mgo.v2"

// Mgo hold our Mongodb session
var Mgo *mgo.Session

func init() {
	session, err := mgo.Dial("mongodb://localhost")
	if err != nil {
		panic(err)
	}
	Mgo = session
}
