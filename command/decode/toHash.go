package decode

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha3"
	"crypto/sha512"
	"encoding/hex"
	"ghid/errHandler"
	"ghid/output"

	"golang.org/x/crypto/blake2b"
	"golang.org/x/crypto/blake2s"
	"golang.org/x/crypto/md4"
	"golang.org/x/crypto/ripemd160"
)

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
