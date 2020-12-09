package pricemonitor

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/mcarloai/mai-v3-broker/conf"
	logger "github.com/sirupsen/logrus"
	"net/http"
	"time"
)

type PriceTable struct {
	Fast    uint64 `json:"fast"`
	Fastest uint64 `json:"fastest"`
	SafeLow uint64 `json:"safeLow"`
	Average uint64 `json:"average"`
}

func (pt *PriceTable) PriceByName(name string) uint64 {
	switch name {
	case "fast":
		return pt.Fast
	case "fastest":
		return pt.Fastest
	case "safeLow":
		return pt.SafeLow
	case "average":
		return pt.Average
	}
	return pt.Average
}

type PriceMonitor struct {
	ctx context.Context
	pt  *PriceTable
}

func NewPriceMonitor(ctx context.Context) *PriceMonitor {
	priceMonitor := &PriceMonitor{
		ctx: ctx,
	}

	go priceMonitor.run()
	return priceMonitor
}

func (p *PriceMonitor) GetGasPrice() uint64 {
	return p.pt.PriceByName(conf.Conf.GasStation.GasLevel)
}

func (p *PriceMonitor) run() {
	for {
		select {
		case <-p.ctx.Done():
			return
		case <-time.After(10 * time.Second):
			priceTable, err := p.getPriceInfo()
			if err != nil {
				logger.Errorf("fail to retrieve lastest gas price info:%w", err)
			} else {
				p.pt = priceTable
			}
		}
	}
}

func (p *PriceMonitor) getPriceInfo() (*PriceTable, error) {
	httpCli := &http.Client{Timeout: conf.Conf.GasStation.Timeout.Duration}
	req, err := http.NewRequestWithContext(p.ctx, "GET", conf.Conf.GasStation.URL, nil)
	if err != nil {
		return nil, fmt.Errorf("build request failed:%w", err)
	}
	rsp, err := httpCli.Do(req)
	if err != nil {
		return nil, fmt.Errorf("get from target failed:%w", err)
	}
	defer rsp.Body.Close()
	content := new(PriceTable)
	err = json.NewDecoder(rsp.Body).Decode(content)
	if err != nil {
		return nil, fmt.Errorf("decode http content failed:%w", err)
	}
	content.Average = content.Average / 10
	content.Fast = content.Fast / 10
	content.Fastest = content.Fastest / 10
	content.SafeLow = content.SafeLow / 10

	logger.Infof("price => %+v", content)

	return content, nil
}
