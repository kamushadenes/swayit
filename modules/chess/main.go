package chess

import (
	"fmt"
	"github.com/kamushadenes/swayit/common"
	"github.com/kamushadenes/swayit/config"
	"os/exec"
	"strings"
)

type Chess struct {
	Name        string
	Description string
	Slug        string
	output      *common.WaybarOutput
}

var module = &Chess{
	Name:        "Chess",
	Description: "Chess.com Notifications",
	Slug:        "chess",
	output:      &common.WaybarOutput{Class: "chess"},
}

func (m *Chess) GetName() string {
	return m.Name
}

func (m *Chess) GetDescription() string {
	return m.Description
}

func (m *Chess) GetSlug() string {
	return m.Slug
}

func (m *Chess) GetWaybarOutput() (*common.WaybarOutput, error) {
	m.output.Text = ""
	m.output.Tooltip = ""
	err := run(m.output)
	m.output.Text = strings.TrimSpace(m.output.Text)
	m.output.Tooltip = strings.TrimSpace(m.output.Tooltip)

	return m.output, err
}

func (m *Chess) SaveLastRun() {
	common.SaveLastRun(m.Slug)
}
func (m *Chess) GetLastRun() string {
	return common.GetLastRun(m.Slug)
}

func (m *Chess) WriteOutput() error {
	return common.WriteOutput(m.Slug, m.output.ToJSON())
}

func (m *Chess) Run() error {
	m.SaveLastRun()

	_, err := m.GetWaybarOutput()

	if err != nil {
		return err
	}

	return m.WriteOutput()
}

func (m *Chess) GetRunInterval() int64 {
	return config.SwayItConfig.Chess.Interval
}

func (m *Chess) GetRunIntervalOnBattery() int64 {
	return config.SwayItConfig.Chess.IntervalOnBattery
}

func (m *Chess) RunCommand(name string, args []string) error {
	switch name {
	case "openReadyGame":
		common.Notify(m.GetName(), "Opening game")
		cmd := exec.Command("xdg-open", "https://www.chess.com/goto_ready_game")
		err := cmd.Start()
		if err != nil {
			return err
		}
	default:
		return fmt.Errorf("unknown command %s", name)
	}

	return nil
}

func (m *Chess) IsEnabled() bool {
	return config.SwayItConfig.Chess.Enabled
}

func (m *Chess) SuspendOnBattery() bool {
	return config.SwayItConfig.Chess.SuspendOnBattery
}

func GetModule() common.Module {
	return module
}
