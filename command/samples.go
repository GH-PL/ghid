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
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			errHandler.Show(errHandler.ErrEmptyArgument)
			cmd.Usage()
			return nil
		}
		if errSamples := samples(args[0]); errSamples != nil {

			errHandler.IsError(&errHandler.IsERROR{
				Err: errSamples,
				Msg: args[0],
			})
			return errSamples
		}
		return nil
	},
}

func init() {
	flags.AddBoolFlags(SamplesCmd, flags.BoolFlags)
}

func samples(str string) error {
	hash := utils.ParseJson()
	if color.NoColor {
		output.DisableColorOutput()
	}

	for _, hashValue := range hash {
		for _, mode := range hashValue.Modes {

			if !strings.EqualFold(str, mode.Name) {
				continue
			}
			if len(mode.Samples) == 0 {
				output.PrintColorText(&output.Text{
					Text:           "No examples found",
					ColorAttribute: color.FgRed,
					Style:          []color.Attribute{color.Bold},
				})
				continue
			}
			for _, samplesValue := range mode.Samples {
				output.PrintColorText(&output.Text{
					Text:           samplesValue,
					ColorAttribute: color.FgBlue,
					Style:          []color.Attribute{color.Bold},
				})
			}
			return nil
		}
	}
	return nil
}
