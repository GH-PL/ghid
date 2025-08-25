package command

import (
	"strings"

	"ghid/err"
	"ghid/flags"
	"ghid/utils"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var SamplesCmd = &cobra.Command{
	Use:   "samples [NAME HASH]",
	Short: "Display hash samples for the given type",
	Long:  "Display hash samples for the given type",
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return err.ErrEmptyArgument
		}
		return samples(args[0])
	},
}

func init() {
	flags.AddBoolFlags(SamplesCmd, flags.BoolFlags)
}

func samples(str string) error {
	hash := utils.ParseJson()
	if color.NoColor {
		utils.DisableColorOutput()
	}

	for _, hashValue := range hash {
		for _, mode := range hashValue.Modes {

			if !strings.EqualFold(str, mode.Name) {
				continue
			}
			if len(mode.Samples) == 0 {
				return err.ErrNoSamplesFound
			}
			for _, samplesValue := range mode.Samples {
				utils.PrintColorText(&utils.Text{
					Text:           samplesValue,
					ColorAttribute: color.FgBlue,
					Style:          []color.Attribute{color.Bold},
				})
			}
			return nil
		}
	}
	return err.ErrNoSamplesFound
}
