package content

import (
	"bytes"
	"fmt"
	"strings"

	"gopkg.in/yaml.v3"
)

func ParsePage(data []byte) (*Page, error) {
	meta, body, err := splitFrontmatter(data)
	if err != nil {
		return nil, err
	}

	var fm Frontmatter
	if meta != nil {
		if err := yaml.Unmarshal(meta, &fm); err != nil {
			return nil, fmt.Errorf("parsing frontmatter: %w", err)
		}
	}

	sections, err := parseSections(body)
	if err != nil {
		return nil, err
	}

	return &Page{Meta: fm, Sections: sections}, nil
}

func splitFrontmatter(data []byte) (meta []byte, body []byte, err error) {
	const delimiter = "---"

	trimmed := bytes.TrimLeft(data, " \t\r\n")
	if !bytes.HasPrefix(trimmed, []byte(delimiter)) {
		return nil, data, nil
	}

	afterFirst := trimmed[len(delimiter):]
	if len(afterFirst) > 0 && afterFirst[0] == '\n' {
		afterFirst = afterFirst[1:]
	} else if len(afterFirst) > 1 && afterFirst[0] == '\r' && afterFirst[1] == '\n' {
		afterFirst = afterFirst[2:]
	} else if len(afterFirst) == 0 {
		return nil, data, nil
	}

	idx := bytes.Index(afterFirst, []byte("\n"+delimiter))
	if idx < 0 {
		return nil, data, nil
	}

	meta = afterFirst[:idx]
	rest := afterFirst[idx+1+len(delimiter):]
	if len(rest) > 0 && rest[0] == '\n' {
		rest = rest[1:]
	} else if len(rest) > 1 && rest[0] == '\r' && rest[1] == '\n' {
		rest = rest[2:]
	}
	return meta, rest, nil
}

func parseSections(body []byte) ([]Section, error) {
	if len(bytes.TrimSpace(body)) == 0 {
		return nil, nil
	}

	var sections []Section
	current := Section{Type: SectionText}
	var buf strings.Builder

	lines := strings.Split(string(body), "\n")
	for _, line := range lines {
		trimmed := strings.TrimSpace(line)

		if trimmed == "@text" {
			current.Content = strings.TrimRight(buf.String(), "\n")
			sections = append(sections, current)
			buf.Reset()
			current = Section{Type: SectionText}
			continue
		}

		if strings.HasPrefix(trimmed, "@code") {
			parts := strings.Fields(trimmed)
			if len(parts) < 2 {
				return nil, fmt.Errorf("@code directive requires a language")
			}
			current.Content = strings.TrimRight(buf.String(), "\n")
			sections = append(sections, current)
			buf.Reset()
			current = Section{Type: SectionCode, Language: parts[1]}
			continue
		}

		buf.WriteString(line)
		buf.WriteByte('\n')
	}

	current.Content = strings.TrimRight(buf.String(), "\n")
	sections = append(sections, current)

	// Filter out empty leading section
	if len(sections) > 0 && strings.TrimSpace(sections[0].Content) == "" && sections[0].Type == SectionText {
		sections = sections[1:]
	}

	return sections, nil
}
