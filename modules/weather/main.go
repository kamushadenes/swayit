package weather

import (
	"fmt"
	"github.com/kamushadenes/swayit/common"
	"github.com/kamushadenes/swayit/config"
	"strings"
)

type Weather struct {
	Name        string
	Description string
	Slug        string
	output      *common.WaybarOutput
}

var module = &Weather{
	Name:        "Weather",
	Description: "Weather Notifications",
	Slug:        "weather",
	output:      &common.WaybarOutput{Class: "weather"},
}

func (m *Weather) GetName() string {
	return m.Name
}

func (m *Weather) GetDescription() string {
	return m.Description
}

func (m *Weather) GetSlug() string {
	return m.Slug
}

func (m *Weather) GetWaybarOutput() (*common.WaybarOutput, error) {
	m.output.Text = ""
	m.output.Tooltip = ""
	err := run(m.output)
	m.output.Text = strings.TrimSpace(m.output.Text)
	m.output.Tooltip = strings.TrimSpace(m.output.Tooltip)

	return m.output, err
}

func (m *Weather) SaveLastRun() {
	common.SaveLastRun(m.Slug)
}
func (m *Weather) GetLastRun() string {
	return common.GetLastRun(m.Slug)
}

func (m *Weather) WriteOutput() error {
	return common.WriteOutput(m.Slug, m.output.ToJSON())
}

func (m *Weather) Run() error {
	m.SaveLastRun()

	_, err := m.GetWaybarOutput()

	if err != nil {
		return err
	}

	return m.WriteOutput()
}

func (m *Weather) GetRunInterval() int64 {
	return config.SwayItConfig.Weather.Interval
}

func (m *Weather) GetRunIntervalOnBattery() int64 {
	return config.SwayItConfig.Weather.IntervalOnBattery
}

func (m *Weather) RunCommand(name string, args []string) error {
	switch name {
	default:
		return fmt.Errorf("unknown command %s", name)
	}
}

func (m *Weather) IsEnabled() bool {
	return config.SwayItConfig.Weather.Enabled
}

func (m *Weather) SuspendOnBattery() bool {
	return config.SwayItConfig.Weather.SuspendOnBattery
}

func GetModule() common.Module {
	return module
}
