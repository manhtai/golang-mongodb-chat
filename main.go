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

	// Start 2 channels to chat & save chat
	r := models.NewRoomChan()
	sm := models.NewSaveMessageChan()

	router := mux.NewRouter()

	router.HandleFunc("/", controllers.Index)

	// Auth handlers
	router.HandleFunc("/auth/login", controllers.Login)
	router.HandleFunc("/auth/{action:(?:login|callback)}/{provider}",
		controllers.LoginHandle)
	router.HandleFunc("/auth/logout", controllers.Logout)

	// Chat handlers
	router.HandleFunc("/channel", controllers.ChannelList)
	router.HandleFunc("/channel/new",
		controllers.MustAuth(controllers.ChannelNew))
	router.HandleFunc("/channel/{id}/chat", models.RoomChat(r, sm))
	router.HandleFunc("/channel/{id}/view",
		controllers.MustAuth(controllers.ChannelView))
	router.HandleFunc("/channel/{id}/history",
		controllers.MustAuth(controllers.ChannelHistory))

	// User handlers
	// router.GET("/user/", UserList)
	// router.GET("/user/:id", UserDetail)

	// The rest, just not found
	router.HandleFunc("/*", http.NotFound)

	log.Println("Starting web server on", port)
	log.Fatal(http.ListenAndServe(port, router))
}
