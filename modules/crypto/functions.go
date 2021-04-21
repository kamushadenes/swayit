package crypto

import (
	"fmt"
	"github.com/kamushadenes/swayit/common"
	"github.com/kamushadenes/swayit/config"
	cmc "github.com/miguelmota/go-coinmarketcap/pro/v1"
	"github.com/rgeoghegan/tabulate"
)

var okColor = "#2ECC40"
var nokColor = "#FF4136"

var client *cmc.Client

func formatChange(c float64) string {
	if c >= 0 {
		return fmt.Sprintf("+%.2f%%", c)
	} else {
		return fmt.Sprintf("%.2f%%", c)
	}
}

func run(w *common.WaybarOutput) error {
	if client == nil {
		client = cmc.NewClient(&cmc.Config{
			ProAPIKey: config.SwayItConfig.Crypto.CMCApiKey,
		})
	}

	listings, err := client.Cryptocurrency.LatestListings(&cmc.ListingOptions{
		Convert: config.SwayItConfig.Crypto.BaseCurrency,
	})
	if err != nil {
		return err
	}

	w.Tooltip = fmt.Sprintf("<b>Source:</b> CoinMarketCap\n<b>Last Update:</b> %s", module.GetLastRun())

	var rows []*Row

	for _, listing := range listings {
		for _, coin := range config.SwayItConfig.Crypto.Coins {
			if listing.Symbol == coin.Symbol {
				quote := listing.Quote[config.SwayItConfig.Crypto.BaseCurrency]
				var row Row
				row.Coin = listing.Symbol
				row.Name = listing.Name
				row.Price = fmt.Sprintf("$ %.2f", quote.Price)
				row.Change1h = formatChange(quote.PercentChange1H)
				row.Change24h = formatChange(quote.PercentChange24H)
				row.Change7d = formatChange(quote.PercentChange7D)

				rows = append(rows, &row)

				if !coin.TooltipOnly {
					if quote.PercentChange24H > 0 {
						w.Text += fmt.Sprintf("<span font='8' color='%s'>%s</span> ", okColor, coin.Icon)
					} else {
						w.Text += fmt.Sprintf("<span font='8' color='%s'>%s</span> ", nokColor, coin.Icon)
					}
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
