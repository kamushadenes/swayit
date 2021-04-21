package power

import (
	"fmt"
	"github.com/kamushadenes/swayit/common"
	"github.com/kamushadenes/swayit/config"
	"strings"
)

type Power struct {
	Name        string
	Description string
	Slug        string
	output      *common.WaybarOutput
}

var module = &Power{
	Name:        "Power",
	Description: "Power Monitor",
	Slug:        "power",
	output:      &common.WaybarOutput{Class: "power"},
}

func (m *Power) GetName() string {
	return m.Name
}

func (m *Power) GetDescription() string {
	return m.Description
}

func (m *Power) GetSlug() string {
	return m.Slug
}

func (m *Power) GetWaybarOutput() (*common.WaybarOutput, error) {
	m.output.Text = ""
	m.output.Tooltip = ""
	err := run(m.output)
	m.output.Text = strings.TrimSpace(m.output.Text)
	m.output.Tooltip = strings.TrimSpace(m.output.Tooltip)

	return m.output, err
}

func (m *Power) SaveLastRun() {
	common.SaveLastRun(m.Slug)
}
func (m *Power) GetLastRun() string {
	return common.GetLastRun(m.Slug)
}

func (m *Power) WriteOutput() error {
	return common.WriteOutput(m.Slug, m.output.ToJSON())
}

func (m *Power) Run() error {
	m.SaveLastRun()

	_, err := m.GetWaybarOutput()

	if err != nil {
		return err
	}

	return m.WriteOutput()
}

func (m *Power) GetRunInterval() int64 {
	return config.SwayItConfig.Power.Interval
}

func (m *Power) GetRunIntervalOnBattery() int64 {
	return config.SwayItConfig.Power.IntervalOnBattery
}

func (m *Power) RunCommand(name string, args []string) error {
	switch name {
	default:
		return fmt.Errorf("unknown command %s", name)
	}

	return nil
}

func (m *Power) IsEnabled() bool {
	return config.SwayItConfig.Power.Enabled
}

func (m *Power) SuspendOnBattery() bool {
	return config.SwayItConfig.Power.SuspendOnBattery
}

func GetModule() common.Module {
	return module
}
