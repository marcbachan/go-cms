package config

import (
	"encoding/json"
	"log"
	"os"
)

type Settings struct {
	PostsDir  string `json:"postsDir"`
	ImagesDir string `json:"imagesDir"`
}

var AppConfig Settings

func LoadConfig(path string) {
	file, err := os.Open(path)
	if err != nil {
		log.Fatalf("Failed to open config file: %v", err)
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&AppConfig); err != nil {
		log.Fatalf("Failed to decode config JSON: %v", err)
	}
}
