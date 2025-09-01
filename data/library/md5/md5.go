package md5

import "ghid/data/library"

const (
	Size        = 16
	BlockSize   = 64
	maxAsmIters = 1024
	maxAsmSize  = BlockSize * maxAsmIters
)

const (
	init0 = 0x67452301
	init1 = 0xEFCDAB89
	init2 = 0x98BADCFE
	init3 = 0x10325476
)

type digest struct {
	s   [4]uint32
	x   [BlockSize]byte
	nx  int
	len uint64
}

func init() {
	library.RegisterHash(library.MD5, New)
}
func (d *digest) Size() int { return Size }

func (d *digest) BlockSize() int { return BlockSize }

func (d *digest) Sum(b []byte) []byte {
	return []byte{}
}

func New() library.Hasher {
	d := new(digest)
	d.Reset()
	return d
}

func (d *digest) Write(p []byte) (int, error) { return 0, nil }

func (d *digest) checkSum() [Size]byte {
	var digest [Size]byte
	return digest
}

func (d *digest) Reset() {}

func Sum(input []byte) [Size]byte {
	var d digest
	d.Reset()
	d.Write(input)
	return d.checkSum()
}
