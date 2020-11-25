package mai3

import (
	"fmt"
	"github.com/mcarloai/mai-v3-broker/common/mai3/crypto"
	"github.com/mcarloai/mai-v3-broker/common/mai3/utils"
	"github.com/shopspring/decimal"
	"math/big"
	"strconv"
	"strings"
)

var EIP712_DOMAIN_TYPEHASH []byte
var EIP712_MAI3_ORDER_TYPE []byte

func init() {
	EIP712_DOMAIN_TYPEHASH = crypto.Keccak256([]byte(`EIP712Domain(string name)`))
	EIP712_MAI3_ORDER_TYPE = crypto.Keccak256([]byte(`Order(address trader,address broker,address relayer,address perpetual,address referrer,int256 amount,int256 priceLimit,uint64 deadline,uint32 version,OrderType orderType,bool isCloseOnly,uint64 salt,uint256 chainID)`))
}

func addTailingZero(data string, length int) string {
	return data + strings.Repeat("0", length-len(data))
}

func addLeadingZero(data string, length int) string {
	return strings.Repeat("0", length-len(data)) + data
}

func GetOrderData(deadline int64, version int32, ordrType int8, isCloseOnly bool, salt int64) ([32]byte, error) {
	var orderData [32]byte
	data := GenerateOrderData(deadline, version, ordrType, isCloseOnly, salt)
	b, err := utils.Hex2Bytes(data)
	if err != nil {
		return orderData, fmt.Errorf("GetOrderData:%w", err)
	}
	if len(b) != len(orderData) {
		return orderData, fmt.Errorf("GetOrderData:bad length")
	}
	copy(orderData[0:32], b)

	return orderData, nil
}

func GenerateOrderData(expiredAtSeconds int64, version int32, ordrType int8, isCloseOnly bool, salt int64) string {
	data := strings.Builder{}
	data.WriteString("0x")
	data.WriteString(addLeadingZero(fmt.Sprintf("%x", expiredAtSeconds), 8*2))
	data.WriteString(addLeadingZero(strconv.FormatInt(int64(version), 16), 4*2))
	data.WriteString(addLeadingZero(strconv.FormatInt(int64(ordrType), 16), 2))

	if isCloseOnly {
		data.WriteString("01")
	} else {
		data.WriteString("00")
	}
	data.WriteString(addLeadingZero(fmt.Sprintf("%x", salt), 8*2))
	return addTailingZero(data.String(), 66)
}

func GetOrderHash(traderAddress, brokerAddress, relayerAddress, contractAddress, referrerAddress, orderData string, amount, price decimal.Decimal, chainID int64) ([]byte, error) {
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
	orderDataBin, err := utils.HexToHash(orderData)
	if err != nil {
		return nil, fmt.Errorf("GetOrderHash:%w", err)
	}
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
			orderDataBin.Bytes(),
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
