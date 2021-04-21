package prometheus

import (
	"fmt"
	"github.com/kamushadenes/swayit/common"
	"github.com/kamushadenes/swayit/config"
	"os/exec"
	"strings"
)

type Prometheus struct {
	Name        string
	Description string
	Slug        string
	output      *common.WaybarOutput
}

var module = &Prometheus{
	Name:        "Prometheus",
	Description: "Prometheus Alarms",
	Slug:        "prometheus",
	output:      &common.WaybarOutput{Class: "prometheus"},
}

func (m *Prometheus) GetName() string {
	return m.Name
}

func (m *Prometheus) GetDescription() string {
	return m.Description
}

func (m *Prometheus) GetSlug() string {
	return m.Slug
}

func (m *Prometheus) GetWaybarOutput() (*common.WaybarOutput, error) {
	m.output.Text = ""
	m.output.Tooltip = ""
	err := run(m.output)
	m.output.Text = strings.TrimSpace(m.output.Text)
	m.output.Tooltip = strings.TrimSpace(m.output.Tooltip)

	return m.output, err
}

func (m *Prometheus) SaveLastRun() {
	common.SaveLastRun(m.Slug)
}
func (m *Prometheus) GetLastRun() string {
	return common.GetLastRun(m.Slug)
}

func (m *Prometheus) WriteOutput() error {
	return common.WriteOutput(m.Slug, m.output.ToJSON())
}

func (m *Prometheus) Run() error {
	m.SaveLastRun()

	_, err := m.GetWaybarOutput()

	if err != nil {
		return err
	}

	return m.WriteOutput()
}

func (m *Prometheus) GetRunInterval() int64 {
	return config.SwayItConfig.Prometheus.Interval
}

func (m *Prometheus) GetRunIntervalOnBattery() int64 {
	return config.SwayItConfig.Prometheus.IntervalOnBattery
}

func (m *Prometheus) RunCommand(name string, args []string) error {
	switch name {
	case "openGrafana":
		common.Notify(m.GetName(), "Opening Grafana")
		cmd := exec.Command("xdg-open", config.SwayItConfig.Prometheus.GrafanaURL)
		err := cmd.Start()
		if err != nil {
			return err
		}
	default:
		return fmt.Errorf("unknown command %s", name)
	}

	return nil
}

func (m *Prometheus) IsEnabled() bool {
	return config.SwayItConfig.Prometheus.Enabled
}

func (m *Prometheus) SuspendOnBattery() bool {
	return config.SwayItConfig.Prometheus.SuspendOnBattery
}

func GetModule() common.Module {
	return module
}
