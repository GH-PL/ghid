package decode

import (
	"ghid/flags"

	"github.com/spf13/cobra"
)

func init() {
	flags.AddBoolFlags(DecodeCmd)
	flags.AddStringFlags(DecodeCmd)
}

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
