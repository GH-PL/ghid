package utils

import (
	"encoding/json"
	"log"
	"os"
	"strings"

	"github.com/GH-PL/ghid/data"
)

func loadJson(filePath string, v interface{}) {
	file, errFile := os.Open(filePath)
	if errFile != nil {
		log.Fatalf("Error opening %s: %v", filePath, errFile)
	}
	defer file.Close()

	if errDecode := json.NewDecoder(file).Decode(v); errDecode != nil {
		log.Fatalf("JSON decode error in %s: %v", filePath, errDecode)
	}
}

func ParseJson() []data.Hash {
	var hashes []data.Hash
	loadJson("data/data.json", &hashes)
	return hashes
}

var PopularHashesSet map[string]struct{}

func LoadPopularHashes() {
	var popularList []string
	loadJson("data/popularHash.json", &popularList)

	PopularHashesSet = make(map[string]struct{})
	for _, name := range popularList {
		name = strings.ToLower(name)
		PopularHashesSet[name] = struct{}{}
	}
}
