package ssh

import (
	"fmt"
	"github.com/kamushadenes/swayit/common"
	"github.com/kamushadenes/swayit/config"
	"os/exec"
	"strings"
)

type SSH struct {
	Name        string
	Description string
	Slug        string
	output      *common.WaybarOutput
}

var module = &SSH{
	Name:        "SSH",
	Description: "SSH Menu",
	Slug:        "ssh",
}

func (m *SSH) GetName() string {
	return m.Name
}

func (m *SSH) GetDescription() string {
	return m.Description
}

func (m *SSH) GetSlug() string {
	return m.Slug
}

func (m *SSH) GetWaybarOutput() (*common.WaybarOutput, error) {
	return nil, fmt.Errorf("not a waybar module")
}

func (m *SSH) SaveLastRun() {
	common.SaveLastRun(m.Slug)
}
func (m *SSH) GetLastRun() string {
	return common.GetLastRun(m.Slug)
}

func (m *SSH) WriteOutput() error {
	return common.WriteOutput(m.Slug, m.output.ToJSON())
}

func (m *SSH) Run() error {
	m.SaveLastRun()

	_, err := m.GetWaybarOutput()

	if err != nil {
		return err
	}

	return m.WriteOutput()
}

func (m *SSH) GetRunInterval() int64 {
	return 0
}

func (m *SSH) GetRunIntervalOnBattery() int64 {
	return 0
}

func (m *SSH) RunCommand(name string, args []string) error {
	switch name {
	case "update":
		return UpdateHosts()
	case "getHosts":
		hosts, err := GetHosts()
		if err != nil {
			return err
		}
		
		fmt.Println(strings.Join(hosts, "\n"))
		
		return nil
	case "pick":
		hosts, err := GetHosts()
		if err != nil {
			return err
		}
		host, _, err := common.Wofi("SSH Hosts", hosts)
		if err != nil {
			return err
		}
		
		var sshArgs []string
		sshArgs = append(sshArgs, config.SwayItConfig.SSH.TerminalArgs...)
		sshArgs = append(sshArgs, config.SwayItConfig.SSH.Command)
		sshArgs = append(sshArgs, host)
		
		cmd := exec.Command(config.SwayItConfig.SSH.Terminal, sshArgs...)
		return cmd.Run()
	default:
		return fmt.Errorf("unknown command %s", name)
	}
}

func (m *SSH) IsEnabled() bool {
	return true
}

func (m *SSH) SuspendOnBattery() bool {
	return false
}

func GetModule() common.Module {
	return module
}
