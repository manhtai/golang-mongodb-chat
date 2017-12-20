package main

import (
	"github.com/manhtai/cusbot/client"
	"html/template"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"sync"
)

type templateHandler struct {
	once     sync.Once
	filename string
	templ    *template.Template
}

func (t *templateHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	t.once.Do(func() {
		t.templ = template.Must(
			template.ParseFiles(filepath.Join("templates", t.filename)))
	})
	data := map[string]interface{}{
		"Host": r.Host,
	}
	t.templ.Execute(w, data)
}

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		port = "8080"
	}

	port = ":" + port
	r := client.NewRoom()

	http.Handle("/chat", &templateHandler{filename: "chat.html"})
	http.Handle("/room", r)
	go r.Run()

	log.Println("Starting web server on", port)

	if err := http.ListenAndServe(port, nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
