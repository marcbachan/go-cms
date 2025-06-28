package handlers

import (
	"cms/config"
	"cms/utils"
	"html/template"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

func NewPostForm(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/newpost.html"))
	tmpl.Execute(w, nil)
}

func EditPostForm(w http.ResponseWriter, r *http.Request) {
	slug := strings.TrimPrefix(r.URL.Path, "/edit/")
	if slug == "" {
		http.Error(w, "Missing slug", http.StatusBadRequest)
		return
	}

	path := filepath.Join(config.AppConfig.PostsDir, slug+".md")
	content, err := os.ReadFile(path)
	if err != nil {
		http.Error(w, "Post not found", http.StatusNotFound)
		return
	}

	post, body := utils.ParseFrontmatter(content)

	tmpl := template.New("editpost.html").Funcs(template.FuncMap{
		"join": strings.Join,
	})
	tmpl = template.Must(tmpl.ParseFiles("templates/editpost.html"))

	tmpl.Execute(w, map[string]interface{}{
		"Post": post,
		"Slug": slug,
		"Body": body,
	})
}

func ListPosts(w http.ResponseWriter, r *http.Request) {
	files, err := os.ReadDir(config.AppConfig.PostsDir)
	if err != nil {
		http.Error(w, "Failed to list posts", http.StatusInternalServerError)
		return
	}

	var slugs []string
	for _, f := range files {
		if !f.IsDir() && strings.HasSuffix(f.Name(), ".md") {
			slug := strings.TrimSuffix(f.Name(), ".md")
			slugs = append(slugs, slug)
		}
	}

	tmpl := template.Must(template.ParseFiles("templates/listposts.html"))
	tmpl.Execute(w, slugs)
}
