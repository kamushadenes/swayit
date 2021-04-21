package cmd

import (
	"fmt"
	"github.com/kamushadenes/swayit/common"
	"github.com/spf13/cobra"
)

var outputCmd = &cobra.Command{
	Use: "output",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) != 1 {
			return fmt.Errorf("must have one argument")
		}

		return nil
	},
	Short: "Get a module output",
	Run: func(cmd *cobra.Command, args []string) {
		follow, _ := cmd.PersistentFlags().GetBool("follow")
		if follow {
			err := common.GetOutputFollow(args[0])
			if err != nil {
				common.Logger.Error().Err(err).Msg("an error has occurred")
			}
		} else {
			out, err := common.GetOutput(args[0])
			if err != nil {
				common.Logger.Error().Err(err).Msg("an error has occurred")
			} else {
				fmt.Println(out)
			}
		}
	},
}

func init() {
	outputCmd.PersistentFlags().BoolP("follow", "f", false, "Follow the output")
	rootCmd.AddCommand(outputCmd)
}
