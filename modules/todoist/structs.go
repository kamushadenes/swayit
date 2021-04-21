package todoist

import (
	"strings"
	"time"
)

type Task struct {
	Id          string
	Priority    string
	Due         *time.Time
	Project     string
	Labels      []string
	Description string
}

type Row struct {
	Id          string
	Priority    string
	Due         string
	Project     string
	Labels      string
	Description string
}

func (t *Task) ToRow() *Row {
	var r Row

	r.Id = t.Id
	r.Priority = t.Priority

	if t.Due != nil {
		r.Due = t.Due.Format("2006-01-02 15:04:05")
	} else {
		r.Due = "-"
	}

	r.Project = t.Project
	r.Labels = strings.Join(t.Labels, ",")
	r.Description = t.Description

	return &r
}
