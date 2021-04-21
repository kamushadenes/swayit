package hash

import (
	"fmt"
	"github.com/kamushadenes/swayit/common"
	"sort"
	"strings"
)

type Hash struct {
	Name        string
	Description string
	Slug        string
	output      *common.WaybarOutput
}

var module = &Hash{
	Name:        "Hash",
	Description: "Quick Hash Menu",
	Slug:        "hash",
}

func (m *Hash) GetName() string {
	return m.Name
}

func (m *Hash) GetDescription() string {
	return m.Description
}

func (m *Hash) GetSlug() string {
	return m.Slug
}

func (m *Hash) GetWaybarOutput() (*common.WaybarOutput, error) {
	return nil, fmt.Errorf("not a waybar module")
}

func (m *Hash) SaveLastRun() {
	common.SaveLastRun(m.Slug)
}
func (m *Hash) GetLastRun() string {
	return common.GetLastRun(m.Slug)
}

func (m *Hash) WriteOutput() error {
	return common.WriteOutput(m.Slug, m.output.ToJSON())
}

func (m *Hash) Run() error {
	m.SaveLastRun()

	_, err := m.GetWaybarOutput()

	if err != nil {
		return err
	}

	return m.WriteOutput()
}

func (m *Hash) GetRunInterval() int64 {
	return 0
}

func (m *Hash) GetRunIntervalOnBattery() int64 {
	return 0
}

func (m *Hash) RunCommand(name string, args []string) error {
	switch name {
	case "pick":
		text, _, err := common.WofiInput("Enter the text to be hashed", nil, "--lines", "1")
		if err != nil {
			return err
		}

		var hashes []string

		hasheds := HashAll(text)

		for k := range hasheds {
			hashes = append(hashes, fmt.Sprintf("<b>%s</b>\n%s", hasheds[k].Name, hasheds[k].Value))
		}
		sort.Slice(hashes, func(i, j int) bool {
			return hashes[i] < hashes[j]
		})

		out, _, err := common.Wofi("Select the hash algorithm to copy", hashes)
		if err != nil {
			return err
		}

		u := strings.Split(out, "\n")[1]

		return common.CopyToClipboard(u)
	default:
		return fmt.Errorf("unknown command %s", name)
	}
}

func (m *Hash) IsEnabled() bool {
	return true
}

func (m *Hash) SuspendOnBattery() bool {
	return false
}

func GetModule() common.Module {
	return module
}
