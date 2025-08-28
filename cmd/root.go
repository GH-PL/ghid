package cmd

import (
	"fmt"
	"os"

	"ghid/flags"
	"ghid/output"

	"ghid/command"
	"ghid/data"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(command.ListCmd)
	rootCmd.AddCommand(command.SamplesCmd)
}

var rootCmd = RootCmd()

func RootCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "ghid [options] <hash>",
		Short: "Ghid — Golang Hash Identifier — is a sample Go application that serves as an analog to the Haiti, HashId program and others.",
		Long:  "Ghid — Golang Hash Identifier — is a sample Go application that serves as an analog to the Haiti, HashId program and others.",
		Args:  cobra.ArbitraryArgs,
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) != 0 {
				showResult := showHashValue(args)
				if !showResult {
					output.PrintColorText(&output.Text{
						Text:           "Not found type for this Hash",
						ColorAttribute: color.FgRed,
						Style:          []color.Attribute{color.Bold},
					})
				} else if !flags.Extended && showResult {
					output.PrintColorText(&output.Text{
						Text:           "You need extended mode",
						ColorAttribute: color.BgYellow,
						Style:          []color.Attribute{color.Bold},
					})
				}

			} else {
				cmd.Help()
				return
			}
		},
		PreRun: func(cmd *cobra.Command, args []string) {
			if flags.VersionFlag {
				fmt.Println(data.VERSION)
				os.Exit(0)
			}
		},
		PersistentPreRun: func(cmd *cobra.Command, args []string) {

			if flags.NoColorFlag {
				output.DisableColorOutput()
			}
		},
	}
	flags.AddBoolFlags(cmd, flags.BoolFlags)
	return cmd
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Printf("root command: %s\n", err)
		os.Exit(1)
	}
}
