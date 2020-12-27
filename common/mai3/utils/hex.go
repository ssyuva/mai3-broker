package utils

import (
	"encoding/hex"
	"strings"
)

func Bytes2Hex(bytes []byte) string {
	return hex.EncodeToString(bytes)
}

func Hex2Bytes(str string) ([]byte, error) {
	if strings.HasPrefix(str, "0x") || strings.HasPrefix(str, "0X") {
		str = str[2:]
	}

	if len(str)%2 == 1 {
		str = "0" + str
	}

	return hex.DecodeString(str)
}

// with prefix '0x'
func Bytes2HexP(bytes []byte) string {
	return "0x" + hex.EncodeToString(bytes)
}

const (
	// HashLength is the expected length of the hash
	HashLength = 32
)

type Hash [HashLength]byte

// BytesToHash sets b to hash.
// If b is larger than len(h), b will be cropped from the left.

func (h *Hash) SetBytes(b []byte) {
	if len(b) > len(h) {
		b = b[len(b)-HashLength:]
	}

	copy(h[HashLength-len(b):], b)
}

func (h Hash) Bytes() []byte { return h[:] }

func HexToHash(s string) (Hash, error) {
	var h Hash
	b, err := Hex2Bytes(s)
	if err != nil {
		return h, err
	}
	h = BytesToHash(b)
	return h, nil
}

func BytesToHash(b []byte) Hash {
	var h Hash
	h.SetBytes(b)
	return h
}
