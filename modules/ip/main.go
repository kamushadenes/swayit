package ip

import (
	"fmt"
	"github.com/kamushadenes/swayit/common"
	"sort"
	"strings"
)

type IP struct {
	Name        string
	Description string
	Slug        string
	output      *common.WaybarOutput
}

var module = &IP{
	Name:        "IP",
	Description: "IP Information Menu",
	Slug:        "ip",
}

func (m *IP) GetName() string {
	return m.Name
}

func (m *IP) GetDescription() string {
	return m.Description
}

func (m *IP) GetSlug() string {
	return m.Slug
}

func (m *IP) GetWaybarOutput() (*common.WaybarOutput, error) {
	return nil, fmt.Errorf("not a waybar module")
}

func (m *IP) SaveLastRun() {
	common.SaveLastRun(m.Slug)
}
func (m *IP) GetLastRun() string {
	return common.GetLastRun(m.Slug)
}

func (m *IP) WriteOutput() error {
	return common.WriteOutput(m.Slug, m.output.ToJSON())
}

func (m *IP) Run() error {
	m.SaveLastRun()

	_, err := m.GetWaybarOutput()

	if err != nil {
		return err
	}

	return m.WriteOutput()
}

func (m *IP) GetRunInterval() int64 {
	return 0
}

func (m *IP) GetRunIntervalOnBattery() int64 {
	return 0
}

func (m *IP) RunCommand(name string, args []string) error {
	switch name {
	case "pick":
		text, _, err := common.WofiInput("Enter the IP to search", nil, "--lines", "1")
		if err != nil {
			return err
		}

		var infos []string

		info, err := GetIPInfo(text)
		if err != nil {
			return err
		}
		
		var infoMap = map[string]string{
			"Country": info.Country,
			"Country Code": info.CountryCode,
			"Region": info.Region,
			"Region Name": info.RegionName,
			"City": info.City,
			"ZIP": info.Zip,
			"Latitude": fmt.Sprintf("%f", info.Lat),
			"Longitude": fmt.Sprintf("%f", info.Lon),
			"Timezone": info.Timezone,
			"ISP": info.Isp,
			"Organization": info.Org,
			"AS": info.As,
			"Continent": info.Continent,
			"Continent Code": info.ContinentCode,
			"Currency": info.Currency,
			"District": info.District,
			"AS Name": info.AsName,
			"Reverse": info.Reverse,
			"Is Mobile?": fmt.Sprintf("%T", info.Mobile),
			"Is Proxy?": fmt.Sprintf("%T", info.Proxy),
			"Is Hosting?": fmt.Sprintf("%T", info.Hosting),
		}
		
		for k, v := range infoMap {
			infos = append(infos, fmt.Sprintf("<b>%s</b>\n%s", k, v))
		}
		sort.Slice(infos, func(i, j int) bool {
			return infos[i] < infos[j]
		})

		out, _, err := common.Wofi("Select the information to copy", infos)
		if err != nil {
			return err
		}

		u := strings.Split(out, "\n")[1]

		return common.CopyToClipboard(u)
	default:
		return fmt.Errorf("unknown command %s", name)
	}
}

func (m *IP) IsEnabled() bool {
	return true
}

func (m *IP) SuspendOnBattery() bool {
	return false
}

func GetModule() common.Module {
	return module
}
