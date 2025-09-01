package detect

import (
	"regexp"

	"ghid/data"
	"ghid/errHandler"
	"ghid/flags"
	"ghid/output"
	"ghid/utils"

	"github.com/spf13/cobra"
)

func Commands() []*cobra.Command {
	var detectCmd = &cobra.Command{
		Use:   "detect [flags] <hash>",
		Short: "Identify the most probable hash type",
		Long:  "Identify the most probable hash type",
		Run: func(cmd *cobra.Command, args []string) {
			matchHash := matchHashTypes(args)
			if !matchHash {
				output.PrintError(errHandler.ErrNotFoundHash)
			} else {
				if !flags.Extended {
					output.PrintWarning("You need extended mode")
				}
			}
		},
	}
	flags.AddCommandFlags(detectCmd, "detect")
	return []*cobra.Command{detectCmd}
}

func matchHashTypes(args []string) bool {
	found := false
	hashes := utils.ParseJson(data.WAY_DATA_JSON)

	for _, hashValue := range hashes {
		for _, valueArgs := range args {
			match, _ := regexp.MatchString(hashValue.Regex, valueArgs)

			if !match {
				continue
			}
			found = true
			for _, modes := range hashValue.Modes {

				if !flags.Extended && !isSimpleHash(modes.Name) {
					continue
				}
				switch {
				case flags.ShortFlag:
					printModeField("Name", &modes.Name)
				case flags.Hashcat:
					printModeField("Hashcat", uintToStr(modes.Hashcat))
				case flags.John:
					printModeField("John", modes.John)

				default:
					printMode(modes)
				}

			}
		}
	}
	return found
}
