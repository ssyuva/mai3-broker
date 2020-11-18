package ethereum

import (
	"crypto/ecdsa"
	"fmt"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/pborman/uuid"
	"github.com/pkg/errors"
)

// Account represents eth user account for further operations
type Account struct {
	address common.Address
	signer  *bind.TransactOpts
	private *ecdsa.PrivateKey
}

// Signer returns signer to sign transaction
func (acc *Account) Signer() *bind.TransactOpts {
	return acc.signer
}

func (acc *Account) String() string {
	return acc.address.Hex()
}

func (acc *Account) Address() common.Address {
	return acc.address
}

func (acc *Account) PersonalSign(msg []byte) ([]byte, error) {
	sigBytes, err := crypto.Sign(wrapMessage(msg), acc.private)
	if err != nil {
		return nil, errors.Wrap(err, "sign failed")
	}
	if sigBytes[64] < 27 {
		sigBytes[64] = sigBytes[64] + 27
	}
	return sigBytes, nil
}

func wrapMessage(data []byte) []byte {
	msg := fmt.Sprintf("\x19Ethereum Signed Message:\n%d%s", len(data), data)
	return crypto.Keccak256Hash([]byte(msg)).Bytes()
}

// PrivateToAccount imports an account from given private key
func PrivateToAccount(p *ecdsa.PrivateKey) *Account {
	addr := crypto.PubkeyToAddress(p.PublicKey).Hex()
	return &Account{
		address: common.HexToAddress(addr),
		signer:  bind.NewKeyedTransactor(p),
		private: p,
	}
}

func HexToPrivate(pk string) (*ecdsa.PrivateKey, error) {
	return crypto.HexToECDSA(pk)
}

func EncryptKey(pk *ecdsa.PrivateKey, password string) ([]byte, error) {
	key := &keystore.Key{
		Id:         uuid.NewRandom(),
		Address:    crypto.PubkeyToAddress(pk.PublicKey),
		PrivateKey: pk,
	}
	keyjson, err := keystore.EncryptKey(key, password, keystore.StandardScryptN, keystore.StandardScryptP)
	if err != nil {
		return nil, err
	}
	return keyjson, err
}

func DecryptKey(cipher []byte, password string) (*ecdsa.PrivateKey, error) {
	key, err := keystore.DecryptKey(cipher, password)
	if err != nil {
		return nil, err
	}
	return key.PrivateKey, nil
}
