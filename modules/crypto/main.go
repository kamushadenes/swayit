package crypto

import (
	"fmt"
	"github.com/kamushadenes/swayit/common"
	"github.com/kamushadenes/swayit/config"
	"strings"
)

type Crypto struct {
	Name        string
	Description string
	Slug        string
	output      *common.WaybarOutput
}

var module = &Crypto{
	Name:        "Crypto",
	Description: "CryptoCoin Ticker",
	Slug:        "crypto",
	output:      &common.WaybarOutput{Class: "crypto"},
}

func (m *Crypto) GetName() string {
	return m.Name
}

func (m *Crypto) GetDescription() string {
	return m.Description
}

func (m *Crypto) GetSlug() string {
	return m.Slug
}

func (m *Crypto) GetWaybarOutput() (*common.WaybarOutput, error) {
	m.output.Text = ""
	m.output.Tooltip = ""
	m.output.Alt = ""
	err := run(m.output)
	m.output.Text = strings.TrimSpace(m.output.Text)
	m.output.Tooltip = strings.TrimSpace(m.output.Tooltip)
	m.output.Alt = strings.TrimSpace(m.output.Alt)

	return m.output, err
}

func (m *Crypto) SaveLastRun() {
	common.SaveLastRun(m.Slug)
}
func (m *Crypto) GetLastRun() string {
	return common.GetLastRun(m.Slug)
}

func (m *Crypto) WriteOutput() error {
	return common.WriteOutput(m.Slug, m.output.ToJSON())
}

func (m *Crypto) Run() error {
	m.SaveLastRun()

	_, err := m.GetWaybarOutput()

	if err != nil {
		return err
	}

	return m.WriteOutput()
}

func (m *Crypto) GetRunInterval() int64 {
	return config.SwayItConfig.Crypto.Interval
}

func (m *Crypto) GetRunIntervalOnBattery() int64 {
	return config.SwayItConfig.Crypto.IntervalOnBattery
}

func (m *Crypto) RunCommand(name string, args []string) error {
	switch name {
	default:
		return fmt.Errorf("unknown command %s", name)
	}

	return nil
}

func (m *Crypto) IsEnabled() bool {
	return config.SwayItConfig.Crypto.Enabled
}

func (m *Crypto) SuspendOnBattery() bool {
	return config.SwayItConfig.Crypto.SuspendOnBattery
}

func GetModule() common.Module {
	return module
}
