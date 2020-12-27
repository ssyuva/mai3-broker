package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHex2Bytes(t *testing.T) {
	bytes1, _ := Hex2Bytes("ffff")
	assert.EqualValues(t, []byte{0xff, 0xff}, bytes1)
	bytes2, _ := Hex2Bytes("fff")
	assert.EqualValues(t, []byte{0x0f, 0xff}, bytes2)
	bytes3, _ := Hex2Bytes("0xffff")
	assert.EqualValues(t, []byte{0xff, 0xff}, bytes3)
}

func TestBytes2Hex(t *testing.T) {
	assert.EqualValues(t, "ffff", Bytes2Hex([]byte{0xff, 0xff}))
	assert.EqualValues(t, "ff0f", Bytes2Hex([]byte{0xff, 0xf}))
}

func TestBytes2HexP(t *testing.T) {
	assert.EqualValues(t, "0xff12", Bytes2HexP([]byte{0xff, 0x12}))
}
