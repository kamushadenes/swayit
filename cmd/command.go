package cmd

import (
	"fmt"
	"github.com/kamushadenes/swayit/common"
	"github.com/kamushadenes/swayit/modules"

	"github.com/spf13/cobra"
)

// commandCmd represents the command command
var commandCmd = &cobra.Command{
	Use:   "command",
	Short: "Run a command from a module",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 2 {
			return fmt.Errorf("must have at least two arguments")
		}

		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		for k := range modules.Modules {
			m := modules.Modules[k]
			if m.GetSlug() == args[0] {
				if len(args) > 2 {
					err := m.RunCommand(args[1], args[2:])
					if err != nil {
						common.Logger.Error().Err(err).Msg("an error has occurred")
					}
				} else {
					err := m.RunCommand(args[1], nil)
					if err != nil {
						common.Logger.Error().Err(err).Msg("an error has occurred")
					}
				}
				return
			}
		}

		common.Logger.Error().Err(fmt.Errorf("module not found")).Str("module", args[0]).Msg("an error has occurred")
	},
}

func init() {
	rootCmd.AddCommand(commandCmd)
}
