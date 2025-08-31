package command

import (
	"context"
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"encoding/hex"
	"ghid/flags"
	"ghid/utils"
	"strings"
	"sync"

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
			passUser = runDecore(parts[1], decodeData.NameHash, decodeData.Dictionary)
		} else {
			nameUser = "unknown"
			passUser = runDecore(parts[0], decodeData.NameHash, decodeData.Dictionary)
		}

		res := nameUser + ":" + passUser
		result = append(result, res)
	}

	utils.CreateTxt(decodeData.WriterFile, strings.Join(result, "\n"))
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

// Go
func runDecore(passUser, nameHash, dictionary string) string {
	words := utils.ParseTxt(dictionary)
	numWorkes := 16
	wordChan := make(chan string)
	resultChan := make(chan string, 1)
	conText, cancel := context.WithCancel(context.Background())
	defer cancel()

	var waitGroup sync.WaitGroup

	for i := 0; i < numWorkes; i++ {
		waitGroup.Add(1)
		go func() {
			defer waitGroup.Done()
			for {
				select {
				case <-conText.Done():
					return
				case word, ok := <-wordChan:
					if !ok {
						return
					}
					word = strings.TrimSpace(word)
					if word == "" {
						continue
					}
					if toHash(word, nameHash) == passUser {
						select {
						case resultChan <- word:
							cancel()
						default:
						}
						return
					}
				}
			}
		}()
	}

	go func() {
		defer close(wordChan)
		for _, word := range words {
			select {
			case <-conText.Done():
				return
			default:
				wordChan <- word
			}
		}
	}()

	go func() {
		waitGroup.Wait()
		close(resultChan)
	}()

	if result, ok := <-resultChan; ok {
		return result
	}
	return passUser + " [not found]"
}
