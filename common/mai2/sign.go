package mai2

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"strings"

	"github.com/mcarloai/mai-v3-broker/common/mai2/crypto"
	"github.com/mcarloai/mai-v3-broker/common/mai2/utils"
	"github.com/shopspring/decimal"
)

const orderDataFeeRateBase int64 = 100000

var EIP712_DOMAIN_TYPEHASH []byte
var EIP712_MAI2_ORDER_TYPE []byte

func init() {
	EIP712_DOMAIN_TYPEHASH = crypto.Keccak256([]byte(`EIP712Domain(string name)`))
	EIP712_MAI2_ORDER_TYPE = crypto.Keccak256([]byte(`Order(address trader,address broker,address perpetual,uint256 amount,uint256 price,bytes32 data)`))
}

func feeRateToHex(rate decimal.Decimal) string {
	f, _ := rate.Float64()
	n := int16(f * float64(orderDataFeeRateBase))
	b := bytes.NewBuffer(nil)
	if err := binary.Write(b, binary.BigEndian, n); err != nil {
		panic(fmt.Errorf("feeRateToHex error %+v", err))
	}
	return utils.Bytes2Hex(b.Bytes())
}

type OrderSignatureType int

const (
	EthSign OrderSignatureType = iota
	EIP712
)

func IsValidOrderSignature(address string, orderID string, signature string) (bool, error) {
	// ethereum signature config: [:32] r[32:64] s[64:]
	// first byte of config is v
	sigBytes, err := utils.Hex2Bytes(signature)
	if err != nil {
		return false, fmt.Errorf("IsValidOrderSignature:hex to bytes error:%w", err)
	}

	if len(sigBytes) != 96 {
		return false, fmt.Errorf("IsValidOrderSignature:order signature for ethereum should have 96 bytes. %s", signature)
	}

	ethSig := make([]byte, 65)
	copy(ethSig[:64], sigBytes[32:])
	ethSig[64] = sigBytes[0]

	// the 2nd byte of config is signature type
	method := EthSign
	if sigBytes[1] == 1 {
		method = EIP712
	}

	res, err := IsValidSignature(address, orderID, utils.Bytes2HexP(ethSig), method)
	if err != nil {
		err = fmt.Errorf("IsValidOrderSignature:valid order signature fail: %w", err)
	}
	return res, err
}

func IsValidSignature(address string, message string, signature string, method OrderSignatureType) (bool, error) {
	if len(address) != 42 {
		return false, fmt.Errorf("IsValidSignature:address must be 42 size long")
	}

	if len(signature) != 132 {
		return false, fmt.Errorf("IsValidSignature:signature must be 132 size long")
	}

	var hashBytes []byte
	if strings.HasPrefix(message, "0x") {
		var err error
		hashBytes, err = utils.Hex2Bytes(message)
		if err != nil {
			return false, fmt.Errorf("IsValidSignature:%w", err)
		}
	} else {
		hashBytes = []byte(message)
	}

	signatureByte, err := utils.Hex2Bytes(signature)
	if err != nil {
		return false, fmt.Errorf("IsValidSignature:%w", err)
	}
	switch method {
	case EthSign:
		pk, err := crypto.PersonalEcRecover(hashBytes, signatureByte)
		if err != nil {
			return false, err
		}
		return "0x"+strings.ToLower(pk) == strings.ToLower(address), nil
	case EIP712:
		pk, err := crypto.EIP712EcRecover(hashBytes, signatureByte)
		if err != nil {
			return false, err
		}
		return "0x"+strings.ToLower(pk) == strings.ToLower(address), nil
	default:
		return false, fmt.Errorf("IsValidSignature:unknown signature method")
	}
}

func getDomainSeparator() []byte {
	return crypto.Keccak256(
		EIP712_DOMAIN_TYPEHASH,
		crypto.Keccak256([]byte("Mai Protocol")),
	)
}

func getEIP712MessageHash(message []byte) []byte {
	return crypto.Keccak256(
		[]byte{'\x19', '\x01'},
		getDomainSeparator(),
		message,
	)
}

func addTailingZero(data string, length int) string {
	return data + strings.Repeat("0", length-len(data))
}

func addLeadingZero(data string, length int) string {
	return strings.Repeat("0", length-len(data)) + data
}
