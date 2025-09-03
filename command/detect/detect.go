package detect

import (
	"regexp"
	"sync"

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
	flags.AddCommandFlags(detectCmd, data.CMD_DETECT)
	return []*cobra.Command{detectCmd}
}

var (
	hashesOnce sync.Once
	hashes     []utils.Hash
)

func loadJsonHash() []utils.Hash {
	hashesOnce.Do(func() {
		hashes = utils.ParseJson(data.WAY_DATA_JSON)
	})
	return hashes
}

func matchHashTypes(args []string) bool {
	found := false
	loadJsonHash()

	for _, arg := range args {
		if matched := search(arg, compiledHashes); matched {
			found = true
		}
	}
	return found
}

func printModeByFlags(mode utils.Modes) {
	if flags.Hashcat {
		printIfExists("Hashcat", uintToStr(mode.Hashcat))
		return
	}
	if flags.John {
		printIfExists("John", toStr(mode.John))
		return
	}
	if flags.ShortFlag {
		printIfExists("Name", mode.Name)
		return
	}
	printMode(mode)
}

type CompiledHash struct {
	Modes []utils.Modes
	Re    *regexp.Regexp
}

func loadCompiledHashes() []CompiledHash {
	raw := utils.ParseJson(data.WAY_DATA_JSON)
	var list []CompiledHash
	for _, h := range raw {
		re := regexp.MustCompile(h.Regex) // компиляция единожды
		list = append(list, CompiledHash{Modes: h.Modes, Re: re})
	}
	return list
}

var compiledHashes []CompiledHash = loadCompiledHashes()

func search(arg string, hashes []CompiledHash) bool {
	for _, hv := range hashes {
		if !hv.Re.MatchString(arg) {
			continue
		}
		for _, mode := range hv.Modes {
			if !flags.Extended && !isSimpleHash(mode.Name) {
				continue
			}
			printModeByFlags(mode)
		}
		return true
	}
	return false
}

/*
	func search(arg string, hashes []utils.Hash) bool {
		for _, hashValue := range hashes {
			if match, _ := regexp.MatchString(hashValue.Regex, arg); !match {
				continue
			}
			for _, mode := range hashValue.Modes {
				if !flags.Extended && !isSimpleHash(mode.Name) {
					continue
				}
				printModeByFlags(mode)
			}
			return true

		}
		return false
	}
*/
