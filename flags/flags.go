package flags

import "github.com/spf13/cobra"

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

func AddBoolFlags(cmd *cobra.Command, flags []BoolFlagsStruct) {
	for _, flag := range flags {
		if flag.Shorthand != "" {
			cmd.Flags().BoolVarP(flag.Target, flag.Name, flag.Shorthand, flag.Value, flag.Usage)
		} else {
			cmd.Flags().BoolVar(flag.Target, flag.Name, flag.Value, flag.Usage)
		}
	}
}
