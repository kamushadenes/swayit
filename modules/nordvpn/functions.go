package nordvpn

import (
	"fmt"
	"github.com/kamushadenes/swayit/common"
	"github.com/kamushadenes/swayit/config"
	"os/exec"
	"strings"
)

func run(w *common.WaybarOutput) error {
	cmd := exec.Command(config.SwayItConfig.NordVPN.Command, "status")

	output, err := cmd.Output()
	if err != nil {
		return err
	}

	nvpnStatus := strings.Replace(string(output), "\r", "", -1)
	nvpnStatus = strings.TrimPrefix(nvpnStatus, "-")
	nvpnStatus = strings.TrimSuffix(nvpnStatus, "-")
	nvpnStatus = strings.TrimSpace(nvpnStatus)

	var status string
	var city string

	statusSlice := statusRegex.FindStringSubmatch(nvpnStatus)
	citySlice := cityRegex.FindStringSubmatch(nvpnStatus)

	if len(status) == 1 {
		status = statusSlice[0]
	}

	if len(city) == 1 {
		city = citySlice[0]
	}

	if status == "Connected" {
		w.Text = fmt.Sprintf("\uf023 %s", city)
		w.Class = "connected"
	} else {
		w.Class = "disconnected"
		w.Text = "\uf09c Disconnected"
		w.Tooltip = nvpnStatus
	}

	return nil
}
