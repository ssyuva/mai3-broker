package mai3

import (
	"github.com/mcarloai/mai-v3-broker/common/mai3/utils"
	"github.com/mcarloai/mai-v3-broker/common/model"
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGenerateOrderFlags(t *testing.T) {
	var (
		orderType   = model.StopLimitOrder
		isCloseOnly = true
	)
	flags := GenerateOrderFlags(orderType, isCloseOnly)
	assert.Equal(t, 0x0a0000000, flags)
}

func TestGenerateOrderData(t *testing.T) {
	var (
		trader         = "0xe87c5839421579552e676ab0627ae03a7bf9b6d1"
		broker         = "0x0ce5953e6f05a6100e2fffb38d4163624912670b"
		relayer        = "0xd595f7c2c071d3fd8f5587931edf34e92f9ad39f"
		referrer       = "0x0000000000000000000000000000000000000000"
		pool           = "0x39b5b39de93e60081dcdc94a8b4180a8063959cc"
		minTradeAmount = decimal.NewFromFloat(1)
		amount         = decimal.NewFromFloat(0.007)
		priceLimit     = decimal.NewFromFloat(1343.3835)
		triggerPrice   = decimal.NewFromFloat(0)
		chainID        = 1337
		expires        = 1611113371
		perpetualIndex = 0
		brokerFeeLimit = 9000000
		flags          = 0x00000000
		salt           = 3057603
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
		trader         = "0xe87c5839421579552e676ab0627ae03a7bf9b6d1"
		broker         = "0x0ce5953e6f05a6100e2fffb38d4163624912670b"
		relayer        = "0xd595f7c2c071d3fd8f5587931edf34e92f9ad39f"
		referrer       = "0x0000000000000000000000000000000000000000"
		pool           = "0x39b5b39de93e60081dcdc94a8b4180a8063959cc"
		minTradeAmount = decimal.NewFromFloat(1)
		amount         = decimal.NewFromFloat(0.007)
		priceLimit     = decimal.NewFromFloat(1343.3835)
		triggerPrice   = decimal.NewFromFloat(0)
		chainID        = 1337
		expires        = 1611113371
		perpetualIndex = 0
		brokerFeeLimit = 9000000
		flags          = 0x00000000
		salt           = 3057603
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
