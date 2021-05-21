package crypto

import (
	"testing"

	"github.com/mcdexio/mai3-broker/common/mai3/utils"
	"github.com/stretchr/testify/assert"
)

func TestNewPrivateKey(t *testing.T) {
	address := "0x93388b4efe13b9b18ed480783c05462409851547"
	prvKeyHex := "95b0a982c0dfc5ab70bf915dcf9f4b790544d25bc5e6cff0f38a59d0bba58651"
	expect := address

	act, _ := NewPrivateKeyByHex(prvKeyHex)
	assert.EqualValues(t, expect, PubKey2Address(act.PublicKey))

	prvKeyHex2 := "95b0a982c0dfc5ab70bf915dcf9f4b790544d25bc5e6cff0f38a59d0bba586"
	act2, err := NewPrivateKeyByHex(prvKeyHex2)
	assert.Nil(t, act2)
	assert.EqualValues(t, err.Error(), "invalid length, need 256 bits")

	prvKeyHex3 := "ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff"
	act3, err := NewPrivateKeyByHex(prvKeyHex3)
	assert.Nil(t, act3)
	assert.EqualValues(t, err.Error(), "invalid private key, >=N")

	prvKeyHex4 := "0000000000000000000000000000000000000000000000000000000000000000"
	act4, err := NewPrivateKeyByHex(prvKeyHex4)
	assert.Nil(t, act4)
	assert.EqualValues(t, err.Error(), "invalid private key, zero or negative")

}

func TestNewPrivateKeyByHex(t *testing.T) {
	address := "0x93388b4efe13b9b18ed480783c05462409851547"
	prvKeyHex := "95b0a982c0dfc5ab70bf915dcf9f4b790544d25bc5e6cff0f38a59d0bba58651"
	prvKeyBytes, _ := utils.Hex2Bytes(prvKeyHex)
	expect := address

	act, _ := NewPrivateKey(prvKeyBytes)

	assert.EqualValues(t, expect, PubKey2Address(act.PublicKey))
}

func TestSign(t *testing.T) {
	prvKeyHex := "95b0a982c0dfc5ab70bf915dcf9f4b790544d25bc5e6cff0f38a59d0bba58651"
	message, _ := utils.Hex2Bytes("9df8dba3720d00bd48ad744722021ef91b035e273bccfb78660ca8df9574b086")
	expect, _ := utils.Hex2Bytes("2736b2ca3e2d4778e53a33e0d9bb2d9bad91ec858ab71ad49e31f540f15728a83dbea28bd686bb66d06e4ad9f48912ef437b92a272ea47563c2df80ed59b508e00")
	actKey, _ := NewPrivateKeyByHex(prvKeyHex)
	act, _ := Sign([]byte(message), actKey)
	assert.EqualValues(t, expect, act)
}

func TestPersonalSignAndPersonalSignByPrivateKey(t *testing.T) {
	prvKeyHex := "95b0a982c0dfc5ab70bf915dcf9f4b790544d25bc5e6cff0f38a59d0bba58651"
	message, _ := utils.Hex2Bytes("9df8dba3720d00bd48ad744722021ef91b035e273bccfb78660ca8df9574b086")
	expect, _ := utils.Hex2Bytes("aa7cd9f5a7eb485771215d45cc2a4c535e270c75c3595ae6b1c158aef72e67066ad5df037ad5945c65da90edfaa4fe418e5b6bd2225ec9d4b704433a779e4bff1b")

	actKey, _ := NewPrivateKeyByHex(prvKeyHex)
	act, _ := PersonalSignByPrivateKey(message, actKey)
	act2, _ := PersonalSign(message, prvKeyHex)
	assert.EqualValues(t, expect, act)
	assert.EqualValues(t, expect, act2)
}

func TestEcRecover(t *testing.T) {
	sign, _ := utils.Hex2Bytes("2736b2ca3e2d4778e53a33e0d9bb2d9bad91ec858ab71ad49e31f540f15728a83dbea28bd686bb66d06e4ad9f48912ef437b92a272ea47563c2df80ed59b508e00")
	message := Keccak256([]byte("some message"))
	expect, _ := utils.Hex2Bytes("0450d7aa97f7496fd412f393e54df0cbe3f6cbeacf15d1ddb12133e408522feb8896dd1652ee84b18788bc7753663302a6489f779352bbfec010ab25c9e3806843")

	act, _ := EcRecover(message, sign)
	assert.EqualValues(t, expect, act)
}

func TestPersonalEcRecover(t *testing.T) {
	address := "0x93388b4efe13b9b18ed480783c05462409851547"
	sign, _ := utils.Hex2Bytes("aa7cd9f5a7eb485771215d45cc2a4c535e270c75c3595ae6b1c158aef72e67066ad5df037ad5945c65da90edfaa4fe418e5b6bd2225ec9d4b704433a779e4bff1b")
	message := Keccak256([]byte("some message"))

	act, err := PersonalEcRecover(message, sign)
	assert.Nil(t, err)
	assert.EqualValues(t, address[2:], act)
}

func TestEIP712EcRecover(t *testing.T) {
	address := "0x39e38953aa0822bac469da1c252815ddefbfb515"
	sign, _ := utils.Hex2Bytes("1ffc7cae091494c377794d935dcae1cf9f6da232bd21943a8a91888b5cc4505376882744e97cef61f2e49e6e7390c3a2cdde1fa5279b3d853fd987d736be77f81b")
	message, _ := utils.Hex2Bytes("8d37453e1248f94614963b3e48b54c5f9393f9216d3c674ff8c0a761be50a78e")

	act, err := EIP712EcRecover(message, sign)
	assert.Nil(t, err)
	assert.EqualValues(t, address[2:], act)
}

func TestSigToPub(t *testing.T) {
	address := "0x93388b4efe13b9b18ed480783c05462409851547"
	message := Keccak256([]byte("some message"))
	sign, _ := utils.Hex2Bytes("aa7cd9f5a7eb485771215d45cc2a4c535e270c75c3595ae6b1c158aef72e67066ad5df037ad5945c65da90edfaa4fe418e5b6bd2225ec9d4b704433a779e4bff00")

	act, _ := SigToPub(hashPersonalMessage(message), sign)
	assert.EqualValues(t, address, PubKey2Address(*act))
}
