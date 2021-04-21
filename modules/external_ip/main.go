package externalIp

import (
	"fmt"
	"github.com/kamushadenes/swayit/common"
	"github.com/kamushadenes/swayit/config"
	"strings"
)

type ExternalIP struct {
	Name        string
	Description string
	Slug        string
	output      *common.WaybarOutput
}

var externalIp = &ExternalIP{
	Name:        "ExternalIP",
	Description: "Get External IP",
	Slug:        "externalIp",
	output:      &common.WaybarOutput{Class: "external_ip"},
}

func (m *ExternalIP) GetName() string {
	return m.Name
}

func (m *ExternalIP) GetDescription() string {
	return m.Description
}

func (m *ExternalIP) GetSlug() string {
	return m.Slug
}

func (m *ExternalIP) GetWaybarOutput() (*common.WaybarOutput, error) {
	m.output.Text = ""
	m.output.Tooltip = ""
	err := run(m.output)
	m.output.Text = strings.TrimSpace(m.output.Text)
	m.output.Tooltip = strings.TrimSpace(m.output.Tooltip)

	return m.output, err
}

func (m *ExternalIP) SaveLastRun() {
	common.SaveLastRun(m.Slug)
}
func (m *ExternalIP) GetLastRun() string {
	return common.GetLastRun(m.Slug)
}

func (m *ExternalIP) WriteOutput() error {
	return common.WriteOutput(m.Slug, m.output.ToJSON())
}

func (m *ExternalIP) Run() error {
	m.SaveLastRun()

	_, err := m.GetWaybarOutput()

	if err != nil {
		return err
	}

	return m.WriteOutput()
}

func (m *ExternalIP) GetRunInterval() int64 {
	return config.SwayItConfig.ExternalIP.Interval
}

func (m *ExternalIP) GetRunIntervalOnBattery() int64 {
	return config.SwayItConfig.ExternalIP.IntervalOnBattery
}

func (m *ExternalIP) RunCommand(name string, args []string) error {
	switch name {
	default:
		return fmt.Errorf("unknown command %s", name)
	}

	return nil
}

func (m *ExternalIP) IsEnabled() bool {
	return config.SwayItConfig.ExternalIP.Enabled
}

func (m *ExternalIP) SuspendOnBattery() bool {
	return config.SwayItConfig.ExternalIP.SuspendOnBattery
}

func GetModule() common.Module {
	return externalIp
}
