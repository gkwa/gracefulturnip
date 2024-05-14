package cmd

import (
	"github.com/spf13/cobra"
	"github.com/taylormonacelli/gracefulturnip/core"
)

// run2Cmd represents the run2 command
var run2Cmd = &cobra.Command{
	Use:   "test2",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		core.Test2()
	},
}

func init() {
	rootCmd.AddCommand(run2Cmd)
}
