package content

type Node struct {
	Name        string
	Path        string
	Description string
	Order       int
	IsCategory  bool
	Children    []*Node
	Page        *Page
}
