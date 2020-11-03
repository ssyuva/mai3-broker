package ethereum

import (
	"encoding/hex"
	"fmt"
	"math/big"
	"strings"

	gethCommon "github.com/ethereum/go-ethereum/common"
)

func HexToAddress(str string) (addr gethCommon.Address, err error) {
	var b []byte
	b, err = fromHex(str)
	if err != nil {
		return
	}
	if len(b) != gethCommon.AddressLength {
		err = fmt.Errorf("invalid eth address:%s", str)
		return
	}
	addr.SetBytes(b)
	return
}

func MustHexToAddress(str string) gethCommon.Address {
	addr, err := HexToAddress(str)
	if err != nil {
		panic("[bug]" + err.Error())
	}
	return addr
}

func HexArrayToAddresses(addresses []string) ([]gethCommon.Address, error) {
	ret := make([]gethCommon.Address, 0)
	for _, addrStr := range addresses {
		addr, err := HexToAddress(addrStr)
		if err != nil {
			return nil, err
		}
		ret = append(ret, addr)
	}
	return ret, nil
}

func HexToHash(str string) (hash gethCommon.Hash, err error) {
	var b []byte
	b, err = fromHex(str)
	if err != nil {
		return
	}
	if len(b) != gethCommon.HashLength {
		err = fmt.Errorf("invalid eth hash:%s", str)
		return
	}
	hash.SetBytes(b)
	return
}

func MustHexToHash(str string) gethCommon.Hash {
	hash, err := HexToHash(str)
	if err != nil {
		panic("[bug]" + err.Error())
	}
	return hash
}

func fromHex(s string) ([]byte, error) {
	if has0xPrefix(s) {
		s = s[2:]
	}
	if len(s)%2 == 1 {
		s = "0" + s
	}
	return hex.DecodeString(s)
}

func has0xPrefix(str string) bool {
	return len(str) >= 2 && str[0] == '0' && (str[1] == 'x' || str[1] == 'X')
}

func StringToBytes32(str string) (result [32]byte) {
	forceConvert := []byte(str)
	if len(forceConvert) > 32 {
		forceConvert = forceConvert[0:32]
	}
	copy(result[:], forceConvert)
	return
}

func BigIntToAddressString(n *big.Int) (addressString string, err error) {
	b := n.Bytes()
	if len(b) != gethCommon.AddressLength {
		err = fmt.Errorf("invalid eth address:%s", n.String())
		return
	}
	var addr gethCommon.Address
	addr.SetBytes(b)
	addressString = strings.ToLower(addr.Hex())
	return
}
