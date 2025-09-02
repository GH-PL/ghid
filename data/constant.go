package data

// String
const (
	VERSION              string = "GHID version 0.0.1"
	WAY_DATA_JSON        string = "data/data.json"
	WAY_POPULAR_HASH_CSV string = "data/popularHash.csv"
	DEFAULT_DECRYPT_FILE string = "decrypt.txt"
)

type Command uint

const (
	CMD_DETECT Command = iota + 1
	CMD_DECODE
	CMD_LIST
	CMD_SAMPLES
	CMD_VERSION
)

const (
	NUM_WORKER int = 2 // Defines the number of worker threads. Actual core usage: total cores / NUM_WORKER
)
