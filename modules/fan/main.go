package fan

import (
	"fmt"
	"github.com/kamushadenes/swayit/common"
	"github.com/kamushadenes/swayit/config"
	"strings"
)

type Fan struct {
	Name        string
	Description string
	Slug        string
	output      *common.WaybarOutput
}

var fan = &Fan{
	Name:        "Fan",
	Description: "Laptop Fan RPM",
	Slug:        "fan",
	output:      &common.WaybarOutput{Class: "fan"},
}

func (m *Fan) GetName() string {
	return m.Name
}

func (m *Fan) GetDescription() string {
	return m.Description
}

func (m *Fan) GetSlug() string {
	return m.Slug
}

func (m *Fan) GetWaybarOutput() (*common.WaybarOutput, error) {
	m.output.Text = ""
	m.output.Tooltip = ""
	err := run(m.output)
	m.output.Text = strings.TrimSpace(m.output.Text)
	m.output.Tooltip = strings.TrimSpace(m.output.Tooltip)

	return m.output, err
}

func (m *Fan) SaveLastRun() {
	common.SaveLastRun(m.Slug)
}
func (m *Fan) GetLastRun() string {
	return common.GetLastRun(m.Slug)
}

func (m *Fan) WriteOutput() error {
	return common.WriteOutput(m.Slug, m.output.ToJSON())
}

func (m *Fan) Run() error {
	m.SaveLastRun()

	_, err := m.GetWaybarOutput()

	if err != nil {
		return err
	}

	return m.WriteOutput()
}

func (m *Fan) GetRunInterval() int64 {
	return config.SwayItConfig.Fan.Interval
}

func (m *Fan) GetRunIntervalOnBattery() int64 {
	return config.SwayItConfig.Fan.IntervalOnBattery
}

func (m *Fan) RunCommand(name string, args []string) error {
	switch name {
	default:
		return fmt.Errorf("unknown command %s", name)
	}

	return nil
}

func (m *Fan) IsEnabled() bool {
	return config.SwayItConfig.Fan.Enabled
}

func (m *Fan) SuspendOnBattery() bool {
	return config.SwayItConfig.Fan.SuspendOnBattery
}

func GetModule() common.Module {
	return fan
}
