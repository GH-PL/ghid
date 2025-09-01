package detect

import (
	"fmt"
	"ghid/data"
	"ghid/output"
	"ghid/utils"
	"strconv"
	"strings"
)

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
	_, ok := utils.ParseCsv(data.WAY_POPULAR_HASH_CSV)[strings.ToLower(name)]
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
