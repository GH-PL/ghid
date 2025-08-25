package command

import (
	"fmt"
	"strings"

	"github.com/GH-PL/ghid/flags"
	"github.com/GH-PL/ghid/utils"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var SamplesCmd = &cobra.Command{
	Use:   "samples [NAME HASH]",
	Short: "Display hash samples for the given type",
	Long:  "Display hash samples for the given type",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			utils.PrintColorText(&utils.Text{
				Text:           "Error: missing hashType argument",
				ColorAttribute: color.FgRed,
				Style:          []color.Attribute{color.Bold},
			})
			return
		}
		samples(args[0])
	},
}

func init() {
	flags.AddBoolFlags(SamplesCmd, flags.BoolFlags)
}

func samples(str string) {
	hash := utils.ParseJson()
	if color.NoColor {
		utils.DisableColorOutput()
	}
	found := false

	for _, hashValue := range hash {
		for _, mode := range hashValue.Modes {
			if strings.EqualFold(str, mode.Name) {
				if len(mode.Samples) == 0 {
					utils.PrintColorText(&utils.Text{
						Text:           "No samples available for this Hash",
						ColorAttribute: color.FgRed,
						Style:          []color.Attribute{color.Bold},
					})
				} else {
					for _, samplesValue := range mode.Samples {
						utils.PrintColorText(&utils.Text{
							Text:           samplesValue,
							ColorAttribute: color.FgBlue,
							Style:          []color.Attribute{color.Bold},
						})
					}
				}

				found = true
			}
		}
	}
	if !found {
		utils.PrintColorText(&utils.Text{
			Text:           fmt.Sprintf("Not found %s", str),
			ColorAttribute: color.FgRed,
			Style:          []color.Attribute{color.Bold},
		})
	}
}
