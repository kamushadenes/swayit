package notes

import (
	"fmt"
	"github.com/kamushadenes/swayit/common"
	"github.com/kamushadenes/swayit/config"
	"io/ioutil"
	"os"
	"path"
	"strings"
)

type Notes struct {
	Name        string
	Description string
	Slug        string
	output      *common.WaybarOutput
}

var module = &Notes{
	Name:        "Notes",
	Description: "Note Taking",
	Slug:        "notes",
}

func (m *Notes) GetName() string {
	return m.Name
}

func (m *Notes) GetDescription() string {
	return m.Description
}

func (m *Notes) GetSlug() string {
	return m.Slug
}

func (m *Notes) GetWaybarOutput() (*common.WaybarOutput, error) {
	return nil, fmt.Errorf("not a waybar module")
}

func (m *Notes) SaveLastRun() {
	common.SaveLastRun(m.Slug)
}
func (m *Notes) GetLastRun() string {
	return common.GetLastRun(m.Slug)
}

func (m *Notes) WriteOutput() error {
	return common.WriteOutput(m.Slug, m.output.ToJSON())
}

func (m *Notes) Run() error {
	m.SaveLastRun()

	_, err := m.GetWaybarOutput()

	if err != nil {
		return err
	}

	return m.WriteOutput()
}

func (m *Notes) GetRunInterval() int64 {
	return 0
}

func (m *Notes) GetRunIntervalOnBattery() int64 {
	return 0
}

func (m *Notes) RunCommand(name string, args []string) error {
	switch name {
	case "meeting":
		md := common.GetData("khal", "currentMeetingMarkdown")
		
		fname := path.Join(config.SwayItConfig.Paths.Notes, "Meetings", md)

		if _, err := os.Stat(fname); os.IsNotExist(err) {
			header := common.GetData("khal", "currentMeetingHeader")
			err = ioutil.WriteFile(fname, []byte(header), 0644)
			if err != nil {
				return err
			}
		}

		err := common.OpenEditor(fname)
		if err != nil {
			return err
		}
	case "pick":
		folders, err := common.ListDirectories(config.SwayItConfig.Paths.Notes)
		if err != nil {
			return err
		}

		_, id, err := common.Wofi("Select the Folder / Category", folders)
		if err != nil {
			return err
		}

		folder := path.Join(config.SwayItConfig.Paths.Notes, folders[id])

		files, err := common.ListFiles(folder)
		if err != nil {
			return err
		}

		name, _, err := common.Wofi("Select or Create Note", files)
		if err != nil {
			return err
		}

		fname := path.Join(folder, name)

		if _, err = os.Stat(fname); os.IsNotExist(err) {
			if !strings.HasSuffix(fname, ".md") {
				fname = fmt.Sprintf("%s.md", fname)
			}
			header := fmt.Sprintf("# %s: %s", folders[id], strings.Title(strings.ToLower(strings.TrimSuffix(name, ".md"))))
			err = ioutil.WriteFile(fname, []byte(header), 0644)
			if err != nil {
				return err
			}
		}

		err = common.OpenEditor(fname)
		if err != nil {
			return err
		}
	default:
		return fmt.Errorf("unknown command %s", name)
	}
	
	return nil
}

func (m *Notes) IsEnabled() bool {
	return true
}

func (m *Notes) SuspendOnBattery() bool {
	return false
}

func GetModule() common.Module {
	return module
}
