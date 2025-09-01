package cmd

import (
	"os"

	"ghid/command/decode"
	"ghid/command/detect"
	"ghid/command/list"
	"ghid/command/samples"
	"ghid/command/version"
	"ghid/flags"
	"ghid/output"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(samples.Commands()...)
	rootCmd.AddCommand(list.Commands()...)
	rootCmd.AddCommand(version.Commands()...)
	rootCmd.AddCommand(decode.Commands()...)
	rootCmd.AddCommand(detect.Commands()...)
}

var rootCmd = &cobra.Command{
	Use:   "ghid",
	Short: "Ghid — Golang Hash Identifier — is a sample Go application that serves as an analog to the Haiti, HashId program and others.",
	Long: `Ghid is a CLI tool for identifying hash types and working with hash data.

Examples:
  ghid detect 5f4dcc3b5aa765d61d8327deb882cf99   # Detect hash type
  ghid decode -f hashes.txt -o output.txt        # Decode hash file
  ghid list                                      # List all supported hash types
  ghid samples md5                               # Show sample hashes for MD5
  ghid version                                   # Display app version`,
	Args: cobra.ArbitraryArgs,
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		if flags.NoColorFlag {
			output.DisableColorOutput()
		}
	},
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
