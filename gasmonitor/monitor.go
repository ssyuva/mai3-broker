package gasmonitor

import (
	"context"
	"github.com/mcarloai/mai-v3-broker/conf"
	logger "github.com/sirupsen/logrus"
	"time"
)

type GasMonitor struct {
	ctx      context.Context
	gasPrice uint64
}

func NewGasMonitor(ctx context.Context) *GasMonitor {
	gasMonitor := &GasMonitor{
		ctx: ctx,
	}

	go gasMonitor.run()
	return gasMonitor
}

func (p *GasMonitor) GetGasPrice() uint64 {
	return p.gasPrice
}

func (p *GasMonitor) run() {
	for {
		select {
		case <-p.ctx.Done():
			return
		case <-time.After(10 * time.Second):
			gasPrice, err := p.getPriceInfo()
			if err != nil {
				logger.Errorf("fail to retrieve lastest gas price info:%w", err)
			} else {
				p.gasPrice = gasPrice
			}
		}
	}
}

func (p *GasMonitor) getPriceInfo() (uint64, error) {
	//TODO update gas price
	return conf.Conf.GasStation.GasPrice, nil
}
