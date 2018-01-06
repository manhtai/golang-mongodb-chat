package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/manhtai/cusbot/controllers"
	"github.com/manhtai/cusbot/models"
)

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		panic("PORT must be set!")
	}
	port = ":" + port

	router := mux.NewRouter()

	r := models.NewRoom()
	go r.Run()

	// TODO: Handle login here
	router.HandleFunc("/", controllers.Index)

	router.HandleFunc("/channel", controllers.ChannelList)
	router.HandleFunc("/channel/new", controllers.ChannelNew)
	router.HandleFunc("/channel/{id}/chat", models.RoomChat(r))
	router.HandleFunc("/channel/{id}/view", controllers.ChannelView)
	router.HandleFunc("/channel/{id}/history", controllers.ChannelHistory)

	// router.GET("/user/", UserList)
	// router.GET("/user/:id", UserDetail)

	log.Println("Starting web server on", port)
	log.Fatal(http.ListenAndServe(port, router))
}
