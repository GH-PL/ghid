package list

import (
	"ghid/data"
	"ghid/flags"
	"ghid/output"
	"ghid/utils"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

func Commands() []*cobra.Command {
	var listCmd = &cobra.Command{
		Use:   "list",
		Short: "Show all Hash name",
		Long:  "Show all Hash name",
		Run: func(cmd *cobra.Command, args []string) {
			showList()
		},
	}

	flags.AddCommandFlags(listCmd, data.CMD_LIST)

	return []*cobra.Command{listCmd}
}

func showList() {
	hash := utils.ParseJson(data.WAY_DATA_JSON)
	if color.NoColor {
		output.DisableColorOutput()
	}
	for _, hashValue := range hash {
		for _, mode := range hashValue.Modes {
			output.PrintBlueText(mode.Name)
		}
	}
}
