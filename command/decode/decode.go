package decode

import (
	"ghid/flags"

	"github.com/spf13/cobra"
)

func Commands() []*cobra.Command {
	var decodeCmd = &cobra.Command{
		Use:   "decode",
		Short: "decode file.txt",
		Long:  "decode file.txt",
		Run: func(cmd *cobra.Command, args []string) {
			Decode(&DecodeData{
				OpenFile:   flags.ReadFile,
				WriterFile: flags.WriterFile,
				NameHash:   flags.NameHash,
				Dictionary: flags.Dictionary,
			})
		},
	}
	flags.AddBoolFlags(decodeCmd)
	flags.AddStringFlags(decodeCmd)

	return []*cobra.Command{decodeCmd}
}
