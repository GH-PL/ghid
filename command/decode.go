package command

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"encoding/hex"
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
			OpenFile:   flags.ReadFile,
			WriterFile: flags.WriterFile,
			NameHash:   flags.NameHash,
			Dictionary: flags.Dictionary,
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
	NameHash   string
	Dictionary string
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
			passUser = runDecode(parts[1], decodeData.NameHash, decodeData.Dictionary)
		} else {
			nameUser = "unknown"
			passUser = runDecode(parts[0], decodeData.NameHash, decodeData.Dictionary)
		}

		res := nameUser + ":" + passUser
		result = append(result, res)
	}

	utils.CreateTxt(decodeData.WriterFile, strings.Join(result, "\n"))
}

func runDecode(passUser, nameHash, dictionary string) string {
	for _, word := range utils.ParseTxt(dictionary) {
		word = strings.TrimSpace(word)
		if word == "" {
			continue
		}
		if toHash(word, nameHash) == passUser {
			return word
		}
	}
	return passUser + " [not found]"
}

func toHash(word string, nameHash string) string {
	switch nameHash {
	case "md5":
		sum := md5.Sum([]byte(word))
		return hex.EncodeToString(sum[:])
	case "sha1":
		sum := sha1.Sum([]byte(word))
		return hex.EncodeToString(sum[:])
	case "sha256":
		sum := sha256.Sum256([]byte(word))
		return hex.EncodeToString(sum[:])
	default:
		return ""
	}
}
