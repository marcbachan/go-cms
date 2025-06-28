package utils

import (
	"strings"

	"cms/model"

	"gopkg.in/yaml.v3"
)

func ParseFrontmatter(content []byte) (model.BlogPost, string) {
	s := string(content)
	parts := strings.SplitN(s, "---", 3)
	if len(parts) < 3 {
		return model.BlogPost{}, s
	}

	var post model.BlogPost
	yaml.Unmarshal([]byte(parts[1]), &post)

	body := strings.TrimSpace(parts[2])
	return post, body
}
