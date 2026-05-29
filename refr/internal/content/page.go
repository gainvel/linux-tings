package content

type Frontmatter struct {
	Title       string   `yaml:"title"`
	Order       int      `yaml:"order"`
	Description string   `yaml:"description"`
	Tags        []string `yaml:"tags"`
}

type Page struct {
	Meta     Frontmatter
	Sections []Section
}
