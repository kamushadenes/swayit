package nordvpn

import (
	"regexp"
)

var (
	statusRegex = regexp.MustCompile("Status: (\\S+)")
	cityRegex   = regexp.MustCompile("City: (.*)")
)
