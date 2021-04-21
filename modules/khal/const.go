package khal

import (
	"fmt"
	"regexp"
)

var (
	eventDelimiter       = "@@@@@@"
	descriptionDelimiter = "######"
	eventFormat          = fmt.Sprintf(
		"{{\"uid\": \"{uid}\", \"start-date\": \"{start-date}\", \"start-time\": \"{start-time}\", \"end-date\": \"{end-date}\", \"end-time\": \"{end-time}\", \"location\": \"{location}\", \"status\": \"{status}\", \"organizer\": \"{organizer}\"}}%s{title}%s{description}%s",
		descriptionDelimiter, descriptionDelimiter, eventDelimiter)
	meetRegex = regexp.MustCompile("https://meet\\.google\\.com/(\\S+)")
	zoomRegex = regexp.MustCompile("https://(?:([a-zA-Z0-9]+)\\.)?zoom\\.us/[jsu]/(\\d+)(?:\\?pwd=([a-zA-Z0-9]+))?")
)
