package version

import (
	"ghid/data"
	"ghid/flags"
	"ghid/output"

	"github.com/spf13/cobra"
)

func Commands() []*cobra.Command {
	var versionCmd = &cobra.Command{
		Use:   "version",
		Short: "Show version App",
		Long:  "Show version App",
		Run: func(cmd *cobra.Command, args []string) {
			output.PrintGreenText(data.VERSION)
		},
	}

	flags.AddCommandFlags(versionCmd, data.CMD_VERSION)

	return []*cobra.Command{versionCmd}
}
