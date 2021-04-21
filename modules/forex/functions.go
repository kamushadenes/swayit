package forex

import (
	"fmt"
	"github.com/kamushadenes/swayit/common"
	"github.com/kamushadenes/swayit/config"
	"github.com/mattevans/dinero"
	"github.com/rgeoghegan/tabulate"
	"time"
)

var okColor = "#2ECC40"
var nokColor = "#FF4136"

var client *dinero.Client
var names map[string]string

func initForex() {
	client = dinero.NewClient(
		config.SwayItConfig.Forex.OpenExchangeRatesAppID,
		"USD",
		time.Duration(config.SwayItConfig.Forex.Interval)*time.Second,
	)

	names = make(map[string]string)
	rsp, err := client.Currencies.List()
	if err != nil {
		panic(err)
	}

	for _, v := range rsp {
		names[v.Code] = v.Name
	}
}

func run(w *common.WaybarOutput) error {
	if client == nil {
		initForex()
	}
	rates, err := client.Rates.List()
	if err != nil {
		return err
	}

	w.Tooltip = fmt.Sprintf("<b>Source:</b> OpenExchangeRates\n<b>Last Update:</b> %s", module.GetLastRun())

	var rows []*Row

	for symbol, price := range rates.Rates {
		for _, coin := range config.SwayItConfig.Forex.Coins {
			if symbol == coin.Symbol {
				var row Row
				row.Coin = symbol
				if name, ok := names[symbol]; ok {
					row.Name = name
				} else {
					row.Name = "-"
				}

				row.Price = fmt.Sprintf("%s %.2f", config.SwayItConfig.Forex.BaseSymbol, 1/price*rates.Rates[config.SwayItConfig.Forex.BaseCurrency])

				rows = append(rows, &row)

				if !coin.TooltipOnly {
					w.Text += fmt.Sprintf("<span font='8' color='%s'>%s</span> ", okColor, coin.Icon)
					w.Alt += fmt.Sprintf("<span font='8'>%s %s</span> ", coin.Icon, row.Price)
				}
				break
			}
		}
	}

	t, err := tabulate.Tabulate(rows, &tabulate.Layout{Format: tabulate.FancyGridFormat})
	if err != nil {
		return err
	}

	w.Tooltip += fmt.Sprintf("\n\n%s", t)

	return nil
}
