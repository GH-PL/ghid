package utils

import (
	"bufio"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"ghid/data"
	"ghid/errHandler"
	"ghid/flags"
	"ghid/output"
	"os"
	"path/filepath"
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
func ParseJson(filePath string) []data.Hash {
	var hashes []data.Hash
	loadJson(filePath, &hashes)
	return hashes
}

// Load a CSV file and return all records as a 2D slice of strings
func loadCsv(filePath string) [][]string {
	file, err := openFile(filePath)

	if err != nil {
		errHandler.ErrorFile("open file", filePath, err)
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
func ParseCsv(filePath string) map[string]struct{} {
	var (
		PopularHashesSet = make(map[string]struct{})
		records          = loadCsv(filePath)
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

func ParseTxt(filePath string) []string {
	// Open File
	file, err := openFile(filePath)

	if err != nil {
		errHandler.ErrorFile("open file", filePath, err)
	}

	defer file.Close()

	// Read file
	var lines []string
	reader := bufio.NewScanner(file)
	for reader.Scan() {
		lines = append(lines, reader.Text())
	}

	if err := reader.Err(); err != nil {
		errHandler.ErrorFile("read txt lines", filePath, err)
	}
	return lines
}

func CreateTxt(nameFile string, decrypt string) {
	filePath := CreateDir(nameFile)

	if nameFile == "" {
		output.PrintWarning("No file path provided — creating decrypt.txt")
		nameFile = data.DEFAULT_DECRYPT_FILE
	}

	file, err := os.Create(filePath)

	if err != nil {
		errHandler.ErrorFile("create file", nameFile, err)
	}

	defer file.Close()

	_, err = file.WriteString(decrypt)

	if err != nil {
		errHandler.ErrorFile("write string in ", nameFile, err)
	}

}

func CreateDir(namefile string) string {
	homeDir, err := os.UserHomeDir()

	if err != nil {
		errHandler.ErrorFile("perform a home directory search ", "os.UserHomeDir", err)
	}

	filePath := filepath.Join(homeDir, ".ghid", namefile)

	err = os.MkdirAll(filepath.Dir(filePath), 0755)
	if err != nil {
		errHandler.ErrorFile("create app dir ", "homeDir/.ghid", err)
	}
	return filePath
}
