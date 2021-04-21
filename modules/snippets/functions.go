package snippets

import "fmt"

func formatSnippetsWofi(snippets []*Snippet) []string {
	var formatted []string

	for _, s := range snippets {
		formatted = append(formatted, fmt.Sprintf("[%s] <b>%s</b>\n%s", s.Category, s.Name, s.Value))
	}

	return formatted
}
