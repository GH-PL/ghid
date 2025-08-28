package command

import (
	"strings"

	"ghid/errHandler"
	"ghid/flags"
	"ghid/output"
	"ghid/utils"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var SamplesCmd = &cobra.Command{
	Use:           "samples [NAME HASH]",
	Short:         "Display hash samples for the given type",
	Long:          "Display hash samples for the given type",
	SilenceErrors: true,
	SilenceUsage:  true,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			output.PrintError(errHandler.ErrEmptyArgument)
			cmd.Usage()
		} else {
			samples(args[0])
		}
	},
}

func init() {
	flags.AddBoolFlags(SamplesCmd, flags.BoolFlags)
}

func samples(str string) {
	hash := utils.ParseJson()
	if color.NoColor {
		output.DisableColorOutput()
	}

	for _, hashValue := range hash {
		for _, mode := range hashValue.Modes {
			if !strings.EqualFold(str, mode.Name) {
				output.PrintError(errHandler.ErrNotFoundName)
				return
			}
			if len(mode.Samples) == 0 {
				output.PrintError(errHandler.ErrNotExampleFound)
				return
			}
			for _, samplesValue := range mode.Samples {
				output.PrintBlueText(samplesValue)
			}
		}
	}

}
