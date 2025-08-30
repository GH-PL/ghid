package command

import (
	"fmt"
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
		parts := strings.SplitN(value, ":", 2)
		if len(parts) != 2 {
			fmt.Println("Invalid string format")
			continue
		}
		var (
			nameUser = parts[0]
			passUser = parts[1]
		)
		res := nameUser + ":" + passUser
		result = append(result, res)
	}
	fmt.Println(result)
	fmt.Println(decodeData.WriterFile)
}
