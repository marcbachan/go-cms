package model

type BlogPost struct {
	Title      string `json:"title"`
	Excerpt    string `json:"excerpt"`
	CoverImage string `json:"coverImage"`
	Date       string `json:"date"`
	OGImage    struct {
		URL string `json:"url"`
	} `json:"ogImage"`
	Tags    []string `json:"tags"`
	Content string   `json:"content"`
	Slug    string   `json:"slug"`
}
