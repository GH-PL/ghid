package command

import (
	"github.com/GH-PL/ghid/flags"
	"github.com/GH-PL/ghid/utils"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var ListCmd = &cobra.Command{
	Use:   "list",
	Short: "Show all Hash name",
	Long:  "Show all Hash name",
	Run: func(cmd *cobra.Command, args []string) {
		showList()
	},
}

func init() {
	flags.AddBoolFlags(ListCmd, flags.BoolFlags)
}

func showList() {
	hash := utils.ParseJson()
	if color.NoColor {
		utils.DisableColorOutput()
	}
	for _, hashValue := range hash {
		for _, mode := range hashValue.Modes {
			utils.PrintColorText(&utils.Text{
				Text:           mode.Name,
				ColorAttribute: color.FgBlue,
				Style:          []color.Attribute{color.Bold},
			})
		}
	}
}
