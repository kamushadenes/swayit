package timezone

import (
	"fmt"
	"github.com/kamushadenes/swayit/common"
	"strings"
	"time"
)

type TimeZone struct {
	Name        string
	Description string
	Slug        string
	output      *common.WaybarOutput
}

var module = &TimeZone{
	Name:        "Timezone",
	Description: "Timezone Information",
	Slug:        "timezone",
}

func (m *TimeZone) GetName() string {
	return m.Name
}

func (m *TimeZone) GetDescription() string {
	return m.Description
}

func (m *TimeZone) GetSlug() string {
	return m.Slug
}

func (m *TimeZone) GetWaybarOutput() (*common.WaybarOutput, error) {
	return nil, fmt.Errorf("not a waybar module")
}

func (m *TimeZone) SaveLastRun() {
	common.SaveLastRun(m.Slug)
}
func (m *TimeZone) GetLastRun() string {
	return common.GetLastRun(m.Slug)
}

func (m *TimeZone) WriteOutput() error {
	return common.WriteOutput(m.Slug, m.output.ToJSON())
}

func (m *TimeZone) Run() error {
	m.SaveLastRun()

	_, err := m.GetWaybarOutput()

	if err != nil {
		return err
	}

	return m.WriteOutput()
}

func (m *TimeZone) GetRunInterval() int64 {
	return 0
}

func (m *TimeZone) GetRunIntervalOnBattery() int64 {
	return 0
}

func (m *TimeZone) RunCommand(name string, args []string) error {
	switch name {
	case "list":
		for _, tz := range ListTimeZones() {
			fmt.Println(tz)
		}
	case "pick":
		timeZones := ListTimeZones()
		
		var items []string
		
		now := time.Now()
		for _, tz := range timeZones {
			loc, err := time.LoadLocation(tz)
			if err != nil {
				common.Logger.Error().Err(err).Msg("an error has occurred")
			}
			t := now.In(loc)
			
			items = append(items, fmt.Sprintf("<b>%s</b>\n%s", tz, t.Format("2006-01-02 15:04:05")))
		}
		
		loc, _, err := common.WofiInput("Timezone", items)
		if err != nil {
			return err
		}

		tstr := strings.Split(loc, "\n")[1]
		return common.CopyToClipboard(tstr)
	default:
		return fmt.Errorf("unknown command %s", name)
	}

	return nil
}

func (m *TimeZone) IsEnabled() bool {
	return true
}

func (m *TimeZone) SuspendOnBattery() bool {
	return false
}

func GetModule() common.Module {
	return module
}
