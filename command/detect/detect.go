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
		Long:  "Identify the most probable hash type\nExample:\nghid detect e99a18c428cb38d5f260853678922e03",
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

type Entry struct {
	Compiled *regexp.Regexp
	Name     string
	HashCat  *uint
	John     *string
}

var (
	entries     []Entry
	hashesOnce  sync.Once
	entriesOnce sync.Once
	hashes      []utils.Hash
)

func loadJsonHash() []utils.Hash {
	hashesOnce.Do(func() {
		hashes = utils.ParseJson(data.WAY_DATA_JSON)
	})
	return hashes
}

func Compiled() []Entry {
	var compiledEntry []Entry
	hashes = loadJsonHash()
	for _, hash := range hashes {
		compiled, err := regexp.Compile(hash.Regex)
		if err != nil {
			output.PrintError(errHandler.ErrNotReadFile)
			continue
		}

		for _, mode := range hash.Modes {
			compiledEntry = append(compiledEntry, Entry{
				Compiled: compiled,
				Name:     mode.Name,
				HashCat:  mode.Hashcat,
				John:     mode.John,
			})
		}
	}
	return compiledEntry
}

func matchHashTypes(args []string) bool {
	found := false
	entriesOnce.Do(func() {
		entries = Compiled()
	})
	for _, arg := range args {
		if matched := search(arg, entries); matched {
			found = true
		}
	}
	return found
}

func search(arg string, entries []Entry) bool {
	var found bool = false
	for _, entry := range entries {
		if !entry.Compiled.MatchString(arg) {
			continue
		}
		if !flags.Extended && !isSimpleHash(entry.Name) {
			continue
		}

		printModeByFlags(utils.Modes{
			Name:    entry.Name,
			Hashcat: entry.HashCat,
			John:    entry.John,
		})

		found = true

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
