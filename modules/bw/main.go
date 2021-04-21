package bw

import (
	"fmt"
	"github.com/kamushadenes/swayit/common"
	"strings"
)

type BW struct {
	Name        string
	Description string
	Slug        string
	output      *common.WaybarOutput
}

var module = &BW{
	Name:        "BitWarden",
	Description: "BitWarden Credential Picker",
	Slug:        "bw",
}

func (m *BW) GetName() string {
	return m.Name
}

func (m *BW) GetDescription() string {
	return m.Description
}

func (m *BW) GetSlug() string {
	return m.Slug
}

func (m *BW) GetWaybarOutput() (*common.WaybarOutput, error) {
	return nil, fmt.Errorf("not a waybar module")
}

func (m *BW) SaveLastRun() {
	common.SaveLastRun(m.Slug)
}
func (m *BW) GetLastRun() string {
	return common.GetLastRun(m.Slug)
}

func (m *BW) WriteOutput() error {
	return common.WriteOutput(m.Slug, m.output.ToJSON())
}

func (m *BW) Run() error {
	m.SaveLastRun()

	_, err := m.GetWaybarOutput()

	if err != nil {
		return err
	}

	return m.WriteOutput()
}

func (m *BW) GetRunInterval() int64 {
	return 0
}

func (m *BW) GetRunIntervalOnBattery() int64 {
	return 0
}

func (m *BW) RunCommand(name string, args []string) error {
	switch name {
	case "pick":
		items, err := GetItems()
		if err != nil {
			return err
		}
		
		_, id, err := common.Wofi("Select an Item", formatItemsWofi(items))
		if err != nil {
			return err
		}
		
		item := items[id]

		options := item.ToWofi()

		_, id, err = common.Wofi("Select the Data to Copy", options)
		if err != nil {
			return err
		}
		
		if id == 0 {
			return common.CopyToClipboard(item.Login.Password.(string))
		} else {
			text := strings.TrimSpace(strings.Split(options[id], "\n")[1])

			return common.CopyToClipboard(text)
		}
	default:
		return fmt.Errorf("unknown command %s", name)
	}
}

func (m *BW) IsEnabled() bool {
	return true
}

func (m *BW) SuspendOnBattery() bool {
	return false
}

func GetModule() common.Module {
	return module
}
