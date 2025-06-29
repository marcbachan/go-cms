package utils

import (
	"log"
	"strings"

	"cms/model"

	"gopkg.in/yaml.v3"
)

func ParseFrontmatter(content []byte) (model.BlogPost, string) {
	s := string(content)
	parts := strings.SplitN(s, "---", 3)
	if len(parts) < 3 {
		log.Println("Frontmatter split failed")
		return model.BlogPost{}, s
	}

	var post model.BlogPost
	if err := yaml.Unmarshal([]byte(parts[1]), &post); err != nil {
		log.Printf("Failed to parse frontmatter YAML: %v", err)
		return model.BlogPost{}, s
	}

	log.Printf("Parsed frontmatter: %+v", post)

	body := strings.TrimSpace(parts[2])
	return post, body
}
