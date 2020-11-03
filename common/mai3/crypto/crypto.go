package crypto

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"errors"
	"fmt"
	"math/big"

	"github.com/btcsuite/btcd/btcec"
	"github.com/mcarloai/mai-v3-broker/common/mai3/utils"
	"golang.org/x/crypto/sha3"
)

var bitCurve = btcec.S256()

func Keccak256(data ...[]byte) []byte {
	d := sha3.NewLegacyKeccak256()
	for _, b := range data {
		_, _ = d.Write(b)
	}
	return d.Sum(nil)
}

func NewPrivateKey(privateKeyBytes []byte) (*ecdsa.PrivateKey, error) {
	priv := new(ecdsa.PrivateKey)
	priv.PublicKey.Curve = bitCurve
	if 8*len(privateKeyBytes) != priv.Params().BitSize {
		return nil, fmt.Errorf("invalid length, need %d bits", priv.Params().BitSize)
	}
	priv.D = new(big.Int).SetBytes(privateKeyBytes)

	// The priv.D must < N
	if priv.D.Cmp(bitCurve.N) >= 0 {
		return nil, fmt.Errorf("invalid private key, >=N")
	}
	// The priv.D must not be zero or negative.
	if priv.D.Sign() <= 0 {
		return nil, fmt.Errorf("invalid private key, zero or negative")
	}

	priv.PublicKey.X, priv.PublicKey.Y = priv.PublicKey.Curve.ScalarBaseMult(privateKeyBytes)
	if priv.PublicKey.X == nil {
		return nil, errors.New("invalid private key")
	}
	return priv, nil
}

func NewPrivateKeyByHex(privateKeyHex string) (*ecdsa.PrivateKey, error) {
	privateKeyBytes, err := utils.Hex2Bytes(privateKeyHex)
	if err != nil {
		return nil, fmt.Errorf("NewPrivateKeyByHex:private key hex to bytes error:%w", err)
	}
	return NewPrivateKey(privateKeyBytes)
}

func Sign(hash []byte, prv *ecdsa.PrivateKey) ([]byte, error) {
	if len(hash) != 32 {
		return nil, fmt.Errorf("hash is required to be exactly 32 bytes (%d)", len(hash))
	}
	if prv.Curve != btcec.S256() {
		return nil, fmt.Errorf("private key curve is not secp256k1")
	}
	sig, err := btcec.SignCompact(btcec.S256(), (*btcec.PrivateKey)(prv), hash, false)
	if err != nil {
		return nil, err
	}
	// Convert to Ethereum signature format with 'recovery id' v at the end.
	v := sig[0] - 27
	copy(sig, sig[1:])
	sig[64] = v
	return sig, nil
}

func PersonalSign(message []byte, privateKey string) ([]byte, error) {
	pk, err := NewPrivateKeyByHex(privateKey)
	if err != nil {
		return nil, err
	}
	return PersonalSignByPrivateKey(message, pk)
}

func PersonalSignByPrivateKey(message []byte, privateKey *ecdsa.PrivateKey) ([]byte, error) {
	personalHash := hashPersonalMessage(message)
	singBytes, err := Sign(personalHash, privateKey)
	if err != nil {
		return nil, err
	}

	// Since we are using HomesteadHash, the v is either 27 or 28
	// Mode details about EIP155 goes https://github.com/ethereum/EIPs/blob/master/EIPS/eip-155.md
	if singBytes[64] < 27 {
		singBytes[64] = singBytes[64] + 27
	}

	return singBytes, nil
}

func EcRecover(hash, sig []byte) ([]byte, error) {
	pub, err := SigToPub(hash, sig)
	if err != nil {
		return nil, err
	}
	bytes := (*btcec.PublicKey)(pub).SerializeUncompressed()
	return bytes, err
}

func PersonalEcRecover(data []byte, sig []byte) (string, error) {
	return EIP712EcRecover(hashPersonalMessage(data), sig)
}

func EIP712EcRecover(data []byte, sig []byte) (string, error) {
	if len(sig) != 65 {
		return "", fmt.Errorf("signature must be 65 bytes long")
	}
	if sig[64] >= 27 {
		sig[64] -= 27
	} else {
		return "", fmt.Errorf("we only accept EIP155. the v is either 27 or 28")
	}

	rpk, err := SigToPub(data, sig)
	if err != nil {
		return "", err
	}

	if rpk == nil || rpk.X == nil || rpk.Y == nil {
		return "", errors.New("")
	}
	pubBytes := elliptic.Marshal(bitCurve, rpk.X, rpk.Y)
	return utils.Bytes2Hex(Keccak256(pubBytes[1:])[12:]), nil
}

func hashPersonalMessage(data []byte) []byte {
	msg := fmt.Sprintf("\x19Ethereum Signed Message:\n%d%s", len(data), data)
	return Keccak256([]byte(msg))
}

func PubKey2Bytes(pub *ecdsa.PublicKey) []byte {
	if pub == nil || pub.X == nil || pub.Y == nil {
		return nil
	}
	return elliptic.Marshal(bitCurve, pub.X, pub.Y)
}

func PubKey2Address(p ecdsa.PublicKey) string {
	pubBytes := PubKey2Bytes(&p)
	return utils.Bytes2HexP(Keccak256(pubBytes[1:])[12:])
}

func SigToPub(hash, sig []byte) (*ecdsa.PublicKey, error) {
	// Convert to btcec input format with 'recovery id' v at the beginning.
	btcSig := make([]byte, 65)
	btcSig[0] = sig[64] + 27
	copy(btcSig[1:], sig)

	pub, _, err := btcec.RecoverCompact(btcec.S256(), btcSig, hash)
	return (*ecdsa.PublicKey)(pub), err
}
