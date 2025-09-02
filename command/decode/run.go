package decode

import (
	"context"
	"fmt"
	"ghid/data"
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
	Core       int
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

		passUser = runDecode(passHash, decodeData.NameHash, dict, decodeData.Core)

		out.WriteString(nameUser)
		out.WriteString(":")
		out.WriteString(passUser)
		out.WriteString("\n")
	}

	utils.CreateTxt(decodeData.WriterFile, out.String())
}

func runDecode(passHash, nameHash string, dictionary []string, core int) string {

	var out strings.Builder

	if len(dictionary) == 0 {
		out.WriteString(errHandler.ErrDictionaryEmpty.Error())
		out.WriteString(" : [empty dictionary file]")
		output.PrintWarning(out.String())

		out.Reset()

		fmt.Fprintf(&out, "%s [dictionary empty]", passHash)
		return out.String()
	}

	hashType, ok := HashFromString(nameHash)
	if !ok {
		fmt.Fprintf(&out, "%s [unknown hash type]", passHash)
		return out.String()
	}

	expectedLen := int(digestSizes[hashType]) * 2
	if expectedLen > 0 && expectedLen != len(passHash) {
		fmt.Println(passHash, " ", len(passHash))

		fmt.Fprintf(&out, "%s [invalid length for hash type: %s (expected %d)]",
			passHash, nameHash, expectedLen)
		return out.String()
	}

	conText, cancel := context.WithCancel(context.Background())
	defer cancel()

	wordChan := make(chan string, 1000)
	resultChan := make(chan string, 1)

	var waitGroup sync.WaitGroup

	//_______________________________________________________________

	var numWorker = runtime.NumCPU() / data.NUM_WORKER
	if core != 2 {
		numWorker = core
	}
	if numWorker < 1 {
		numWorker = 1
	}
	if numWorker > runtime.NumCPU() {
		numWorker = runtime.NumCPU()
	}

	cleanDict := make([]string, 0, len(dictionary))
	for _, word := range dictionary {
		word = strings.TrimSpace(word)
		if word == "" {
			continue
		}
		cleanDict = append(cleanDict, word)
	}

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
					if toHash(word, hashType) == passHash {
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
		for _, word := range cleanDict {
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
	return passHash + " [not found]"
}
