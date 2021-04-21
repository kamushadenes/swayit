package forex

import (
	"fmt"
	"github.com/kamushadenes/swayit/common"
	"github.com/kamushadenes/swayit/config"
	"sort"
	"strconv"
	"strings"
)

type Forex struct {
	Name        string
	Description string
	Slug        string
	output      *common.WaybarOutput
}

var module = &Forex{
	Name:        "Forex",
	Description: "FIAT Ticker",
	Slug:        "forex",
	output:      &common.WaybarOutput{Class: "forex"},
}

func (m *Forex) GetName() string {
	return m.Name
}

func (m *Forex) GetDescription() string {
	return m.Description
}

func (m *Forex) GetSlug() string {
	return m.Slug
}

func (m *Forex) GetWaybarOutput() (*common.WaybarOutput, error) {
	m.output.Text = ""
	m.output.Tooltip = ""
	m.output.Alt = ""
	err := run(m.output)
	m.output.Text = strings.TrimSpace(m.output.Text)
	m.output.Tooltip = strings.TrimSpace(m.output.Tooltip)
	m.output.Alt = strings.TrimSpace(m.output.Alt)

	return m.output, err
}

func (m *Forex) SaveLastRun() {
	common.SaveLastRun(m.Slug)
}
func (m *Forex) GetLastRun() string {
	return common.GetLastRun(m.Slug)
}

func (m *Forex) WriteOutput() error {
	return common.WriteOutput(m.Slug, m.output.ToJSON())
}

func (m *Forex) Run() error {
	m.SaveLastRun()

	_, err := m.GetWaybarOutput()

	if err != nil {
		return err
	}

	return m.WriteOutput()
}

func (m *Forex) GetRunInterval() int64 {
	return config.SwayItConfig.Forex.Interval
}

func (m *Forex) GetRunIntervalOnBattery() int64 {
	return config.SwayItConfig.Forex.IntervalOnBattery
}

func (m *Forex) RunCommand(name string, args []string) error {
	switch name {
	case "convert":
		out, _, err := common.WofiInput("Enter the amount to convert", []string{"1.00"}, "--lines", "1")
		if err != nil {
			return err
		}
		
		f, err := strconv.ParseFloat(strings.TrimSpace(out), 64)
		if err != nil {
			return err
		}
		
		initForex()

		rates, err := client.Rates.List()
		if err != nil {
			return err
		}
		
		var currencies []string
		for coin, name := range names {
			currencies = append(currencies, fmt.Sprintf("<b>%s</b>\n%s", name, coin))
		}
		sort.Slice(currencies, func(i,j int) bool {
			return currencies[i] < currencies[j]
		})

		out, _, err = common.WofiInput("Select the source currency for conversion", currencies)
		if err != nil {
			return err
		}
		
		base := strings.Split(out, "\n")[1]

		var conversions []string
		for coin, rate := range rates.Rates {
			conversions = append(conversions, fmt.Sprintf("<b>%s</b>\n%s\n%.2f", names[coin], coin, 1/(1/rate*rates.Rates[base])*f))
		}
		sort.Slice(conversions, func(i,j int) bool {
			return conversions[i] < conversions[j]
		})

		out, _, err = common.WofiInput("Select the destination currency for conversion", conversions)
		if err != nil {
			return err
		}
		
		res := strings.Split(out, "\n")[2]

		return common.CopyToClipboard(res)
	default:
		return fmt.Errorf("unknown command %s", name)
	}
}

func (m *Forex) IsEnabled() bool {
	return config.SwayItConfig.Forex.Enabled
}

func (m *Forex) SuspendOnBattery() bool {
	return config.SwayItConfig.Forex.SuspendOnBattery
}

func GetModule() common.Module {
	return module
}
