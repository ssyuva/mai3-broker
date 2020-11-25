package mai3

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"github.com/mcarloai/mai-v3-broker/common/mai3/crypto"
	"github.com/mcarloai/mai-v3-broker/common/mai3/utils"
	"github.com/shopspring/decimal"
	"math/big"
	"strings"
)

const orderDataFeeRateBase int64 = 100000

var EIP712_DOMAIN_TYPEHASH []byte
var EIP712_MAI3_ORDER_TYPE []byte

func init() {
	EIP712_DOMAIN_TYPEHASH = crypto.Keccak256([]byte(`EIP712Domain(string name)`))
	EIP712_MAI3_ORDER_TYPE = crypto.Keccak256([]byte(`Order(address trader,address broker,address relayer,address perpetual,address referrer,int256 amount,int256 priceLimit,uint64 deadline,uint32 version,OrderType orderType,bool isCloseOnly,uint64 salt,uint256 chainID)`))
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

func addTailingZero(data string, length int) string {
	return data + strings.Repeat("0", length-len(data))
}

func addLeadingZero(data string, length int) string {
	return strings.Repeat("0", length-len(data)) + data
}

func GetOrderHash(traderAddress, brokerAddress, relayerAddress, contractAddress, referrerAddress string, amount, price decimal.Decimal,
	expiredAtSeconds int64, version int32, ordrType int8, isCloseOnly bool, salt, chainID int64) ([]byte, error) {
	trader, err := utils.HexToHash(traderAddress)
	if err != nil {
		return nil, fmt.Errorf("GetOrderHash:%w", err)
	}
	broker, err := utils.HexToHash(brokerAddress)
	if err != nil {
		return nil, fmt.Errorf("GetOrderHash:%w", err)
	}
	relayer, err := utils.HexToHash(relayerAddress)
	if err != nil {
		return nil, fmt.Errorf("GetOrderHash:%w", err)
	}
	contract, err := utils.HexToHash(contractAddress)
	if err != nil {
		return nil, fmt.Errorf("GetOrderHash:%w", err)
	}
	referrer, err := utils.HexToHash(referrerAddress)
	if err != nil {
		return nil, fmt.Errorf("GetOrderHash:%w", err)
	}

	amountBin := utils.BytesToHash(utils.MustDecimalToBigInt(amount).Bytes())
	priceBin := utils.BytesToHash(utils.MustDecimalToBigInt(price).Bytes())
	expiredAtSecondsBin := utils.BytesToHash(utils.Int64ToBytes(expiredAtSeconds))
	versionBin := utils.BytesToHash(utils.Int32ToBytes(version))
	ordrTypeBin := utils.BytesToHash(utils.Int8ToBytes(ordrType))
	isCloseOnlyBin := utils.BytesToHash(utils.BoolToBytes(isCloseOnly))
	saltBin := utils.BytesToHash(utils.Int64ToBytes(salt))
	chainIDBin := utils.BytesToHash(big.NewInt(chainID).Bytes())

	hash := getEIP712MessageHash(
		crypto.Keccak256(
			EIP712_MAI3_ORDER_TYPE,
			trader.Bytes(),
			broker.Bytes(),
			relayer.Bytes(),
			contract.Bytes(),
			referrer.Bytes(),
			amountBin.Bytes(),
			priceBin.Bytes(),
			expiredAtSecondsBin.Bytes(),
			versionBin.Bytes(),
			ordrTypeBin.Bytes(),
			isCloseOnlyBin.Bytes(),
			saltBin.Bytes(),
			chainIDBin.Bytes(),
		),
	)
	return hash, nil
}

func getDomainSeparator() []byte {
	return crypto.Keccak256(
		EIP712_DOMAIN_TYPEHASH,
		crypto.Keccak256([]byte("Mai Protocol v3")),
	)
}

func getEIP712MessageHash(message []byte) []byte {
	hash := crypto.Keccak256(
		[]byte{'\x19', '\x01'},
		getDomainSeparator(),
		message,
	)

	return crypto.Keccak256(
		[]byte("\x19Ethereum Signed Message:\n32"),
		hash,
	)
}
