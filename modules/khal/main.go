package khal

import (
	"fmt"
	"github.com/kamushadenes/swayit/common"
	"github.com/kamushadenes/swayit/config"
	"os/exec"
	"strings"
)

type Khal struct {
	Name        string
	Description string
	Slug        string
	output      *common.WaybarOutput
}

var module = &Khal{
	Name:        "Khal",
	Description: "Khal Event Integration",
	Slug:        "khal",
	output:      &common.WaybarOutput{Class: "khal"},
}

func (m *Khal) GetName() string {
	return m.Name
}

func (m *Khal) GetDescription() string {
	return m.Description
}

func (m *Khal) GetSlug() string {
	return m.Slug
}

func (m *Khal) GetWaybarOutput() (*common.WaybarOutput, error) {
	m.output.Text = ""
	m.output.Tooltip = ""
	err := run(m.output)
	m.output.Text = strings.TrimSpace(m.output.Text)
	m.output.Tooltip = strings.TrimSpace(m.output.Tooltip)

	return m.output, err
}

func (m *Khal) SaveLastRun() {
	common.SaveLastRun(m.Slug)
}
func (m *Khal) GetLastRun() string {
	return common.GetLastRun(m.Slug)
}

func (m *Khal) WriteOutput() error {
	return common.WriteOutput(m.Slug, m.output.ToJSON())
}

func (m *Khal) Run() error {
	m.SaveLastRun()

	_, err := m.GetWaybarOutput()

	if err != nil {
		return err
	}

	return m.WriteOutput()
}

func (m *Khal) GetRunInterval() int64 {
	return config.SwayItConfig.Khal.Interval
}

func (m *Khal) GetRunIntervalOnBattery() int64 {
	return config.SwayItConfig.Khal.IntervalOnBattery
}

func (m *Khal) RunCommand(name string, args []string) error {
	switch name {
	case "getMeetingMarkdown":
		fmt.Println(common.GetData(m.GetSlug(), "currentMeetingMarkdown"))
		return nil
	case "getMeeting":
		fmt.Println(common.GetData(m.GetSlug(), "currentMeeting"))
		return nil
	case "getMeetingLink":
		fmt.Println(common.GetData(m.GetSlug(), "currentMeetingLink"))
		return nil
	case "openMeetingLink":
		common.Notify(m.GetName(), "Opening meeting")
		cmd := exec.Command("xdg-open", common.GetData(m.GetSlug(), "currentMeetingLink"))
		err := cmd.Start()
		if err != nil {
			return err
		}
	default:
		return fmt.Errorf("unknown command %s", name)
	}

	return nil
}

func (m *Khal) IsEnabled() bool {
	return config.SwayItConfig.Khal.Enabled
}

func (m *Khal) SuspendOnBattery() bool {
	return config.SwayItConfig.Khal.SuspendOnBattery
}

func GetModule() common.Module {
	return module
}
