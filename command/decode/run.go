package decode

import (
	"context"
	"fmt"
	"ghid/errHandler"
	"ghid/output"
	"ghid/utils"
	"runtime"
	"strings"
	"sync"
)

// __________________________________________________
type DecodeData struct {
	OpenFile   string
	WriterFile string
	NameHash   string
	Dictionary string
}

func Decode(decodeData *DecodeData) {
	var (
		out  strings.Builder
		dict []string = utils.ParseTxt(decodeData.Dictionary)
	)

	for _, value := range utils.ParseTxt(decodeData.OpenFile) {
		var nameUser, passUser, passHash string
		parts := strings.SplitN(value, ":", 2)

		nameUser = "unknown"
		passHash = strings.TrimSpace(parts[0])
		if len(parts) == 2 {
			nameUser = parts[0]
			passHash = parts[1]
		}

		passUser = runDecode(passHash, decodeData.NameHash, dict)

		out.WriteString(nameUser)
		out.WriteString(":")
		out.WriteString(passUser)
		out.WriteString("\n")
	}

	utils.CreateTxt(decodeData.WriterFile, out.String())
}

func runDecode(passUser, nameHash string, dictionary []string) string {

	var out strings.Builder

	if len(dictionary) == 0 {
		out.WriteString(errHandler.ErrDictionaryEmpty.Error())
		out.WriteString(" : [empty dictionary file]")
		output.PrintWarning(out.String())

		out.Reset()

		out.WriteString(passUser)
		out.WriteString(" [dictionary empty]")
		return out.String()
	}

	hashType, ok := HashFromString(nameHash)
	if !ok {
		out.WriteString(passUser)
		out.WriteString(" [unknown hash type]")
		return out.String()
	}

	expectedLen := int(digestSizes[hashType]) * 2
	if expectedLen > 0 && expectedLen != len(passUser) {
		out.WriteString(passUser)
		out.WriteString(" [invalid length for hash type: ")
		out.WriteString(nameHash)
		out.WriteString(" (expected ")
		out.WriteString(fmt.Sprintf("%d", expectedLen))
		out.WriteString(")]")
		return out.String()
	}

	conText, cancel := context.WithCancel(context.Background())
	defer cancel()

	wordChan := make(chan string)
	resultChan := make(chan string, 1)

	var waitGroup sync.WaitGroup

	//_______________________________________________________________

	var numWorker = runtime.NumCPU()
	for i := 0; i < numWorker; i++ {
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
					if toHash(word, hashType) == passUser {
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
		for _, word := range dictionary {
			word = strings.TrimSpace(word)
			if word == "" {
				continue
			}
			select {
			case <-conText.Done():
				return
			case wordChan <- word:
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
