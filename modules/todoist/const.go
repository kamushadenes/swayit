package todoist

import "regexp"

var (
	idRegex = regexp.MustCompile("\\[(\\d+)\\]")
)