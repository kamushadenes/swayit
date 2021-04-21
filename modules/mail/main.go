package mail

import (
	"fmt"
	"github.com/kamushadenes/swayit/common"
	"github.com/kamushadenes/swayit/config"
	"os/exec"
	"strings"
)

type Mail struct {
	Name        string
	Description string
	Slug        string
	output      *common.WaybarOutput
}

var module = &Mail{
	Name:        "Mail",
	Description: "Mail Notifications",
	Slug:        "mail",
	output:      &common.WaybarOutput{Class: "mail"},
}

func (m *Mail) GetName() string {
	return m.Name
}

func (m *Mail) GetDescription() string {
	return m.Description
}

func (m *Mail) GetSlug() string {
	return m.Slug
}

func (m *Mail) GetWaybarOutput() (*common.WaybarOutput, error) {
	m.output.Text = ""
	m.output.Tooltip = ""
	err := run(m.output)
	m.output.Text = strings.TrimSpace(m.output.Text)
	m.output.Tooltip = strings.TrimSpace(m.output.Tooltip)

	return m.output, err
}

func (m *Mail) SaveLastRun() {
	common.SaveLastRun(m.Slug)
}
func (m *Mail) GetLastRun() string {
	return common.GetLastRun(m.Slug)
}

func (m *Mail) WriteOutput() error {
	return common.WriteOutput(m.Slug, m.output.ToJSON())
}

func (m *Mail) Run() error {
	m.SaveLastRun()

	_, err := m.GetWaybarOutput()

	if err != nil {
		return err
	}

	return m.WriteOutput()
}

func (m *Mail) GetRunInterval() int64 {
	return config.SwayItConfig.Mail.Interval
}

func (m *Mail) GetRunIntervalOnBattery() int64 {
	return config.SwayItConfig.Mail.IntervalOnBattery
}

func (m *Mail) RunCommand(name string, args []string) error {
	switch name {
	case "openMail":
		common.Notify(m.GetName(), "Opening web mail")
		cmd := exec.Command("xdg-open", "https://mail.google.com")
		err := cmd.Start()
		if err != nil {
			return err
		}
	case "sync":
		common.Notify(m.GetName(), "Synchronizing mail")
		cmd := exec.Command("mbsync", "-a")
		err := cmd.Run()
		if err != nil {
			return err
		}
	default:
		return fmt.Errorf("unknown command %s", name)
	}

	return nil
}

func (m *Mail) IsEnabled() bool {
	return config.SwayItConfig.Mail.Enabled
}

func (m *Mail) SuspendOnBattery() bool {
	return config.SwayItConfig.Mail.SuspendOnBattery
}

func GetModule() common.Module {
	return module
}
