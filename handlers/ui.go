package handlers

import (
	"html/template"
	"net/http"
)

func NewPostForm(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/newpost.html"))
	tmpl.Execute(w, nil)
}
