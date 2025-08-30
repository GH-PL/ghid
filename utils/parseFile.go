package utils

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"ghid/data"
	"ghid/errHandler"
	"ghid/flags"
	"ghid/output"
	"os"
	"strings"
)

// Open a file given its file path
func openFile(filePath string) (*os.File, error) {
	return os.OpenFile(filePath, os.O_RDONLY, 0)
}

// Load a JSON file and decode it into the provided interface
func loadJson(filePath string, v interface{}) {
	file, err := openFile(filePath)

	if err != nil {
		errHandler.ErrorFile("open file", filePath, err)
	}
	defer file.Close()

	if errDecode := json.NewDecoder(file).Decode(v); errDecode != nil {
		errHandler.ErrorFile("decode", filePath, errDecode)
	}
}

// Parse the JSON file and return a slice of data.Hash
func ParseJson() []data.Hash {
	var hashes []data.Hash
	loadJson(data.WAY_DATA_JSON, &hashes)
	return hashes
}

// Load a CSV file and return all records as a 2D slice of strings
func loadCsv(filePath string) [][]string {
	file, err := openFile(filePath)

	if err != nil {
		errHandler.ErrorFile("open file", data.WAY_POPULAR_HASH_CSV, err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()

	if err != nil {
		errHandler.ErrorFile("reader records", data.WAY_POPULAR_HASH_CSV, err)
	}

	return records
}

// Parse the CSV file and return a map of hash names to struct{}
func ParceCsv() map[string]struct{} {
	var (
		PopularHashesSet = make(map[string]struct{})
		records          = loadCsv(data.WAY_POPULAR_HASH_CSV)
	)

	// If the CSV file is empty, enable the extended mode and return it.
	if records == nil {
		output.PrintWarning(fmt.Sprintf("File %s is empty. It will include the extended mode.", data.WAY_POPULAR_HASH_CSV))
		flags.Extended = true
	}

	// Loop through each record and add the hash name to the map (converted to lowercase)
	for _, record := range records {
		for _, name := range record {
			name = strings.ToLower(name)
			PopularHashesSet[name] = struct{}{}
		}
	}
	return PopularHashesSet
}
