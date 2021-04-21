package todoist

import (
	"fmt"
	"github.com/kamushadenes/swayit/common"
	"github.com/kamushadenes/swayit/config"
	"io/ioutil"
	"os"
	"os/exec"
	"path"
	"strings"
)

type Todoist struct {
	Name        string
	Description string
	Slug        string
	output      *common.WaybarOutput
}

var module = &Todoist{
	Name:        "Todoist",
	Description: "Todoist Tasks",
	Slug:        "todoist",
	output:      &common.WaybarOutput{Class: "todoist"},
}

func (m *Todoist) GetName() string {
	return m.Name
}

func (m *Todoist) GetDescription() string {
	return m.Description
}

func (m *Todoist) GetSlug() string {
	return m.Slug
}

func (m *Todoist) GetWaybarOutput() (*common.WaybarOutput, error) {
	m.output.Text = ""
	m.output.Tooltip = ""
	err := run(m.output)
	m.output.Text = strings.TrimSpace(m.output.Text)
	m.output.Tooltip = strings.TrimSpace(m.output.Tooltip)

	return m.output, err
}

func (m *Todoist) SaveLastRun() {
	common.SaveLastRun(m.Slug)
}
func (m *Todoist) GetLastRun() string {
	return common.GetLastRun(m.Slug)
}

func (m *Todoist) WriteOutput() error {
	return common.WriteOutput(m.Slug, m.output.ToJSON())
}

func (m *Todoist) Run() error {
	m.SaveLastRun()

	_, err := m.GetWaybarOutput()

	if err != nil {
		return err
	}

	return m.WriteOutput()
}

func (m *Todoist) GetRunInterval() int64 {
	return config.SwayItConfig.Todoist.Interval
}

func (m *Todoist) GetRunIntervalOnBattery() int64 {
	return config.SwayItConfig.Todoist.IntervalOnBattery
}

func (m *Todoist) RunCommand(name string, args []string) error {
	switch name {
	case "annotateTask":
		tasks, err := getTasks()
		if err != nil {
			return err
		}

		_, id, err := common.Wofi("Select Task to Annotate", formatTasksWofi(tasks))
		if err != nil {
			return err
		}

		fname := path.Join(config.SwayItConfig.Paths.Notes, "Tasks", fmt.Sprintf("%s.md", tasks[id].Id))

		if _, err = os.Stat(fname); os.IsNotExist(err) {
			header := fmt.Sprintf("# Task: %s", tasks[id].Description)
			if len(tasks[id].Labels) > 0 {
				header += fmt.Sprintf("# Labels: %s", strings.Join(tasks[id].Labels, ","))
			}
			err = ioutil.WriteFile(fname, []byte(header), 0644)
			if err != nil {
				return err
			}
		}

		err = common.OpenEditor(fname)
		if err != nil {
			return err
		}
	case "addTask":
		tasks, err := getTasks()
		if err != nil {
			return err
		}

		var projectsMap = make(map[string]bool)

		for _, t := range tasks {
			projectsMap[t.Project] = true
		}

		var projects []string
		for k := range projectsMap {
			projects = append(projects, k)
		}

		project, _, err := common.WofiInput("Select Project", projects)
		if err != nil {
			return err
		}

		if len(project) > 0 {
			desc, _, err := common.WofiInput("Task Description", nil, "--lines", "-1")
			if err != nil {
				return err
			}

			cmd := exec.Command(config.SwayItConfig.Todoist.Command, "add", "--project-name", project, desc)
			err = cmd.Run()
			if err != nil {
				return err
			}
		}
	case "closeTask":
		tasks, err := getTasks()
		if err != nil {
			return err
		}

		_, id, err := common.Wofi("Select Task to Complete", formatTasksWofi(tasks))
		if err != nil {
			return err
		}

		cmd := exec.Command(config.SwayItConfig.Todoist.Command, "close", tasks[id].Id)
		err = cmd.Run()
		if err != nil {
			return err
		}
	case "removeTask":
		tasks, err := getTasks()
		if err != nil {
			return err
		}

		_, id, err := common.Wofi("Select Task to Remove", formatTasksWofi(tasks))
		
		decision := common.WofiConfirm(fmt.Sprintf("Are you sure you want to remove task %s - %s?", tasks[id].Id, tasks[id].Description))
		
		if decision {
			cmd := exec.Command(config.SwayItConfig.Todoist.Command, "delete", tasks[id].Id)
			err = cmd.Run()
			if err != nil {
				return err
			}
		}
	default:
		return fmt.Errorf("unknown command %s", name)
	}

	return nil
}

func (m *Todoist) IsEnabled() bool {
	return config.SwayItConfig.Todoist.Enabled
}

func (m *Todoist) SuspendOnBattery() bool {
	return config.SwayItConfig.Todoist.SuspendOnBattery
}

func GetModule() common.Module {
	return module
}
