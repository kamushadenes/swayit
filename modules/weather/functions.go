package weather

import (
	"encoding/json"
	"fmt"
	"github.com/kamushadenes/swayit/common"
	"github.com/kamushadenes/swayit/config"
	"net/http"
	"strconv"
	"strings"
)

func formatTime(t string) string {
	h := strings.Replace(t, "00", "", -1)
	h = fmt.Sprintf("%02s", h)

	return fmt.Sprintf("%s:00", h)
}

func formatChances(hour *WttrInHour) string {
	var conditions []string
	
	for c := range chances {
		var i int64
		var err error

		switch c {
		case "chanceoffog":
			i, err = strconv.ParseInt(hour.Chanceoffog, 10, 64)
		case "chanceoffrost":
			i, err = strconv.ParseInt(hour.Chanceoffrost, 10, 64)
		case "chanceofovercast":
			i, err = strconv.ParseInt(hour.Chanceofovercast, 10, 64)
		case "chanceofrain":
			i, err = strconv.ParseInt(hour.Chanceofrain, 10, 64)
		case "chanceofsnow":
			i, err = strconv.ParseInt(hour.Chanceofsnow, 10, 64)
		case "chanceofsunshine":
			i, err = strconv.ParseInt(hour.Chanceofsunshine, 10, 64)
		case "chanceofthunder":
			i, err = strconv.ParseInt(hour.Chanceofthunder, 10, 64)
		case "chanceofwindy":
			i, err = strconv.ParseInt(hour.Chanceofwindy, 10, 64)
		}

		if err != nil {
			common.Logger.Error().Err(err).Msg("an error has occurred")
			continue
		}
		
		if i > 0 {
			conditions = append(conditions, fmt.Sprintf("%s %d%%", chances[c], i))
		}
	}

	return strings.Join(conditions, ", ")
}

func getWeather() (*WttrIn, error) {
	resp, err := http.Get(config.SwayItConfig.Weather.URL)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var t WttrIn

	err = json.NewDecoder(resp.Body).Decode(&t)
	if err != nil {
		return nil, err
	}

	return &t, nil
}

func run(w *common.WaybarOutput) error {
	weather, err := getWeather()
	if err != nil {
		return err
	}

	w.Text = fmt.Sprintf("%s %s°C",
		weatherCodes[weather.CurrentCondition[0].Weathercode],
		weather.CurrentCondition[0].Feelslikec)

	w.Tooltip = "<b>Source:</b> wttr.in"
	w.Tooltip += fmt.Sprintf("\n<b>Last Update:</b> %s", module.GetLastRun())
	w.Tooltip += "\n\n"
	w.Tooltip += fmt.Sprintf("<b>%s %s°C</b>\n", weather.CurrentCondition[0].Weatherdesc[0].Value, weather.CurrentCondition[0].TempC)
	w.Tooltip += fmt.Sprintf("Feels like: %s°C\n", weather.CurrentCondition[0].Feelslikec)
	w.Tooltip += fmt.Sprintf("Wind: %s Km/h\n", weather.CurrentCondition[0].Windspeedkmph)
	w.Tooltip += fmt.Sprintf("Humidity: %s%%\n", weather.CurrentCondition[0].Humidity)

	for k, day := range weather.Weather {
		w.Tooltip += "\n<b>"
		if k == 0 {
			w.Tooltip += "Today, "
		} else if k == 1 {
			w.Tooltip += "Tomorrow, "
		}

		w.Tooltip += fmt.Sprintf("%s</b>\n", day.Date)
		w.Tooltip += fmt.Sprintf("⬆️ %s°C ⬇️ %s°C", day.Maxtempc, day.Mintempc)
		w.Tooltip += fmt.Sprintf("🌅 %s 🌇 %s\n", day.Astronomy[0].Sunrise, day.Astronomy[0].Sunset)

		for _, hour := range day.Hourly {
			if k == 0 {
				// format time
			}
			w.Tooltip += fmt.Sprintf("%s %s %s %s, %s\n", formatTime(hour.Time), weatherCodes[hour.Weathercode], fmt.Sprintf("%s°C", hour.Feelslikec), hour.Weatherdesc[0].Value, formatChances(hour))
		}
	}
	return nil
}
