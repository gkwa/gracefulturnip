package cmd

import (
	"github.com/spf13/cobra"
	"github.com/taylormonacelli/gracefulturnip/core"
)

// run2Cmd represents the run2 command
var run2Cmd = &cobra.Command{
	Use:   "run2",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		core.Run2()
	},
}

func init() {
	rootCmd.AddCommand(run2Cmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// run2Cmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// run2Cmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
