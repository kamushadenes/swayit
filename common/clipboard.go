package common

import (
	"io"
	"os/exec"
	"strings"
)

func CopyToClipboard(text string) error {
	cmd := exec.Command("wl-copy")

	stdin, err := cmd.StdinPipe()
	if err != nil {
		return err
	}

	go func() {
		defer stdin.Close()
		_, _ = io.WriteString(stdin, text)
	}()

	err = cmd.Run()
	if err != nil {
		return err
	}

	return nil
}

func PasteFromClipboard() (string, error) {
	cmd := exec.Command("wl-paste", "-n", "-p")

	out, err := cmd.Output()

	return strings.TrimSpace(string(out)), err
}