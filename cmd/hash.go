package cmd

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"ghid/data"
	"ghid/flags"
	"ghid/output"
	"ghid/utils"

	"github.com/fatih/color"
)

func init() {
	utils.LoadPopularHashes()
}

func showHashValue(args []string) bool {
	found := false
	hashes := utils.ParseJson()

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

func printMode(modes data.Modes) {

	output.PrintColorText(&output.Text{
		Text:           fmt.Sprintf("- %s\n", modes.Name),
		ColorAttribute: color.FgGreen,
		Style:          []color.Attribute{color.Bold},
	})

	if modes.Hashcat != nil {
		output.PrintColorText(&output.Text{
			Text:           fmt.Sprintf("  Hashcat: %d\n", *modes.Hashcat),
			ColorAttribute: color.FgGreen,
			Style:          []color.Attribute{color.Bold},
		})
	}
	if modes.John != nil {
		output.PrintColorText(&output.Text{
			Text:           fmt.Sprintf("  John: %s\n", *modes.John),
			ColorAttribute: color.FgGreen,
			Style:          []color.Attribute{color.Bold},
		})
	}
}

func isSimpleHash(name string) bool {
	_, ok := utils.PopularHashesSet[strings.ToLower(name)]
	return ok
}

func printModeField(label string, name *string) {
	if name == nil {
		output.PrintColorText(&output.Text{
			Text:           fmt.Sprintf("  %s: not available\n", label),
			ColorAttribute: color.FgYellow,
			Style:          []color.Attribute{color.Bold},
		})
		return
	}

	output.PrintColorText(&output.Text{
		Text:           fmt.Sprintf("  %s: %s\n", label, *name),
		ColorAttribute: color.FgGreen,
		Style:          []color.Attribute{color.Bold},
	})

}

func uintToStr(uInt *uint) *string {
	if uInt == nil {
		return nil
	}
	str := strconv.FormatUint(uint64(*uInt), 10)
	return &str
}
