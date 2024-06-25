package pwntools

import (
	"encoding/base64"
	"encoding/binary"
	"encoding/hex"
)

type number interface {
	int | string
}

// UnHex hex-decodes a string
func UnHex(h string) []byte {
	b, err := hex.DecodeString(h)

	if err != nil {
		Error(err.Error())
	}

	return b
}

// Hex creates a hex-string from a byte slice
func Hex(b []byte) string {
	return hex.EncodeToString(b)
}

// Xor applies a byte-wise XOR operation to the byte slices passed as arguments
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

// B64d Base64-decodes a string
func B64d(e string) []byte {
	d, err := base64.StdEncoding.DecodeString(e)

	if err != nil {
		Error(err.Error())
	}

	return d
}

// B64e Base64-encodes a string
func B64e(d []byte) string {
	return base64.StdEncoding.EncodeToString(d)
}

// P64 packs a 64-bit unsigned integer
func P64(v uint64) []byte {
	b := make([]byte, 8)
	binary.LittleEndian.PutUint64(b, v)
	return b
}

// P32 packs a 32-bit unsigned integer
func P32(v uint32) []byte {
	b := make([]byte, 4)
	binary.LittleEndian.PutUint32(b, v)
	return b
}

// P16 packs a 16-bit unsigned integer
func P16(v uint16) []byte {
	b := make([]byte, 2)
	binary.LittleEndian.PutUint16(b, v)
	return b
}

// P8 packs an 8-bit unsigned integer
func P8(v uint8) []byte {
	return []byte{v}
}

// U64 unpacks a 64-bit unsigned integer
func U64(b []byte) uint64 {
	if len(b) != 8 {
		Error("U64 requires a buffer of 8 bytes")
	}

	return binary.LittleEndian.Uint64(b)
}

// U32 unpacks a 32-bit unsigned integer
func U32(b []byte) uint32 {
	if len(b) != 4 {
		Error("U32 requires a buffer of 4 bytes")
	}

	return binary.LittleEndian.Uint32(b)
}

// U16 unpacks a 16-bit unsigned integer
func U16(b []byte) uint16 {
	if len(b) != 2 {
		Error("U16 requires a buffer of 2 bytes")
	}

	return binary.LittleEndian.Uint16(b)
}

// U8 unpacks an 8-bit unsigned integer
func U8(b []byte) uint8 {
	if len(b) != 1 {
		Error("U8 requires a buffer of 1 byte")
	}

	return b[0]
}
