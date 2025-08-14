package cmd

import (
	"HaitiGo/flags"
	"fmt"
	"os"

	"github.com/GH-PL/ghid/command"
	"github.com/GH-PL/ghid/data"

	"github.com/GH-PL/ghid/utils"

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
				if !showHashValue(args) {
					utils.PrintColorText(&utils.Text{
						Text:           "Not found type for this Hash",
						ColorAttribute: color.FgRed,
						Style:          []color.Attribute{color.Bold},
					})
				}
			} else {
				cmd.Help()
				return
			}
		},
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			if flags.VersionFlag {
				fmt.Println("Version: ", data.VERSION)
				os.Exit(0)
			}
			if flags.NoColorFlag {
				utils.DisableColorOutput()
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
