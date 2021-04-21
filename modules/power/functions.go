package power

import (
	"fmt"
	"github.com/kamushadenes/swayit/common"
	"github.com/kamushadenes/swayit/config"
	"io/ioutil"
	"path"
	"strconv"
	"strings"
)

var powerSupplyPath = "/sys/class/power_supply"

func run(w *common.WaybarOutput) error {
	b, err := ioutil.ReadFile(path.Join(powerSupplyPath, config.SwayItConfig.Power.Battery, "current_now"))
	if err != nil {
		return err
	}

	currentNow, err := strconv.ParseInt(strings.TrimSpace(string(b)), 10, 64)
	if err != nil {
		return err
	}

	b, err = ioutil.ReadFile(path.Join(powerSupplyPath, config.SwayItConfig.Power.Battery, "voltage_now"))
	if err != nil {
		return err
	}

	voltageNow, err := strconv.ParseInt(strings.TrimSpace(string(b)), 10, 64)
	if err != nil {
		return err
	}

	watts := currentNow * voltageNow / 1000000000000

	w.Text = fmt.Sprintf("\uF0E7 %dW", watts)

	w.Tooltip = fmt.Sprintf("<b>Current:</b> %d mA", currentNow/1000)
	w.Tooltip += fmt.Sprintf("\n<b>Voltage:</b> %d V", voltageNow/1000000)
	w.Tooltip += fmt.Sprintf("\n<b>Watts:</b> %d W", watts)

	return nil
}
