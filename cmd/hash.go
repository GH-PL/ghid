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
)

func matchHashTypes(args []string) bool {
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
	output.PrintGreenText(fmt.Sprintf("- %s\n", modes.Name))

	if modes.Hashcat != nil {
		output.PrintGreenText(fmt.Sprintf("  Hashcat: %d\n", *modes.Hashcat))
	}
	if modes.John != nil {
		output.PrintGreenText(fmt.Sprintf("  John: %s\n", *modes.John))
	}
}

func isSimpleHash(name string) bool {
	_, ok := utils.ParceCsv()[strings.ToLower(name)]
	return ok
}

func printModeField(label string, name *string) {
	if name == nil {
		output.PrintWarning(fmt.Sprintf("  %s: not available\n", label))
		return
	}
	output.PrintGreenText(fmt.Sprintf("  %s: %s\n", label, *name))
}

func uintToStr(uInt *uint) *string {
	if uInt == nil {
		return nil
	}
	str := strconv.FormatUint(uint64(*uInt), 10)
	return &str
}
