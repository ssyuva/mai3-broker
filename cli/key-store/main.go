package main

import (
	"flag"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/mcarloai/mai-v3-broker/common/model"
	"github.com/mcarloai/mai-v3-broker/dao"
	"github.com/pborman/uuid"
)

var privateKey string
var password string
var dataBaseURL string

func Init() {
	flag.StringVar(&privateKey, "key", "", "relayer private key")
	flag.StringVar(&dataBaseURL, "url", "", "database url for kv store")
	flag.StringVar(&password, "pass", "", "password for encrypto private key")
}

func main() {
	var err error
	Init()
	flag.Parse()

	private, err := crypto.HexToECDSA(privateKey)
	if err != nil {
		fmt.Println(err)

		return
	}
	addr := crypto.PubkeyToAddress(private.PublicKey).Hex()

	key := &keystore.Key{
		Id:         uuid.NewRandom(),
		Address:    crypto.PubkeyToAddress(private.PublicKey),
		PrivateKey: private,
	}
	keyBytes, err := keystore.EncryptKey(key, password, keystore.StandardScryptN, keystore.StandardScryptP)
	if err != nil {
		fmt.Println(err)
		return
	}

	// init database
	if err = dao.ConnectPostgres(dataBaseURL); err != nil {
		fmt.Println(err)
		return
	}

	dao := dao.New()
	err = dao.Put(&model.KVStore{
		Key:      addr,
		Category: "keystore",
		Value:    keyBytes,
	})
	if err != nil {
		fmt.Println(err)
	}

	return
}
