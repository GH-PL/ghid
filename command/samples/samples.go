package samples

import (
	"strings"

	"ghid/data"
	"ghid/errHandler"
	"ghid/flags"
	"ghid/output"
	"ghid/utils"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

func Commands() []*cobra.Command {

	var samplesCmd = &cobra.Command{
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

	flags.AddCommandFlags(samplesCmd, data.CMD_SAMPLES)

	return []*cobra.Command{samplesCmd}
}

func samples(str string) {
	hash := utils.ParseJson(data.WAY_DATA_JSON)
	if color.NoColor {
		output.DisableColorOutput()
	}

	for _, hashValue := range hash {
		for _, mode := range hashValue.Modes {
			if strings.EqualFold(str, mode.Name) {
				if len(mode.Samples) == 0 {
					output.PrintError(errHandler.ErrNotExampleFound)
					return
				}

				var out strings.Builder
				for _, samplesValue := range mode.Samples {
					out.WriteString(samplesValue + "\n")
				}
				output.PrintBlueText(out.String())
				return
			}
		}
	}
	output.PrintError(errHandler.ErrNotFoundName)
}
