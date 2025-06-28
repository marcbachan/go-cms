package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"

	"cms/config"
	"cms/model"
	"cms/storage"
	"cms/utils"

	"github.com/gorilla/mux"
)

func CreatePost(w http.ResponseWriter, r *http.Request) {
	var post model.BlogPost
	if err := json.NewDecoder(r.Body).Decode(&post); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	filename := fmt.Sprintf("%s.md", post.Slug)
	fullPath := filepath.Join(config.AppConfig.PostsDir, filename)

	if err := storage.WriteMarkdownWithFrontmatter(fullPath, post); err != nil {
		http.Error(w, "Failed to write post", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "Post written to %s\n", fullPath)
}

func UploadImage(w http.ResponseWriter, r *http.Request) {
	title := r.FormValue("title")
	if title == "" {
		http.Error(w, "Missing title field", http.StatusBadRequest)
		return
	}

	file, header, err := r.FormFile("image")
	if err != nil {
		http.Error(w, "Image not provided", http.StatusBadRequest)
		return
	}
	defer file.Close()

	slug := utils.Slugify(title)
	imageDir := filepath.Join(config.AppConfig.ImagesDir, slug)
	os.MkdirAll(imageDir, os.ModePerm)

	dstPath := filepath.Join(imageDir, header.Filename)
	dst, err := os.Create(dstPath)
	if err != nil {
		http.Error(w, "Failed to save image", http.StatusInternalServerError)
		return
	}
	defer dst.Close()
	io.Copy(dst, file)

	webPath := fmt.Sprintf("/assets/img/%s/%s", slug, header.Filename)
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"url": webPath,
	})
}

func UpdatePost(w http.ResponseWriter, r *http.Request) {
	slug := mux.Vars(r)["slug"]
	if slug == "" {
		http.Error(w, "Missing slug", http.StatusBadRequest)
		return
	}

	var post model.BlogPost
	if err := json.NewDecoder(r.Body).Decode(&post); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	path := filepath.Join(config.AppConfig.PostsDir, slug+".md")
	if err := storage.WriteMarkdownWithFrontmatter(path, post); err != nil {
		http.Error(w, "Failed to update post", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Post updated")
}

func DeletePost(w http.ResponseWriter, r *http.Request) {
	slug := mux.Vars(r)["slug"]
	path := filepath.Join(config.AppConfig.PostsDir, slug+".md")

	if err := os.Remove(path); err != nil {
		http.Error(w, "Could not delete post", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Deleted"))
}
