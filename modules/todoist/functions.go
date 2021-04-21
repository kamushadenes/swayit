package todoist

import (
	"bytes"
	"encoding/csv"
	"fmt"
	"github.com/kamushadenes/swayit/common"
	"github.com/kamushadenes/swayit/config"
	"github.com/rgeoghegan/tabulate"
	"io"
	"os/exec"
	"strings"
	"time"
)

func formatTasksWofi(tasks []*Task) []string {
	var formatted []string
	
	for _, t := range tasks {
		formatted = append(formatted, fmt.Sprintf("<b>[%s]</b> - <b>%s</b>\n%s", t.Id, t.Project, t.Description))
	}
	
	return formatted
}

func parseLabels(labels string) []string {
	var ls []string

	for _, v := range strings.Split(labels, ",") {
		ls = append(ls, strings.TrimPrefix(v, "@"))
	}

	return ls
}

func getTasks() ([]*Task, error) {
	var tasks []*Task

	cmd := exec.Command(config.SwayItConfig.Todoist.Command, "--csv", "list")

	output, err := cmd.Output()
	if err != nil {
		return tasks, err
	}

	reader := bytes.NewReader(output)

	r := csv.NewReader(reader)

taskLoop:
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			return tasks, err
		}

		var t Task
		t.Id = record[0]
		t.Priority = record[1]
		due, err := time.Parse("06/01/02(Mon) 15:04", record[2])
		if err == nil {
			t.Due = &due
		}
		t.Project = strings.TrimPrefix(record[3], "#")
		t.Labels = parseLabels(record[4])
		t.Description = record[5]

		for _, p := range config.SwayItConfig.Todoist.ExcludedProjects {
			if strings.Contains(t.Project, p) {
				continue taskLoop
			}
		}

		tasks = append(tasks, &t)
	}
	
	return tasks, nil
}

func run(w *common.WaybarOutput) error {
	tasks, err := getTasks()
	if err != nil {
		return err
	}

	if len(tasks) > 0 {
		w.Text = fmt.Sprintf("\uF0AE %d", len(tasks))

		w.Tooltip = "<b>Tasks</b>\n"

		var rows []*Row

		for _, t := range tasks {
			rows = append(rows, t.ToRow())
		}

		table, err := tabulate.Tabulate(rows, &tabulate.Layout{Format: tabulate.FancyGridFormat})
		if err != nil {
			return err
		}

		w.Tooltip += table
	}

	return nil
}
