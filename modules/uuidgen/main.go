package uuidgen

import (
	"fmt"
	"github.com/kamushadenes/swayit/common"
	uuid "github.com/satori/go.uuid"
	"strings"
)

type UUIDGen struct {
	Name        string
	Description string
	Slug        string
	output      *common.WaybarOutput
}

var module = &UUIDGen{
	Name:        "UUIDGen",
	Description: "UUID Generation Menu",
	Slug:        "uuidgen",
}

func (m *UUIDGen) GetName() string {
	return m.Name
}

func (m *UUIDGen) GetDescription() string {
	return m.Description
}

func (m *UUIDGen) GetSlug() string {
	return m.Slug
}

func (m *UUIDGen) GetWaybarOutput() (*common.WaybarOutput, error) {
	return nil, fmt.Errorf("not a waybar module")
}

func (m *UUIDGen) SaveLastRun() {
	common.SaveLastRun(m.Slug)
}
func (m *UUIDGen) GetLastRun() string {
	return common.GetLastRun(m.Slug)
}

func (m *UUIDGen) WriteOutput() error {
	return common.WriteOutput(m.Slug, m.output.ToJSON())
}

func (m *UUIDGen) Run() error {
	m.SaveLastRun()

	_, err := m.GetWaybarOutput()

	if err != nil {
		return err
	}

	return m.WriteOutput()
}

func (m *UUIDGen) GetRunInterval() int64 {
	return 0
}

func (m *UUIDGen) GetRunIntervalOnBattery() int64 {
	return 0
}

func (m *UUIDGen) RunCommand(name string, args []string) error {
	switch name {
	case "pick":
		var uuidList []string

		u1 := uuid.NewV1()
		uuidList = append(uuidList, fmt.Sprintf("<b>UUID v1</b>\n%s", u1.String()))

		u4 := uuid.NewV4()
		uuidList = append(uuidList, fmt.Sprintf("<b>UUID v4</b>\n%s", u4.String()))

		out, _, err := common.Wofi("Select the UUID", uuidList)
		if err != nil {
			return err
		}
		
		u := strings.Split(out, "\n")[1]

		return common.CopyToClipboard(u)
	default:
		return fmt.Errorf("unknown command %s", name)
	}
}

func (m *UUIDGen) IsEnabled() bool {
	return true
}

func (m *UUIDGen) SuspendOnBattery() bool {
	return false
}

func GetModule() common.Module {
	return module
}
