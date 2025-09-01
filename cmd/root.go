package cmd

import (
	"os"

	"ghid/command/decode"
	"ghid/errHandler"
	"ghid/flags"
	"ghid/output"

	"ghid/command"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(command.ListCmd)
	rootCmd.AddCommand(command.SamplesCmd)
	rootCmd.AddCommand(command.VersionCmd)
	rootCmd.AddCommand(decode.Commands()...)
}

var rootCmd = RootCmd()

func RootCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "ghid [command] [flags] <hash>",
		Short: "Ghid — Golang Hash Identifier — is a sample Go application that serves as an analog to the Haiti, HashId program and others.",
		Long:  "Ghid — Golang Hash Identifier — is a sample Go application that serves as an analog to the Haiti, HashId program and others.",
		Args:  cobra.ArbitraryArgs,
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) < 1 {
				cmd.Help()
				return
			}
			matchHash := matchHashTypes(args)
			if !matchHash {
				output.PrintError(errHandler.ErrNotFoundHash)
			} else {
				if !flags.Extended {
					output.PrintWarning("You need extended mode")
				}
			}

		},
		PersistentPreRun: func(cmd *cobra.Command, args []string) {

			if flags.NoColorFlag {
				output.DisableColorOutput()
			}
		},
	}
	flags.AddBoolFlags(cmd)
	return cmd
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
