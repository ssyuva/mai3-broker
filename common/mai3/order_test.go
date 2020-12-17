package mai3

import (
	"github.com/mcarloai/mai-v3-broker/common/mai3/utils"
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGenerateOrderData(t *testing.T) {
	var (
		expires   = 1606217568
		version   = 1
		orderType = 1
		closeOnly = true
		salt      = 123456
	)
	orderData := GenerateOrderData(int64(expires), int32(version), int8(orderType), closeOnly, int64(salt))
	assert.Equal(t, "0x000000005fbcef60000000010101000000000001e24000000000000000000000", orderData)
}

func TestGetOrderHash(t *testing.T) {
	var (
		trader     = "0x0000000000000000000000000000000000000001"
		broker     = "0x0000000000000000000000000000000000000002"
		relayer    = "0x0000000000000000000000000000000000000003"
		perpetual  = "0x0000000000000000000000000000000000000004"
		referrer   = "0x0000000000000000000000000000000000000005"
		amount     = decimal.NewFromFloat(1000).Div(decimal.NewFromInt(1000000000000000000))
		priceLimit = decimal.NewFromFloat(2000).Div(decimal.NewFromInt(1000000000000000000))
		orderData  = "0x000000005fbcef60000000010101000000000001e24000000000000000000000"
		chainID    = 1
	)
	orderHash, err := GetOrderHash(trader, broker, relayer, perpetual, referrer, orderData, amount, priceLimit, int64(chainID))
	assert.Nil(t, err)
	assert.Equal(t, "0x5cb3d66525445c1556d92eb52a8783bd5f0a81d8862bf72c6c2dfcef4e1c1290", utils.Bytes2HexP(orderHash))
}
