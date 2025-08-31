package flags

import (
	"ghid/data"

	"github.com/spf13/cobra"
)

// _________Bool flags___________
var (
	VersionFlag bool
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
		Name:      "version",
		Shorthand: "v",
		Value:     false,
		Usage:     "Show version",
		Target:    &VersionFlag,
	},
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

func AddBoolFlags(cmd *cobra.Command) {
	for _, flag := range BoolFlags {
		if flag.Shorthand != "" {
			cmd.Flags().BoolVarP(flag.Target, flag.Name, flag.Shorthand, flag.Value, flag.Usage)
		} else {
			cmd.Flags().BoolVar(flag.Target, flag.Name, flag.Value, flag.Usage)
		}
	}
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

func AddStringFlags(cmd *cobra.Command) {
	for _, flag := range StringFlags {
		if flag.Shorthand != "" {
			cmd.Flags().StringVarP(flag.Target, flag.Name, flag.Shorthand, flag.Value, flag.Usage)
		} else {
			cmd.Flags().StringVar(flag.Target, flag.Name, flag.Value, flag.Usage)
		}
	}
}
