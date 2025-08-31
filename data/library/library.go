package main

import (
	"hash"
	"strconv"
)

type Hash uint

// uint id, iota = 0, 1...19
const (
	MD4         Hash = 1 + iota // import golang.org/x/crypto/md4
	MD5                         // import crypto/md5
	SHA1                        // import crypto/sha1
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

var digestSizes = []uint8{
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

func (hashId Hash) HashFunc() Hash {
	return hashId
}

// return string value Hash.
func (hashId Hash) String() string {
	switch hashId {
	case MD4:
		return "MD4"
	case MD5:
		return "MD5"
	case SHA1:
		return "SHA-1"
	case SHA224:
		return "SHA-224"
	case SHA256:
		return "SHA-256"
	case SHA384:
		return "SHA-384"
	case SHA512:
		return "SHA-512"
	case MD5SHA1:
		return "MD5+SHA1"
	case RIPEMD160:
		return "RIPEMD-160"
	case SHA3_224:
		return "SHA3-224"
	case SHA3_256:
		return "SHA3-256"
	case SHA3_384:
		return "SHA3-384"
	case SHA3_512:
		return "SHA3-512"
	case SHA512_224:
		return "SHA-512/224"
	case SHA512_256:
		return "SHA-512/256"
	case BLAKE2s_256:
		return "BLAKE2s-256"
	case BLAKE2b_256:
		return "BLAKE2b-256"
	case BLAKE2b_384:
		return "BLAKE2b-384"
	case BLAKE2b_512:
		return "BLAKE2b-512"
	default:
		return "unknown hash value " + strconv.Itoa(int(hashId))
	}
}

// from uint id make func
var hashes = make([]func() hash.Hash, maxHash)

// return func for input uint id.
func (hashId Hash) New() hash.Hash {
	if hashId > 0 && hashId < maxHash {
		f := hashes[hashId]
		if f != nil {
			return f()
		}
	}
	panic("crypto: requested hash function #" + strconv.Itoa(int(hashId)) + " is unavailable")
}
