package pwntools

import (
	"encoding/base64"
	"encoding/hex"
)

type Number interface {
	int | string
}

func UnHex(h string) []byte {
	b, err := hex.DecodeString(h)

	if err != nil {
		panic(err)
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

func B64d(enc string) []byte {
	dec, _ := base64.StdEncoding.DecodeString(enc)
	return dec
}

func B64e(data []byte) string {
	return base64.StdEncoding.EncodeToString(data)
}
