package main

import (
	"log"
	"net/http"
	"os"

	"cms/config"
	"cms/handlers"

	"github.com/gorilla/mux"
)

func main() {
	config.LoadConfig("config.json")

	os.MkdirAll(config.AppConfig.PostsDir, os.ModePerm)
	os.MkdirAll(config.AppConfig.ImagesDir, os.ModePerm)

	r := mux.NewRouter()

	r.HandleFunc("/api/posts", handlers.CreatePost).Methods("POST")
	r.HandleFunc("/api/upload", handlers.UploadImage).Methods("POST")
	r.HandleFunc("/new", handlers.NewPostForm).Methods("GET")

	log.Println("Server running at :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
