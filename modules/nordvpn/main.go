package nordvpn

import (
	"fmt"
	"github.com/kamushadenes/swayit/common"
	"github.com/kamushadenes/swayit/config"
	"os/exec"
	"strings"
)

type NordVPN struct {
	Name        string
	Description string
	Slug        string
	output      *common.WaybarOutput
}

var module = &NordVPN{
	Name:        "NordVPN",
	Description: "NordVPN Status",
	Slug:        "nordvpn",
	output:      &common.WaybarOutput{Class: "nordvpn"},
}

func (m *NordVPN) GetName() string {
	return m.Name
}

func (m *NordVPN) GetDescription() string {
	return m.Description
}

func (m *NordVPN) GetSlug() string {
	return m.Slug
}

func (m *NordVPN) GetWaybarOutput() (*common.WaybarOutput, error) {
	m.output.Text = ""
	m.output.Tooltip = ""
	err := run(m.output)
	m.output.Text = strings.TrimSpace(m.output.Text)
	m.output.Tooltip = strings.TrimSpace(m.output.Tooltip)

	return m.output, err
}

func (m *NordVPN) SaveLastRun() {
	common.SaveLastRun(m.Slug)
}
func (m *NordVPN) GetLastRun() string {
	return common.GetLastRun(m.Slug)
}

func (m *NordVPN) WriteOutput() error {
	return common.WriteOutput(m.Slug, m.output.ToJSON())
}

func (m *NordVPN) Run() error {
	m.SaveLastRun()

	_, err := m.GetWaybarOutput()

	if err != nil {
		return err
	}

	return m.WriteOutput()
}

func (m *NordVPN) GetRunInterval() int64 {
	return config.SwayItConfig.NordVPN.Interval
}

func (m *NordVPN) GetRunIntervalOnBattery() int64 {
	return config.SwayItConfig.NordVPN.IntervalOnBattery
}

func (m *NordVPN) RunCommand(name string, args []string) error {
	switch name {
	case "connect":
		common.Notify(m.GetName(), "Connecting")
		cmd := exec.Command(config.SwayItConfig.NordVPN.Command, "connect")
		err := cmd.Run()
		if err != nil {
			return err
		}
	case "disconnect":
		common.Notify(m.GetName(), "Disconnecting")
		cmd := exec.Command(config.SwayItConfig.NordVPN.Command, "disconnect")
		err := cmd.Run()
		if err != nil {
			return err
		}
	default:
		return fmt.Errorf("unknown command %s", name)
	}

	return nil
}

func (m *NordVPN) IsEnabled() bool {
	return config.SwayItConfig.NordVPN.Enabled
}

func (m *NordVPN) SuspendOnBattery() bool {
	return config.SwayItConfig.NordVPN.SuspendOnBattery
}

func GetModule() common.Module {
	return module
}
