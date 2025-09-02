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
	var sb strings.Builder

	fmt.Fprintf(&sb, "- %s\n", modes.Name)

	if modes.Hashcat != nil {
		fmt.Fprintf(&sb, "  Hashcat: %d\n", *modes.Hashcat)
	}
	if modes.John != nil {
		fmt.Fprintf(&sb, "  John: %s\n", *modes.John)
	}

	output.PrintGreenText(sb.String())
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
