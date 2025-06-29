package main

import (
	"github.com/joho/godotenv"

	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"

	"cms/config"
	"cms/handlers"
)

func main() {
	_ = godotenv.Load() // load .env file automatically
	config.LoadConfig("config.json")

	// Set up secure cookie store
	sessionSecret := os.Getenv("SESSION_SECRET")
	if sessionSecret == "" {
		log.Fatal("SESSION_SECRET not set")
	}

	store := sessions.NewCookieStore([]byte(sessionSecret))
	handlers.SetStore(store)

	r := mux.NewRouter()

	// Public routes
	r.HandleFunc("/login", handlers.LoginForm).Methods("GET")
	r.HandleFunc("/", handlers.LoginForm).Methods("GET")
	r.HandleFunc("/login", handlers.Login).Methods("POST")
	r.HandleFunc("/logout", handlers.Logout).Methods("GET")

	// Static file server
	r.PathPrefix("/styles/").Handler(http.StripPrefix("/styles/", http.FileServer(http.Dir("./public/styles/"))))

	// Auth-protected routes
	protected := r.NewRoute().Subrouter()
	protected.Use(handlers.RequireLogin)

	// UI
	protected.HandleFunc("/new", handlers.NewPostForm).Methods("GET")
	protected.HandleFunc("/edit/{slug}", handlers.EditPostForm).Methods("GET")
	protected.HandleFunc("/posts", handlers.ListPosts).Methods("GET")

	// API
	protected.HandleFunc("/api/posts", handlers.CreatePost).Methods("POST")
	protected.HandleFunc("/api/posts/{slug}", handlers.UpdatePost).Methods("PUT")
	protected.HandleFunc("/api/posts/{slug}", handlers.DeletePost).Methods("DELETE")
	protected.HandleFunc("/api/upload", handlers.UploadImage).Methods("POST")

	log.Println("CMS running on http://localhost:8080")
	http.ListenAndServe(":8080", r)
}
