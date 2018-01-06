package controllers

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/manhtai/cusbot/config"
	"github.com/manhtai/cusbot/models"
	"gopkg.in/mgo.v2/bson"
)

// ChannelList lists all the channel available to chat
func ChannelList(w http.ResponseWriter, r *http.Request) {
	var data []models.Channel
	config.Mgo.DB("cusbot").C("channels").Find(nil).All(&data)
	config.Templ.ExecuteTemplate(w, "channel-list.html", data)
}

// ChannelNew is used to create new chat channel
func ChannelNew(w http.ResponseWriter, r *http.Request) {

	data := map[string]interface{}{}

	if r.Method == http.MethodPost {
		// Stub an user to be populated from the body
		channel := models.Channel{}

		// Populate the user data
		err := r.ParseForm()
		if err != nil {
			log.Fatal(err)
		}

		channel.Name = r.FormValue("name")

		if channel.Name == "" {
			channel.Name = "No name"
		}

		channel.ID = bson.NewObjectId()

		// Write the user to mongo
		config.Mgo.DB("cusbot").C("channels").Insert(channel)

		data["channel"] = channel
	}

	config.Templ.ExecuteTemplate(w, "channel-new.html", data)
}

// ChannelView is where we chat, it displays history along with current chat in the channel
func ChannelView(w http.ResponseWriter, r *http.Request) {
	data := map[string]interface{}{
		"Host": r.Host,
	}

	var channel models.Channel
	vars := mux.Vars(r)
	id := bson.ObjectIdHex(vars["id"])

	config.Mgo.DB("cusbot").C("channels").FindId(id).One(&channel)
	data["channel"] = channel

	config.Templ.ExecuteTemplate(w, "channel-view.html", data)
}

// ChannelHistory hold chat history in a channel
func ChannelHistory(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	var channel models.Channel
	config.Mgo.DB("cusbot").C("channels").FindId(vars["id"]).One(&channel)

	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte("[]"))
}
