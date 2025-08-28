package utils

import (
	"encoding/json"
	"ghid/data"
	"ghid/errHandler"
	"os"
	"strings"
)

func loadJson(filePath string, v interface{}) {

	file, err := os.OpenFile(filePath, os.O_RDONLY, 0)
	if err != nil {
		errHandler.ErrorFile("open file", filePath, err)
	}
	defer file.Close()

	if errDecode := json.NewDecoder(file).Decode(v); errDecode != nil {
		errHandler.ErrorFile("decode", filePath, errDecode)
	}
}

func ParseJson() []data.Hash {
	var hashes []data.Hash
	loadJson(data.WAY_DATA_JSON, &hashes)
	return hashes
}

var PopularHashesSet map[string]struct{}

func LoadPopularHashes() {
	var popularList []string
	loadJson(data.WAY_POPULAR_HASH_JSON, &popularList)

	PopularHashesSet = make(map[string]struct{})
	for _, name := range popularList {
		name = strings.ToLower(name)
		PopularHashesSet[name] = struct{}{}
	}
}
