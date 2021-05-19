package gasmonitor

import (
	"context"
	"time"

	"github.com/mcarloai/mai-v3-broker/common/chain"
	"github.com/mcarloai/mai-v3-broker/conf"
	"github.com/shopspring/decimal"
	logger "github.com/sirupsen/logrus"
)

type GasMonitor struct {
	ctx      context.Context
	chainCli chain.ChainClient
	gasPrice decimal.Decimal
}

var gwei, _ = decimal.NewFromString("1000000000")

func NewGasMonitor(ctx context.Context, cli chain.ChainClient) *GasMonitor {
	gasMonitor := &GasMonitor{
		ctx:      ctx,
		chainCli: cli,
	}

	return gasMonitor
}

// GetGasPrice return gas in Gwei
func (p *GasMonitor) GetGasPrice() decimal.Decimal {
	return p.gasPrice
}

// GasPriceGwei return gas in eth decimal
func (p *GasMonitor) GasPriceGwei() decimal.Decimal {
	return p.gasPrice.Mul(gwei)
}

func (p *GasMonitor) Run() error {
	logger.Infof("gas price monitor start")
	for {
		select {
		case <-p.ctx.Done():
			logger.Infof("gas price monitor end")
			return nil
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

func (p *GasMonitor) getPriceInfo() (decimal.Decimal, error) {
	//TODO update gas price
	if conf.Conf.GasArbEnable {
		res, err := p.chainCli.GetGasPrice(p.ctx, conf.Conf.GasArbAddress)
		if err != nil {
			return decimal.Zero, err
		}
		return res[0], nil
	} else {
		return decimal.NewFromInt(int64(conf.Conf.GasPrice)), nil
	}
}
