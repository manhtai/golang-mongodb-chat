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
	router.HandleFunc("/", controllers.RoomList)

	r := models.NewRoom()
	router.Handle("/chat", r)
	go r.Run()

	router.HandleFunc("/room/", controllers.RoomList)
	router.HandleFunc("/room/new", controllers.RoomNew)
	router.HandleFunc("/room/chat/{id}", controllers.RoomDetail)

	// router.GET("/user/", UserList)
	// router.GET("/user/:id", UserDetail)

	log.Println("Starting web server on", port)
	log.Fatal(http.ListenAndServe(port, router))
}
