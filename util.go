package pwntools

import (
	"encoding/base64"
	"encoding/binary"
	"encoding/hex"
)

type number interface {
	int | string
}

func UnHex(h string) []byte {
	b, err := hex.DecodeString(h)

	if err != nil {
		Error(err.Error())
	}

	return b
}

func Hex(b []byte) string {
	return hex.EncodeToString(b)
}

func Xor(a []byte, b ...[]byte) []byte {
	length := len(a)

	for j := 0; j < len(b); j++ {
		if len(b[j]) > length {
			length = len(b[j])
		}
	}

	r := make([]byte, length)

	for i := 0; i < length; i++ {
		rr := a[i%len(a)]

		for j := 0; j < len(b); j++ {
			rr ^= b[j][i%len(b[j])]
		}

		r = append(r, rr)
	}

	return r
}

func B64d(e string) []byte {
	d, err := base64.StdEncoding.DecodeString(e)

	if err != nil {
		Error(err.Error())
	}

	return d
}

func B64e(d []byte) string {
	return base64.StdEncoding.EncodeToString(d)
}

func P64(v uint64) []byte {
	b := make([]byte, 8)
	binary.LittleEndian.PutUint64(b, v)
	return b
}

func P32(v uint32) []byte {
	b := make([]byte, 4)
	binary.LittleEndian.PutUint32(b, v)
	return b
}

func P16(v uint16) []byte {
	b := make([]byte, 2)
	binary.LittleEndian.PutUint16(b, v)
	return b
}

func P8(v uint8) []byte {
	return []byte{v}
}

func U64(b []byte) uint64 {
	if len(b) != 8 {
		Error("U64 requires a buffer of 8 bytes")
	}

	return binary.LittleEndian.Uint64(b)
}

func U32(b []byte) uint32 {
	if len(b) != 4 {
		Error("U32 requires a buffer of 4 bytes")
	}

	return binary.LittleEndian.Uint32(b)
}
func U16(b []byte) uint16 {
	if len(b) != 2 {
		Error("U16 requires a buffer of 2 bytes")
	}

	return binary.LittleEndian.Uint16(b)
}
func U8(b []byte) uint8 {
	if len(b) != 1 {
		Error("U8 requires a buffer of 1 byte")
	}

	return b[0]
}
