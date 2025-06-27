package utils

import (
	"regexp"
	"strings"
)

func Slugify(title string) string {
	slug := strings.ToLower(title)
	slug = strings.ReplaceAll(slug, " ", "_")
	slug = regexp.MustCompile(`[^\w_]+`).ReplaceAllString(slug, "")
	return slug
}
