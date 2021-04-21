package common

import (
	"io"
	"os/exec"
	"strconv"
	"strings"
)

func Wofi(prompt string, items []string, args ...string) (string, int64, error) {
	return WofiInput(prompt, items, append(args, "-Ddmenu-print_line_num=true")...)
}

func WofiInput(prompt string, items []string, args ...string) (string, int64, error) {
	var nargs []string
	nargs = append(nargs, "-d", "-p", prompt, "-k", "/dev/null", "-Ddmenu-separator=|")
	nargs = append(nargs, args...)

	cmd := exec.Command("wofi", nargs...)

	if len(items) > 0 {
		stdin, err := cmd.StdinPipe()
		if err != nil {
			return "", 0, err
		}
		go func() {
			defer stdin.Close()
			_, _ = io.WriteString(stdin, strings.Join(items, "|"))
		}()
	}

	out, err := cmd.Output()

	strOut := strings.TrimSpace(string(out))

	if err != nil {
		return strOut, 0, err
	}

	i, err := strconv.ParseInt(strOut, 10, 64)
	if err == nil {
		if int64(len(items)) > i {
			return items[i], i, err
		} else {
			return strOut, i, err
		}
	}

	return strOut, 0, nil
}

func WofiConfirm(prompt string) bool {
	out, _, _ := WofiInput(prompt, []string{"Yes", "No"})

	if out == "Yes" {
		return true
	}
	
	return false
}