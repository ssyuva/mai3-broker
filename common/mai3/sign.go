package mai3

import (
	"fmt"
	"strings"

	"github.com/mcdexio/mai3-broker/common/mai3/crypto"
	"github.com/mcdexio/mai3-broker/common/mai3/utils"
)

const (
	EthSign = "ethSign"
	EIP712  = "eip712"
)

func IsValidSignature(address string, message string, signature string, method string) (bool, error) {
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
