package perpetualsyncer

import (
	"context"
	"encoding/json"
	"github.com/mcarloai/mai-v3-broker/common/model"
	"github.com/mcarloai/mai-v3-broker/common/utils"
	"github.com/mcarloai/mai-v3-broker/conf"
	"github.com/mcarloai/mai-v3-broker/dao"
	logger "github.com/sirupsen/logrus"
	"net"
	"net/http"
	"time"
)

var transport = &http.Transport{
	DialContext: (&net.Dialer{
		Timeout: 500 * time.Millisecond,
	}).DialContext,
	TLSHandshakeTimeout: 1000 * time.Millisecond,
	MaxIdleConns:        100,
	IdleConnTimeout:     30 * time.Second,
}

type PerpetualSyncer struct {
	ctx            context.Context
	dao            dao.DAO
	httpClient     *utils.HttpClient
	perpetualCache map[string]bool
}

func NewPerpetualSyncer(ctx context.Context, dao dao.DAO) (*PerpetualSyncer, error) {
	transport := &http.Transport{
		DialContext: (&net.Dialer{
			Timeout: 500 * time.Millisecond,
		}).DialContext,
		TLSHandshakeTimeout: 1000 * time.Millisecond,
		MaxIdleConns:        100,
		IdleConnTimeout:     30 * time.Second,
	}
	perpetualSyncer := &PerpetualSyncer{
		ctx:            ctx,
		dao:            dao,
		httpClient:     utils.NewHttpClient(transport),
		perpetualCache: make(map[string]bool),
	}

	err := perpetualSyncer.initPerpetualCache()
	if err != nil {
		return nil, err
	}

	return perpetualSyncer, nil
}

func (p *PerpetualSyncer) initPerpetualCache() error {
	perpetuals, err := p.dao.QueryPerpetuals(false)
	if err != nil {
		return err
	}
	for _, perp := range perpetuals {
		key := perp.LiquidityPoolAddress + "-" + perp.PerpetualIndex
		p.perpetualCache[key] = true
	}
	return nil
}

func (p *PerpetualSyncer) Run() error {
	for {
		select {
		case <-p.ctx.Done():
			return nil
		case <-time.After(10 * time.Second):
			p.syncPerpetual()
		}
	}
}

func (p *PerpetualSyncer) syncPerpetual() {
	var params struct {
		Query string `json:"query"`
	}
	params.Query = `{
		perpetuals {
			index
			liquidityPool {
				id
			}
		}
	}`

	err, code, res := p.httpClient.Post(conf.Conf.Subgraph.URL, nil, params, nil)
	if err != nil || code != 200 {
		logger.Infof("get perpetuals error. err:%s, code:%d", err, code)
		return
	}
	var response struct {
		Data struct {
			Perpetuals []struct {
				Index         string `json:"index"`
				LiquidityPool struct {
					ID string `json:"id"`
				}
			} `json:"perpetuals"`
		} `json:"data"`
	}

	err = json.Unmarshal(res, &response)
	if err != nil {
		logger.Infof("Unmarshal error. err:%s", err)
		return
	}

	for _, perp := range response.Data.Perpetuals {
		key := perp.LiquidityPool.ID + "-" + perp.Index
		if _, ok := p.perpetualCache[key]; ok {
			continue
		}
		newPerpetual := &model.Perpetual{
			PerpetualIndex:       perp.Index,
			LiquidityPoolAddress: perp.LiquidityPool.ID,
		}
		err = p.dao.CreatePerpetual(newPerpetual)
		if err != nil {
			logger.Errorf("CreatePerpetual fail: %s", err)
		}
		p.perpetualCache[key] = true
	}
}
