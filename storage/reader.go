package storage

import (
	"cms/model"
	"fmt"
	"os"
	"strings"

	"gopkg.in/yaml.v3"
)

func ReadMarkdownWithFrontmatter(path string) (model.BlogPost, string, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return model.BlogPost{}, "", fmt.Errorf("failed to read file: %w", err)
	}

	content := string(data)

	if !strings.HasPrefix(content, "---") {
		return model.BlogPost{}, "", fmt.Errorf("missing frontmatter in %s", path)
	}

	parts := strings.SplitN(content, "---", 3)
	if len(parts) < 3 {
		return model.BlogPost{}, "", fmt.Errorf("invalid frontmatter format in %s", path)
	}

	frontmatterYaml := strings.TrimSpace(parts[1])
	body := strings.TrimSpace(parts[2])

	var post model.BlogPost
	if err := yaml.Unmarshal([]byte(frontmatterYaml), &post); err != nil {
		return model.BlogPost{}, "", fmt.Errorf("failed to parse frontmatter: %w", err)
	}

	return post, body, nil
}
