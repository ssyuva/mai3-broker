package mai3

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"github.com/mcarloai/mai-v3-broker/common/mai3/crypto"
	"github.com/mcarloai/mai-v3-broker/common/mai3/utils"
	"github.com/shopspring/decimal"
	"strconv"
	"strings"
)

const orderDataFeeRateBase int64 = 100000

var EIP712_DOMAIN_TYPEHASH []byte
var EIP712_MAI3_ORDER_TYPE []byte

func init() {
	EIP712_DOMAIN_TYPEHASH = crypto.Keccak256([]byte(`EIP712Domain(string name)`))
	EIP712_MAI3_ORDER_TYPE = crypto.Keccak256([]byte(`Order(address trader,address broker,address perpetual,uint256 amount,uint256 price,bytes32 data)`))
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

func GenerateOrderData(version MaiProtocolVersion, expiredAtSeconds, salt, chainID int64, asMakerFeeRate, asTakerFeeRate, makerRebateRate decimal.Decimal, isSell, isMarket, isMakerOnly bool, isInverse bool) string {
	data := strings.Builder{}
	data.WriteString("0x")
	data.WriteString(addLeadingZero(strconv.FormatInt(int64(version), 16), 2))
	if isSell {
		data.WriteString("01")
	} else {
		data.WriteString("00")
	}

	if isMarket {
		data.WriteString("01")
	} else {
		data.WriteString("00")
	}

	data.WriteString(addLeadingZero(fmt.Sprintf("%x", expiredAtSeconds), 5*2))
	data.WriteString(addLeadingZero(feeRateToHex(asMakerFeeRate), 2*2))
	data.WriteString(addLeadingZero(feeRateToHex(asTakerFeeRate), 2*2))
	data.WriteString(addLeadingZero(feeRateToHex(makerRebateRate), 2*2))
	data.WriteString(addLeadingZero(fmt.Sprintf("%x", salt), 8*2))

	if isMakerOnly {
		data.WriteString("01")
	} else {
		data.WriteString("00")
	}

	if isInverse {
		data.WriteString("01")
	} else {
		data.WriteString("00")
	}
	data.WriteString(addLeadingZero(fmt.Sprintf("%x", chainID), 8*2))

	return addTailingZero(data.String(), 66)
}

func GetOrderHash(traderAddress, relayerAddress, contractAddress, orderData string, amount, price decimal.Decimal) ([]byte, error) {
	trader, err := utils.HexToHash(traderAddress)
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

	amountBin := utils.BytesToHash(utils.MustDecimalToBigInt(amount).Bytes())
	priceBin := utils.BytesToHash(utils.MustDecimalToBigInt(price).Bytes())
	orderDataBin, err := utils.HexToHash(orderData)
	if err != nil {
		return nil, fmt.Errorf("GetOrderHash:%w", err)
	}

	hash := getEIP712MessageHash(
		crypto.Keccak256(
			EIP712_MAI3_ORDER_TYPE,
			trader.Bytes(),
			relayer.Bytes(),
			contract.Bytes(),
			amountBin.Bytes(),
			priceBin.Bytes(),
			orderDataBin.Bytes(),
		),
	)
	return hash, nil
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
