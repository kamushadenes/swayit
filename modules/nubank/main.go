package nubank

import (
	"github.com/kamushadenes/swayit/common"
	"github.com/kamushadenes/swayit/config"
)

type Nubank struct {
	Name        string
	Description string
	Slug        string
	output      *common.WaybarOutput
}

var module = &Nubank{
	Name:        "Nubank",
	Description: "Nubank Statements",
	Slug:        "nubank",
	output:      &common.WaybarOutput{Class: "nubank"},
}

func (m *Nubank) GetName() string {
	return m.Name
}

func (m *Nubank) GetDescription() string {
	return m.Description
}

func (m *Nubank) GetSlug() string {
	return m.Slug
}

func (m *Nubank) GetWaybarOutput() (*common.WaybarOutput, error) {
	err := run(m.output)

	return m.output, err
}

func (m *Nubank) SaveLastRun() {
	common.SaveLastRun(m.Slug)
}
func (m *Nubank) GetLastRun() string {
	return common.GetLastRun(m.Slug)
}

func (m *Nubank) WriteOutput() error {
	return common.WriteOutput(m.Slug, m.output.ToJSON())
}

func (m *Nubank) Run() error {
	m.SaveLastRun()

	_, err := m.GetWaybarOutput()

	if err != nil {
		return err
	}

	return m.WriteOutput()
}

func (m *Nubank) GetRunInterval() int64 {
	return config.SwayItConfig.Nubank.Interval
}

func (m *Nubank) GetRunIntervalOnBattery() int64 {
	return config.SwayItConfig.Nubank.IntervalOnBattery
}

func (m *Nubank) RunCommand(name string, args []string) error {
	return nil
}

func (m *Nubank) IsEnabled() bool {
	return config.SwayItConfig.Nubank.Enabled
}

func (m *Nubank) SuspendOnBattery() bool {
	return config.SwayItConfig.Nubank.SuspendOnBattery
}

func GetModule() common.Module {
	return module
}
