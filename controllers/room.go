package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/manhtai/cusbot/config"
	"github.com/manhtai/cusbot/models"
	"gopkg.in/mgo.v2/bson"
)

// RoomList lists all the room available
func RoomList(w http.ResponseWriter, r *http.Request) {
	var data []models.Room
	config.Mgo.DB("cusbot").C("rooms").Find(nil).All(&data)
	config.Templ.ExecuteTemplate(w, "room-list.html", data)
}

// RoomNew is used to create new chat room
func RoomNew(w http.ResponseWriter, r *http.Request) {
	// Stub an user to be populated from the body
	room := models.Room{}

	// Populate the user data
	json.NewDecoder(r.Body).Decode(&room)

	// Add an Id
	room.Id = bson.NewObjectId()

	// Write the user to mongo
	config.Mgo.DB("cusbot").C("rooms").Insert(room)

	// Marshal provided interface into JSON structure
	rj, _ := json.Marshal(room)

	config.Templ.ExecuteTemplate(w, "room-new.html", rj)
}

// RoomDetail is where we chat, it holds history of all chat in the room
func RoomDetail(w http.ResponseWriter, r *http.Request) {
	data := map[string]interface{}{
		"Host": r.Host,
	}
	vars := mux.Vars(r)
	var room models.Room
	config.Mgo.DB("cusbot").C("rooms").FindId(vars["id"]).One(&room)
	data["room"] = room
	config.Templ.ExecuteTemplate(w, "room-detail.html", data)
}
