package ethereum

import (
	"crypto/ecdsa"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/google/uuid"
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
func PrivateToAccount(p *ecdsa.PrivateKey, chainID *big.Int) (*Account, error) {
	addr := crypto.PubkeyToAddress(p.PublicKey).Hex()
	signer, err := bind.NewKeyedTransactorWithChainID(p, chainID)
	if err != nil {
		return nil, err
	}
	return &Account{
		address: common.HexToAddress(addr),
		signer:  signer,
		private: p,
	}, nil
}

func (c *Client) HexToPrivate(pk string) (*ecdsa.PrivateKey, string, error) {
	private, err := crypto.HexToECDSA(pk)
	if err != nil {
		return nil, "", err
	}
	addr := crypto.PubkeyToAddress(private.PublicKey).Hex()
	return private, addr, nil
}

func (c *Client) EncryptKey(pk *ecdsa.PrivateKey, password string) ([]byte, error) {
	id, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}

	key := &keystore.Key{
		Id:         id,
		Address:    crypto.PubkeyToAddress(pk.PublicKey),
		PrivateKey: pk,
	}
	keyjson, err := keystore.EncryptKey(key, password, keystore.StandardScryptN, keystore.StandardScryptP)
	if err != nil {
		return nil, err
	}
	return keyjson, err
}

func (c *Client) DecryptKey(cipher []byte, password string) (*ecdsa.PrivateKey, error) {
	key, err := keystore.DecryptKey(cipher, password)
	if err != nil {
		return nil, err
	}
	return key.PrivateKey, nil
}
