package snippets

import (
	"fmt"
	"github.com/kamushadenes/swayit/common"
	"github.com/kamushadenes/swayit/config"
	"io/ioutil"
	"path"
	"strings"
)

type Snippets struct {
	Name        string
	Description string
	Slug        string
	output      *common.WaybarOutput
}

var module = &Snippets{
	Name:        "Snippets",
	Description: "Snippets Menu",
	Slug:        "snippets",
}

func (m *Snippets) GetName() string {
	return m.Name
}

func (m *Snippets) GetDescription() string {
	return m.Description
}

func (m *Snippets) GetSlug() string {
	return m.Slug
}

func (m *Snippets) GetWaybarOutput() (*common.WaybarOutput, error) {
	return nil, fmt.Errorf("not a waybar module")
}

func (m *Snippets) SaveLastRun() {
	common.SaveLastRun(m.Slug)
}
func (m *Snippets) GetLastRun() string {
	return common.GetLastRun(m.Slug)
}

func (m *Snippets) WriteOutput() error {
	return common.WriteOutput(m.Slug, m.output.ToJSON())
}

func (m *Snippets) Run() error {
	m.SaveLastRun()

	_, err := m.GetWaybarOutput()

	if err != nil {
		return err
	}

	return m.WriteOutput()
}

func (m *Snippets) GetRunInterval() int64 {
	return 0
}

func (m *Snippets) GetRunIntervalOnBattery() int64 {
	return 0
}

func (m *Snippets) RunCommand(name string, args []string) error {
	switch name {
	case "pick":
		files, err := common.ListFiles(config.SwayItConfig.Paths.Snippets)
		if err != nil {
			return err
		}

		var snippets = make([]*Snippet, 0)

		for _, f := range files {
			b, err := ioutil.ReadFile(path.Join(config.SwayItConfig.Paths.Snippets, f))
			if err != nil {
				continue
			}

			data := string(b)
			lines := strings.Split(data, "\n")

			var s Snippet
			s.Category = strings.TrimSpace(lines[0])
			s.Name = strings.TrimSpace(lines[2])
			s.Value = strings.TrimSpace(strings.Join(lines[4:], "\n"))

			snippets = append(snippets, &s)
		}

		_, id, err := common.Wofi("Select the Snippet", formatSnippetsWofi(snippets))
		if err != nil {
			return err
		}

		common.Notify("Snippet Copied", fmt.Sprintf("%s - %s", snippets[id].Category, snippets[id].Name))

		return common.CopyToClipboard(snippets[id].Value)
	default:
		return fmt.Errorf("unknown command %s", name)
	}
}

func (m *Snippets) IsEnabled() bool {
	return true
}

func (m *Snippets) SuspendOnBattery() bool {
	return false
}

func GetModule() common.Module {
	return module
}
