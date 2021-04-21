package intel_gpu

import (
	"fmt"
	"github.com/kamushadenes/swayit/common"
	"github.com/kamushadenes/swayit/config"
	"strings"
)

type IntelGPU struct {
	Name        string
	Description string
	Slug        string
	output      *common.WaybarOutput
}

var module = &IntelGPU{
	Name:        "IntelGPU",
	Description: "Intel GPU Top",
	Slug:        "intelGpu",
	output:      &common.WaybarOutput{Class: "intelGpu"},
}

func (m *IntelGPU) GetName() string {
	return m.Name
}

func (m *IntelGPU) GetDescription() string {
	return m.Description
}

func (m *IntelGPU) GetSlug() string {
	return m.Slug
}

func (m *IntelGPU) GetWaybarOutput() (*common.WaybarOutput, error) {
	m.output.Text = ""
	m.output.Tooltip = ""
	err := run(m.output)
	m.output.Text = strings.TrimSpace(m.output.Text)
	m.output.Tooltip = strings.TrimSpace(m.output.Tooltip)

	return m.output, err
}

func (m *IntelGPU) SaveLastRun() {
	common.SaveLastRun(m.Slug)
}
func (m *IntelGPU) GetLastRun() string {
	return common.GetLastRun(m.Slug)
}

func (m *IntelGPU) WriteOutput() error {
	return common.WriteOutput(m.Slug, m.output.ToJSON())
}

func (m *IntelGPU) Run() error {
	m.SaveLastRun()

	_, err := m.GetWaybarOutput()

	if err != nil {
		return err
	}

	return m.WriteOutput()
}

func (m *IntelGPU) GetRunInterval() int64 {
	return config.SwayItConfig.IntelGPU.Interval
}

func (m *IntelGPU) GetRunIntervalOnBattery() int64 {
	return config.SwayItConfig.IntelGPU.IntervalOnBattery
}

func (m *IntelGPU) RunCommand(name string, args []string) error {
	switch name {
	default:
		return fmt.Errorf("unknown command %s", name)
	}

	return nil
}

func (m *IntelGPU) IsEnabled() bool {
	return config.SwayItConfig.IntelGPU.Enabled
}

func (m *IntelGPU) SuspendOnBattery() bool {
	return config.SwayItConfig.IntelGPU.SuspendOnBattery
}

func GetModule() common.Module {
	return module
}
