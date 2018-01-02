package main

import (
	"log"
	"net/http"
	"os"

	"github.com/manhtai/cusbot/client"
	"github.com/manhtai/cusbot/config"
)

func chatHandle(w http.ResponseWriter, r *http.Request) {
	data := map[string]interface{}{
		"Host": r.Host,
	}
	config.Templ.ExecuteTemplate(w, "chat.html", data)
}

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		port = "8080"
	}

	port = ":" + port
	r := client.NewRoom()

	http.HandleFunc("/chat", chatHandle)
	http.Handle("/room", r)
	go r.Run()

	log.Println("Starting web server on", port)

	if err := http.ListenAndServe(port, nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
