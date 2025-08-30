package command

import (
	"ghid/flags"
	"ghid/utils"
	"strings"

	"github.com/spf13/cobra"
)

var DecodeCmd = &cobra.Command{
	Use:   "decode",
	Short: "decode file.txt",
	Long:  "decode file.txt",
	Run: func(cmd *cobra.Command, args []string) {
		decode(&DecodeData{
			OpenFile:   flags.OpenFile,
			WriterFile: flags.WriterFile,
		})
	},
}

func init() {
	flags.AddBoolFlags(DecodeCmd)
	flags.AddStringFlags(DecodeCmd)
}

// __________________________________________________
type DecodeData struct {
	OpenFile   string
	WriterFile string
}

func decode(decodeData *DecodeData) {
	var result []string
	for _, value := range utils.ParseTxt(decodeData.OpenFile) {
		var nameUser, passUser string
		parts := strings.SplitN(value, ":", 2)
		if len(parts) > 2 || len(parts) == 0 {
			continue
		}
		if len(parts) == 2 {
			nameUser = parts[0]
			passUser = parts[1]
		} else {
			nameUser = "unknown"
			passUser = parts[0]
		}

		res := nameUser + ":" + passUser
		result = append(result, res)
	}

	utils.CreateTxt(decodeData.WriterFile, strings.Join(result, "\n"))
}
