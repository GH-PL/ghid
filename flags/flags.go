package flags

import (
	"ghid/data"

	"github.com/spf13/cobra"
)

// _________Bool flags___________
var (
	ShortFlag   bool
	NoColorFlag bool
	Hashcat     bool
	John        bool
	Extended    bool
)

type BoolFlagsStruct struct {
	Name      string
	Shorthand string
	Value     bool
	Usage     string
	Target    *bool
}

var BoolFlags = []BoolFlagsStruct{
	{
		Name:      "short",
		Shorthand: "s",
		Value:     false,
		Usage:     "Short message",
		Target:    &ShortFlag,
	},
	{
		Name:      "no-color",
		Shorthand: "n",
		Value:     false,
		Usage:     "Disable color output",
		Target:    &NoColorFlag,
	},
	{
		Name:      "hashcat-only",
		Shorthand: "c",
		Value:     false,
		Usage:     "Show only hashcat references",
		Target:    &Hashcat,
	},
	{
		Name:      "john-only",
		Shorthand: "j",
		Value:     false,
		Usage:     "Show only john the ripper references",
		Target:    &John,
	},
	{
		Name:      "extended",
		Shorthand: "e",
		Value:     false,
		Usage:     "List all possible hash algorithms including ones using salt",
		Target:    &Extended,
	},
}

// ________________String flags___________________
var (
	ReadFile   string
	WriterFile string
	NameHash   string
	Dictionary string
)

type StringFlagsStruct struct {
	Name      string
	Shorthand string
	Value     string
	Usage     string
	Target    *string
}

var StringFlags = []StringFlagsStruct{
	{
		Name:      "read",
		Shorthand: "r",
		Value:     "",
		Usage:     "Read file",
		Target:    &ReadFile,
	},
	{
		Name:      "writer",
		Shorthand: "w",
		Value:     data.DEFAULT_DECRYPT_FILE,
		Usage:     "Writer file",
		Target:    &WriterFile,
	},
	{
		Name:      "hash-type",
		Shorthand: "t",
		Value:     "md5",
		Usage:     "Type hash",
		Target:    &NameHash,
	},
	{
		Name:      "dictionary",
		Shorthand: "d",
		Value:     "",
		Usage:     "Dictionary",
		Target:    &Dictionary,
	},
}

// _______Int flags___________
// Int
var (
	NumWorker int
)

type IntFlagsStruct struct {
	Name      string
	Shorthand string
	Value     int
	Usage     string
	Target    *int
}

var IntFlags = []IntFlagsStruct{
	{
		Name:      "limit",
		Shorthand: "l",
		Value:     data.NUM_WORKER,
		Usage:     "Number Worker, default: Core/2",
		Target:    &NumWorker,
	},
}

// _______________Map [command] [Flags]__________________________
var FlagsPerCommand = map[data.Command][]string{
	data.CMD_DECODE:  {"short", "read", "writer", "hash-type", "dictionary", "limit"},
	data.CMD_DETECT:  {"short", "extended", "hashcat-only", "john-only", "no-color"}, // "read", "writer"
	data.CMD_LIST:    {"no-color"},                                                   // "hashcat-only", "john-only"
	data.CMD_SAMPLES: {"no-color"},                                                   // "hashcat-only", "john-only"
	data.CMD_VERSION: {"no-color"},
}

// ____________________Func Add Flags to Command______________
func AddCommandFlags(cmd *cobra.Command, commandName data.Command) {
	// Add Bool-flags.
	for _, flagName := range FlagsPerCommand[commandName] {
		for _, flag := range BoolFlags {
			if flag.Name == flagName {
				if flag.Shorthand != "" {
					cmd.Flags().BoolVarP(flag.Target, flag.Name, flag.Shorthand, flag.Value, flag.Usage)
				} else {
					cmd.Flags().BoolVar(flag.Target, flag.Name, flag.Value, flag.Usage)
				}
			}
		}
		// Add String-flags.
		for _, flag := range StringFlags {
			if flag.Name == flagName {
				if flag.Shorthand != "" {
					cmd.Flags().StringVarP(flag.Target, flag.Name, flag.Shorthand, flag.Value, flag.Usage)
				} else {
					cmd.Flags().StringVar(flag.Target, flag.Name, flag.Value, flag.Usage)
				}
			}
		}
		// Add Int-flags.
		for _, flag := range IntFlags {
			if flag.Name == flagName {
				if flag.Shorthand != "" {
					cmd.Flags().IntVarP(flag.Target, flag.Name, flag.Shorthand, flag.Value, flag.Usage)
				} else {
					cmd.Flags().IntVar(flag.Target, flag.Name, flag.Value, flag.Usage)
				}
			}
		}
	}
}
