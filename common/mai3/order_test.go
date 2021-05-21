package mai3

import (
	"testing"

	"github.com/mcdexio/mai3-broker/common/mai3/utils"
	"github.com/mcdexio/mai3-broker/common/model"
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
)

func TestStopLimitOrderFlags(t *testing.T) {
	var (
		orderType   = model.StopLimitOrder
		isCloseOnly = true
	)
	flags := GenerateOrderFlags(orderType, isCloseOnly)
	assert.Equal(t, 0x0a8000000, flags)
}

func TestLimitOrderFlags(t *testing.T) {
	var (
		orderType   = model.LimitOrder
		isCloseOnly = false
	)
	flags := GenerateOrderFlags(orderType, isCloseOnly)
	assert.Equal(t, 0x008000000, flags)
}

func TestTakeProfitOrderFlags(t *testing.T) {
	var (
		orderType   = model.TakeProfitOrder
		isCloseOnly = true
	)
	flags := GenerateOrderFlags(orderType, isCloseOnly)
	assert.Equal(t, 0x098000000, flags)
}

func TestGenerateOrderData(t *testing.T) {
	var (
		trader         = "0x1111111111111111111111111111111111111111"
		broker         = "0x2222222222222222222222222222222222222222"
		relayer        = "0x3333333333333333333333333333333333333333"
		referrer       = "0x0000000000000000000000000000000000000000"
		pool           = "0x5555555555555555555555555555555555555555"
		minTradeAmount = decimal.NewFromFloat(7).Div(decimal.NewFromInt(1000000000000000000))
		amount         = decimal.NewFromFloat(-1)
		priceLimit     = decimal.NewFromFloat(9).Div(decimal.NewFromInt(1000000000000000000))
		triggerPrice   = decimal.NewFromFloat(10).Div(decimal.NewFromInt(1000000000000000000))
		chainID        = 15
		expires        = 11
		perpetualIndex = 6
		brokerFeeLimit = 12
		flags          = 0xffffffff
		salt           = 14
		signType       = EIP712
		v              = "1b"
		r              = "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"
		s              = "bbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbb"
	)
	orderData := GenerateOrderData(trader, broker, relayer, referrer, pool, minTradeAmount, amount, priceLimit, triggerPrice, int64(chainID),
		int64(expires), int64(perpetualIndex), int64(brokerFeeLimit), int64(flags), int64(salt), signType, v, r, s)
	assert.Equal(t, "0x111111111111111111111111111111111111111122222222222222222222222222222222222222223333333333333333333333333333333333333333000000000000000000000000000000000000000055555555555555555555555555555555555555550000000000000000000000000000000000000000000000000000000000000007fffffffffffffffffffffffffffffffffffffffffffffffff21f494c589c00000000000000000000000000000000000000000000000000000000000000000009000000000000000000000000000000000000000000000000000000000000000a000000000000000000000000000000000000000000000000000000000000000f000000000000000b000000060000000cffffffff0000000e1b01aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaabbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbb", orderData)
}

func TestGetOrderHash(t *testing.T) {
	var (
		trader         = "0x1111111111111111111111111111111111111111"
		broker         = "0x2222222222222222222222222222222222222222"
		relayer        = "0x3333333333333333333333333333333333333333"
		referrer       = "0x4444444444444444444444444444444444444444"
		pool           = "0x5555555555555555555555555555555555555555"
		minTradeAmount = decimal.NewFromFloat(7).Div(decimal.NewFromInt(1000000000000000000))
		amount         = decimal.NewFromFloat(8).Div(decimal.NewFromInt(1000000000000000000))
		priceLimit     = decimal.NewFromFloat(9).Div(decimal.NewFromInt(1000000000000000000))
		triggerPrice   = decimal.NewFromFloat(10).Div(decimal.NewFromInt(1000000000000000000))
		chainID        = 15
		expires        = 11
		perpetualIndex = 6
		brokerFeeLimit = 12
		flags          = 0xffffffff
		salt           = 14
	)
	orderHash, err := GetOrderHash(trader, broker, relayer, referrer, pool, minTradeAmount, amount, priceLimit, triggerPrice, int64(chainID),
		int64(expires), int64(perpetualIndex), int64(brokerFeeLimit), int64(flags), int64(salt))
	assert.Nil(t, err)
	assert.Equal(t, "0xc6dd6530bc669ead14b253033fd6267e180b98cecf9d67a0df3955472552e867", utils.Bytes2HexP(orderHash))
}

func TestBigIntToBytes(t *testing.T) {
	assert.Equal(t, "00000000000000000000000000000000000000000000000029a2241af62c0000", encodeNumber(decimal.NewFromFloat(3)))
	assert.Equal(t, "ffffffffffffffffffffffffffffffffffffffffffffffffd65ddbe509d40000", encodeNumber(decimal.NewFromFloat(-3)))

}
