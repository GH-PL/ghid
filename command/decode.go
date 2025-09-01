package command

import (
	"context"
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
	"fmt"
	"ghid/errHandler"
	"ghid/flags"
	"ghid/output"
	"ghid/utils"
	"runtime"
	"strings"
	"sync"

	"github.com/spf13/cobra"
	"golang.org/x/crypto/blake2b"
	"golang.org/x/crypto/blake2s"
	"golang.org/x/crypto/md4"
	"golang.org/x/crypto/ripemd160"
	"golang.org/x/crypto/sha3"
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

func toHash(word string, nameHash Hash) string {

	switch nameHash {
	case MD4:
		hash := md4.New()
		hash.Write([]byte(word))
		return hex.EncodeToString(hash.Sum(nil))
	case MD5:
		sum := md5.Sum([]byte(word))
		return hex.EncodeToString(sum[:])
	case SHA1:
		sum := sha1.Sum([]byte(word))
		return hex.EncodeToString(sum[:])
	case SHA224:
		sum := sha256.Sum224([]byte(word))
		return hex.EncodeToString(sum[:])
	case SHA256:
		sum := sha256.Sum256([]byte(word))
		return hex.EncodeToString(sum[:])
	case SHA3:
		sum := sha3.Sum256([]byte(word))
		return hex.EncodeToString(sum[:])
	case SHA384:
		sum := sha3.Sum384([]byte(word))
		return hex.EncodeToString(sum[:])
	case SHA512:
		sum := sha512.Sum512([]byte(word))
		return hex.EncodeToString(sum[:])
	case MD5SHA1:
		md5Sum := md5.Sum([]byte(word))
		sha1Sum := sha1.Sum([]byte(word))
		combined := append(md5Sum[:], sha1Sum[:]...)
		return hex.EncodeToString(combined)
	case RIPEMD160:
		hash := ripemd160.New()
		hash.Write([]byte(word))
		return hex.EncodeToString(hash.Sum(nil))
	case SHA3_224:
		sum := sha3.Sum224([]byte(word))
		return hex.EncodeToString(sum[:])
	case SHA3_256:
		sum := sha3.Sum256([]byte(word))
		return hex.EncodeToString(sum[:])
	case SHA3_384:
		sum := sha3.Sum384([]byte(word))
		return hex.EncodeToString(sum[:])
	case SHA3_512:
		sum := sha3.Sum512([]byte(word))
		return hex.EncodeToString(sum[:])
	case SHA512_224:
		sum := sha512.Sum512_224([]byte(word))
		return hex.EncodeToString(sum[:])
	case SHA512_256:
		sum := sha512.Sum512_256([]byte(word))
		return hex.EncodeToString(sum[:])
	case BLAKE2s_256:
		sum := blake2s.Sum256([]byte(word))
		return hex.EncodeToString(sum[:])
	case BLAKE2b_256:
		sum := blake2b.Sum256([]byte(word))
		return hex.EncodeToString(sum[:])
	case BLAKE2b_512:
		sum := blake2b.Sum512([]byte(word))
		return hex.EncodeToString(sum[:])

	default:
		output.PrintError(errHandler.ErrNotTypeHash)
		return ""
	}
}

// Go
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

type Hash uint8

const (
	MD4         Hash = 1 + iota // import golang.org/x/crypto/md4
	MD5                         // import crypto/md5
	SHA1                        // import crypto/sha1
	SHA3                        // import golang.org/x/crypto/sha3
	SHA224                      // import crypto/sha256
	SHA256                      // import crypto/sha256
	SHA384                      // import crypto/sha512
	SHA512                      // import crypto/sha512
	MD5SHA1                     // no implementation; MD5+SHA1 used for TLS RSA
	RIPEMD160                   // import golang.org/x/crypto/ripemd160
	SHA3_224                    // import golang.org/x/crypto/sha3
	SHA3_256                    // import golang.org/x/crypto/sha3
	SHA3_384                    // import golang.org/x/crypto/sha3
	SHA3_512                    // import golang.org/x/crypto/sha3
	SHA512_224                  // import crypto/sha512
	SHA512_256                  // import crypto/sha512
	BLAKE2s_256                 // import golang.org/x/crypto/blake2s
	BLAKE2b_256                 // import golang.org/x/crypto/blake2b
	BLAKE2b_384                 // import golang.org/x/crypto/blake2b
	BLAKE2b_512                 // import golang.org/x/crypto/blake2b
	maxHash
)

var digestSizes = map[Hash]uint8{
	MD4:         16,
	MD5:         16,
	SHA1:        20,
	SHA224:      28,
	SHA256:      32,
	SHA384:      48,
	SHA512:      64,
	SHA512_224:  28,
	SHA512_256:  32,
	SHA3_224:    28,
	SHA3_256:    32,
	SHA3_384:    48,
	SHA3_512:    64,
	MD5SHA1:     36,
	RIPEMD160:   20,
	BLAKE2s_256: 32,
	BLAKE2b_256: 32,
	BLAKE2b_384: 48,
	BLAKE2b_512: 64,
}

func (h Hash) String() string {
	switch h {
	case MD4:
		return "MD4"
	case MD5:
		return "MD5"
	case SHA1:
		return "SHA1"
	case SHA224:
		return "SHA224"
	case SHA256:
		return "SHA256"
	case SHA384:
		return "SHA384"
	case SHA512:
		return "SHA512"
	case MD5SHA1:
		return "MD5SHA1"
	case RIPEMD160:
		return "RIPEMD160"
	case SHA3_224:
		return "SHA3_224"
	case SHA3_256:
		return "SHA3_256"
	case SHA3_384:
		return "SHA3_384"
	case SHA3_512:
		return "SHA3_512"
	case SHA512_224:
		return "SHA512_224"
	case SHA512_256:
		return "SHA512_256"
	case BLAKE2s_256:
		return "BLAKE2s_256"
	case BLAKE2b_256:
		return "BLAKE2b_256"
	case BLAKE2b_384:
		return "BLAKE2b_384"
	case BLAKE2b_512:
		return "BLAKE2b_512"
	default:
		return "UNKNOWN"
	}
}

func HashFromString(nameHash string) (Hash, bool) {
	nameHash = strings.ToUpper(nameHash)
	for hash := Hash(1); hash < maxHash; hash++ {
		if hash.String() == nameHash {
			return hash, true
		}
	}
	return 0, false
}
