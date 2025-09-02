package detect

import (
	"fmt"
	"ghid/data"
	"ghid/output"
	"ghid/utils"
	"strconv"
	"strings"
	"sync"
)

func printMode(modes utils.Modes) {
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

var (
	once         sync.Once
	simpleHashes map[string]struct{}
)

func isSimpleHash(name string) bool {

	once.Do(func() {
		simpleHashes = utils.ParseCsv(data.WAY_POPULAR_HASH_CSV)
	})
	_, ok := simpleHashes[strings.ToLower(name)]
	return ok
}

func printIfExists(label string, name string) {
	if name == "" {
		output.PrintWarning(fmt.Sprintf("  %s: not available\n", label))
		return
	}
	output.PrintGreenText(fmt.Sprintf("  %s: %s\n", label, name))
}

func uintToStr(uInt *uint) string {
	if uInt == nil {
		return ""
	}
	return strconv.FormatUint(uint64(*uInt), 10)
}

func toStr(str *string) string {
	if str == nil {
		return ""
	}
	return *str
}
