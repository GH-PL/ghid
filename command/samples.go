package command

import (
	"errors"
	"strings"

	"ghid/err"
	"ghid/flags"
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
			err.Show(err.ErrEmptyArgument)
			cmd.Usage()
			return nil
		}
		if errSamples := samples(args[0]); errSamples != nil {
			if errors.Is(errSamples, err.ErrNoSamplesFound) {
				err.Show(err.ErrNoSamplesFound)
				return nil
			}
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
