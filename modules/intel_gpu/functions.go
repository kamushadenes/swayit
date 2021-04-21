package intel_gpu

import (
	"bufio"
	"encoding/json"
	"fmt"
	"github.com/kamushadenes/swayit/common"
	"github.com/kamushadenes/swayit/config"
	"os/exec"
	"strings"
)

func run(w *common.WaybarOutput) error {
	cmd := exec.Command(config.SwayItConfig.IntelGPU.SudoCommand, config.SwayItConfig.IntelGPU.Command, "-J", "-s", fmt.Sprintf("%d", config.SwayItConfig.IntelGPU.Interval*1000))

	stdout, _ := cmd.StdoutPipe()
	err := cmd.Start()
	if err != nil {
		return err
	}

	scanner := bufio.NewScanner(stdout)

	var msg string
	var out Output

	var openCount = 0
	var closeCount = 0

	for scanner.Scan() {
		m := strings.TrimSpace(scanner.Text())

		for _, v := range m {
			switch v {
			case '{':
				openCount++
			case '}':
				closeCount++
			}
		}

		msg += m

		if closeCount == openCount && (openCount > 0 && closeCount > 0) {
			msg = strings.TrimSuffix(msg, ",")
			err := json.Unmarshal([]byte(msg), &out)
			if err != nil {
				return err
			} else {
				w.Text = fmt.Sprintf("\uf729 %.0f%%", out.Engines.Render3D0.Busy)

				w.Tooltip = fmt.Sprintf("<b>Frequency:</b>  %.2f %s", out.Frequency.Actual, out.Frequency.Unit)
				w.Tooltip += fmt.Sprintf("\n<b>Interrupts:</b> %.2f %s", out.Interrupts.Count, out.Interrupts.Unit)
				w.Tooltip += fmt.Sprintf("\n<b>Power:</b>      %.2f %s", out.Power.Value, out.Power.Unit)
				w.Tooltip += fmt.Sprintf("\n<b>IMC Reads:</b>  %.2f %s", out.ImcBandwidth.Reads, out.ImcBandwidth.Unit)
				w.Tooltip += fmt.Sprintf("\n<b>IMC Writes:</b> %.2f %s", out.ImcBandwidth.Writes, out.ImcBandwidth.Unit)

				w.Tooltip += "\n\n<b>Engines</b>\n"
				w.Tooltip += fmt.Sprintf("\n<b>Render:</b>        %.0f%%", out.Engines.Render3D0.Busy)
				w.Tooltip += fmt.Sprintf("\n<b>Blitter:</b>       %.0f%%", out.Engines.Blitter0.Busy)
				w.Tooltip += fmt.Sprintf("\n<b>Video 0:</b>       %.0f%%", out.Engines.Video0.Busy)
				w.Tooltip += fmt.Sprintf("\n<b>Video 1:</b>       %.0f%%", out.Engines.Video1.Busy)
				w.Tooltip += fmt.Sprintf("\n<b>Video Enhance:</b> %.0f%%", out.Engines.Videoenhance0.Busy)

				err := module.WriteOutput()
				if err != nil {
					return err
				}
			}

			openCount = 0
			closeCount = 0
			msg = ""
		}
	}

	cmd.Wait()

	return nil
}
