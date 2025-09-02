package data

// String
const (
	VERSION              string = "GHID version 0.0.1"
	WAY_DATA_JSON        string = "data/data.json"
	WAY_POPULAR_HASH_CSV string = "data/popularHash.csv"
	DEFAULT_DECRYPT_FILE string = "decrypt.txt"
)

// Int
const ()

/*
// _________Name command_________
const (
	CMD_DETECT  string = "detect"
	CMD_DECODE  string = "decode"
	CMD_LIST    string = "list"
	CMD_SAMPLES string = "samples"
	CMD_VERSION string = "version"
)
*/
type Command uint

const (
	CMD_DETECT Command = iota + 1
	CMD_DECODE
	CMD_LIST
	CMD_SAMPLES
	CMD_VERSION
)
