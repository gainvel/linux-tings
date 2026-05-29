package content

type SectionType int

const (
	SectionText SectionType = iota
	SectionCode
)

type Section struct {
	Type     SectionType
	Language string
	Content  string
}
