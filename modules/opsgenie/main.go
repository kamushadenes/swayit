package opsgenie

import (
	"fmt"
	"github.com/kamushadenes/swayit/common"
	"github.com/kamushadenes/swayit/config"
	"strings"
)

type OpsGenie struct {
	Name        string
	Description string
	Slug        string
	output      *common.WaybarOutput
}

var module = &OpsGenie{
	Name:        "OpsGenie",
	Description: "OpsGenie On-Call Notifications",
	Slug:        "opsgenie",
	output:      &common.WaybarOutput{Class: "opsgenie"},
}

func (m *OpsGenie) GetName() string {
	return m.Name
}

func (m *OpsGenie) GetDescription() string {
	return m.Description
}

func (m *OpsGenie) GetSlug() string {
	return m.Slug
}

func (m *OpsGenie) GetWaybarOutput() (*common.WaybarOutput, error) {
	m.output.Text = ""
	m.output.Tooltip = ""
	err := run(m.output)
	m.output.Text = strings.TrimSpace(m.output.Text)
	m.output.Tooltip = strings.TrimSpace(m.output.Tooltip)

	return m.output, err
}

func (m *OpsGenie) SaveLastRun() {
	common.SaveLastRun(m.Slug)
}
func (m *OpsGenie) GetLastRun() string {
	return common.GetLastRun(m.Slug)
}

func (m *OpsGenie) WriteOutput() error {
	return common.WriteOutput(m.Slug, m.output.ToJSON())
}

func (m *OpsGenie) Run() error {
	m.SaveLastRun()

	_, err := m.GetWaybarOutput()

	if err != nil {
		return err
	}

	return m.WriteOutput()
}

func (m *OpsGenie) GetRunInterval() int64 {
	return config.SwayItConfig.OpsGenie.Interval
}

func (m *OpsGenie) GetRunIntervalOnBattery() int64 {
	return config.SwayItConfig.OpsGenie.IntervalOnBattery
}

func (m *OpsGenie) RunCommand(name string, args []string) error {
	switch name {
	default:
		return fmt.Errorf("unknown command %s", name)
	}
}

func (m *OpsGenie) IsEnabled() bool {
	return config.SwayItConfig.OpsGenie.Enabled
}

func (m *OpsGenie) SuspendOnBattery() bool {
	return config.SwayItConfig.OpsGenie.SuspendOnBattery
}

func GetModule() common.Module {
	return module
}
