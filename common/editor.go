package common

import (
	"github.com/kamushadenes/swayit/config"
	"os/exec"
)

func OpenEditor(fname string) error {
	var args []string
	args = append(args, config.SwayItConfig.Editor.Command)
	args = append(args, config.SwayItConfig.Editor.Args...)
	args = append(args, fname)
	
	cmd := exec.Command(config.SwayItConfig.Editor.Terminal, args...)
	
	err := cmd.Run()
	if err != nil {
		return err
	}
	
	return nil
}