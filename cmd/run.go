package cmd

import (
	"fmt"
	"github.com/kamushadenes/swayit/common"
	"github.com/kamushadenes/swayit/modules"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/spf13/cobra"
)

// runCmd represents the run command
var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Run a module",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) != 1 {
			return fmt.Errorf("must have one argument")
		}

		return nil
	},
	Long: `Run a module and save it's output to "paths.output".`,
	Run: func(cmd *cobra.Command, args []string) {
		forever, _ := cmd.PersistentFlags().GetBool("forever")

		var toRun []common.Module

		for k := range modules.Modules {
			m := modules.Modules[k]
			if m.GetSlug() == args[0] || args[0] == "all" {
				toRun = append(toRun, m)
			}
		}
		
		if len(toRun) == 0 {
			common.Logger.Error().Err(fmt.Errorf("module not found")).Str("module", args[0]).Msg("an error has occurred")
		}

		done := make(chan bool)
		for k := range toRun {
			m := toRun[k]
			if !forever {
				common.Logger.Info().Str("module", m.GetName()).Msg("running")
				err := m.Run()
				if err != nil {
					common.Logger.Error().Err(err).Msg("an error has occurred")
				}
			} else {
				if m.GetSlug() == "intelGpu" {
					go func() {
						if m.IsEnabled() {
							err := m.Run()
							if err != nil {
								common.Logger.Error().Err(err).Msg("an error has occurred")
							}
						}
					}()
					continue
				}
				
				if m.GetRunInterval() == 0 || m.GetRunIntervalOnBattery() == 0 {
					continue
				}

				common.Logger.Info().Str("module", m.GetName()).Int64("every", m.GetRunInterval()).Msg("running")
				
				ticker := time.NewTicker(time.Duration(m.GetRunInterval()) * time.Second)
				batTicker := time.NewTicker(time.Duration(m.GetRunIntervalOnBattery()) * time.Second)

				go func() {
					err := m.Run()
					if err != nil {
						common.Logger.Error().Err(err).Msg("an error has occurred")
					}
					for {
						select {
						case <-done:
							return
						case <-ticker.C:
							if m.IsEnabled() {
								if common.IsACConnected() {
									err = m.Run()
									if err != nil {
										common.Logger.Error().Err(err).Msg("an error has occurred")
									}
								}
							}
						case <-batTicker.C:
							if m.IsEnabled() {
								if !common.IsACConnected() {
									if !m.SuspendOnBattery() {
										err = m.Run()
										if err != nil {
											common.Logger.Error().Err(err).Msg("an error has occurred")
										}
									}
								}
							}
						}
					}
				}()
			}
		}

		if forever {
			signalChan := make(chan os.Signal, 1)
			signal.Notify(signalChan,
				syscall.SIGHUP,
				syscall.SIGINT,
				syscall.SIGTERM,
				syscall.SIGQUIT)
			exitChan := make(chan int)
			go func() {
				for {
					select {
					case <-signalChan:
						done <- true
						exitChan <- 0
					}
				}
			}()

			code := <-exitChan
			os.Exit(code)
		}
	},
}

func init() {
	runCmd.PersistentFlags().BoolP("forever", "f", false, "Run forever according to the interval configured")
	rootCmd.AddCommand(runCmd)
}
