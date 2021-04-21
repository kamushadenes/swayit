package itau

import (
	"fmt"
	"github.com/kamushadenes/swayit/common"
	"github.com/kamushadenes/swayit/config"
	"strings"
)

type Itau struct {
	Name        string
	Description string
	Slug        string
	output      *common.WaybarOutput
}

var module = &Itau{
	Name:        "Itau",
	Description: "Itau Statements",
	Slug:        "itau",
	output:      &common.WaybarOutput{Class: "itau"},
}

func (m *Itau) GetName() string {
	return m.Name
}

func (m *Itau) GetDescription() string {
	return m.Description
}

func (m *Itau) GetSlug() string {
	return m.Slug
}

func (m *Itau) GetWaybarOutput() (*common.WaybarOutput, error) {
	m.output.Text = ""
	m.output.Tooltip = ""
	err := run(m.output)
	m.output.Text = strings.TrimSpace(m.output.Text)
	m.output.Tooltip = strings.TrimSpace(m.output.Tooltip)

	return m.output, err
}

func (m *Itau) SaveLastRun() {
	common.SaveLastRun(m.Slug)
}
func (m *Itau) GetLastRun() string {
	return common.GetLastRun(m.Slug)
}

func (m *Itau) WriteOutput() error {
	return common.WriteOutput(m.Slug, m.output.ToJSON())
}

func (m *Itau) Run() error {
	m.SaveLastRun()

	_, err := m.GetWaybarOutput()

	if err != nil {
		return err
	}

	return m.WriteOutput()
}

func (m *Itau) GetRunInterval() int64 {
	return config.SwayItConfig.Itau.Interval
}

func (m *Itau) GetRunIntervalOnBattery() int64 {
	return config.SwayItConfig.Itau.IntervalOnBattery
}

func (m *Itau) RunCommand(name string, args []string) error {
	switch name {
	default:
		return fmt.Errorf("unknown command %s", name)
	}

	return nil
}

func (m *Itau) IsEnabled() bool {
	return config.SwayItConfig.Itau.Enabled
}

func (m *Itau) SuspendOnBattery() bool {
	return config.SwayItConfig.Itau.SuspendOnBattery
}

func GetModule() common.Module {
	return module
}
