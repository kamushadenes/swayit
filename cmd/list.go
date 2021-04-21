package cmd

import (
	"fmt"
	"github.com/kamushadenes/swayit/modules"
	"sort"

	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List available modules",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("%-12s %-12s %s\n\n", "Name", "Slug", "Description")
		
		sort.Slice(modules.Modules, func(i, j int) bool {
			return modules.Modules[i].GetName() < modules.Modules[j].GetName()
		})
		
		for _, m := range modules.Modules {
			fmt.Printf("%-12s %-12s %s\n", m.GetName(), m.GetSlug(), m.GetDescription())
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
