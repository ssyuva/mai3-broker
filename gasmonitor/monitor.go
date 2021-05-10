package gasmonitor

import (
	"context"
	"time"

	"github.com/mcarloai/mai-v3-broker/conf"
	"github.com/shopspring/decimal"
	logger "github.com/sirupsen/logrus"
)

type GasMonitor struct {
	ctx      context.Context
	gasPrice uint64
}

var gwei, _ = decimal.NewFromString("1000000000")

func NewGasMonitor(ctx context.Context) *GasMonitor {
	gasMonitor := &GasMonitor{
		ctx: ctx,
	}

	go gasMonitor.run()
	return gasMonitor
}

// GetGasPrice return gas in Gwei
func (p *GasMonitor) GetGasPrice() uint64 {
	return p.gasPrice
}

// GasPriceGwei return gas in eth decimal
func (p *GasMonitor) GasPriceGwei() decimal.Decimal {
	return decimal.NewFromInt(int64(p.gasPrice)).Div(gwei)
}

func (p *GasMonitor) run() {
	logger.Infof("gas price monitor start")
	for {
		select {
		case <-p.ctx.Done():
			logger.Infof("gas price monitor end")
			return
		case <-time.After(10 * time.Second):
			gasPrice, err := p.getPriceInfo()
			if err != nil {
				logger.Errorf("fail to retrieve lastest gas price info:%s", err)
			} else {
				p.gasPrice = gasPrice
			}
		}
	}
}

func (p *GasMonitor) getPriceInfo() (uint64, error) {
	//TODO update gas price
	return conf.Conf.GasPrice, nil
}
