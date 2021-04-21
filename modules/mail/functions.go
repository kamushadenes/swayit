package mail

import (
	"fmt"
	"github.com/kamushadenes/swayit/common"
	"github.com/kamushadenes/swayit/config"
	"os/exec"
	"strings"
)

func run(w *common.WaybarOutput) error {
	mainCmd := exec.Command(config.SwayItConfig.Mail.MuPath, "find", config.SwayItConfig.Mail.MainQuery)
	b, err := mainCmd.Output()
	if err != nil {
		return nil
	}

	mainCount := len(strings.Split(strings.TrimSpace(string(b)), "\n"))

	if mainCount > 0 {
		w.Text = fmt.Sprintf("\uF0E0 %d", mainCount)

		for k, q := range config.SwayItConfig.Mail.Queries {
			cmd := exec.Command(config.SwayItConfig.Mail.MuPath, "find", q)
			b, err = cmd.Output()
			if err != nil {
				continue
			}
			cnt := len(strings.Split(strings.TrimSpace(string(b)), "\n"))

			w.Tooltip += fmt.Sprintf(" %s: %d unread mail\n", k, cnt)
		}
	}

	return nil
}
