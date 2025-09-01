package command

import (
	"ghid/data"
	"ghid/flags"
	"ghid/output"

	"github.com/spf13/cobra"
)

var VersionCmd = &cobra.Command{
	Use:   "version",
	Short: "Show version App",
	Long:  "Show version App",
	Run: func(cmd *cobra.Command, args []string) {
		output.PrintGreenText(data.VERSION)
	},
}

func init() {
	flags.AddBoolFlags(ListCmd)
}
