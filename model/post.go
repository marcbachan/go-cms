package model

type BlogPost struct {
	Title      string `yaml:"title" json:"title"`
	Excerpt    string `yaml:"excerpt" json:"excerpt"`
	CoverImage string `yaml:"coverImage" json:"coverImage"`
	Date       string `yaml:"date" json:"date"`

	OGImage struct {
		URL string `yaml:"url" json:"url"`
	} `yaml:"ogImage" json:"ogImage"`

	Tags    []string `yaml:"tags" json:"tags"`
	Content string   `yaml:"-" json:"content"`
	Slug    string   `yaml:"-" json:"slug"`
}
