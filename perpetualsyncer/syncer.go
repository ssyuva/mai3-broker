package perpetualsyncer

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/mcarloai/mai-v3-broker/common/model"
	"github.com/mcarloai/mai-v3-broker/common/utils"
	"github.com/mcarloai/mai-v3-broker/conf"
	"github.com/mcarloai/mai-v3-broker/dao"
	logger "github.com/sirupsen/logrus"
	"net"
	"net/http"
	"strconv"
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
	ctx               context.Context
	dbDao             dao.DAO
	httpClient        *utils.HttpClient
	syncedBlockNumber int64
}

func NewPerpetualSyncer(ctx context.Context, dbDao dao.DAO) (*PerpetualSyncer, error) {
	transport := &http.Transport{
		DialContext: (&net.Dialer{
			Timeout: 500 * time.Millisecond,
		}).DialContext,
		TLSHandshakeTimeout: 1000 * time.Millisecond,
		MaxIdleConns:        100,
		IdleConnTimeout:     30 * time.Second,
	}
	perpetualSyncer := &PerpetualSyncer{
		ctx:        ctx,
		dbDao:      dbDao,
		httpClient: utils.NewHttpClient(transport),
	}

	blockNumber, err := perpetualSyncer.dbDao.GetPerpetualSyncedBlockNumber()
	if err != nil && !dao.IsRecordNotFound(err) {
		return nil, err
	}
	perpetualSyncer.syncedBlockNumber = blockNumber

	return perpetualSyncer, nil
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
	queryFormat := `{
		perpetuals(where:{createdAtBlockNumber_gt:%d} orderBy:createdAtBlockNumber) {
			index
			symbol
			collateralName
			operatorAddress
			liquidityPool {
				id
			}
			createdAtBlockNumber
		}
	}`
	params.Query = fmt.Sprintf(queryFormat, p.syncedBlockNumber)

	err, code, res := p.httpClient.Post(conf.Conf.SubgraphURL, nil, params, nil)
	if err != nil || code != 200 {
		logger.Infof("get perpetuals error. err:%s, code:%d", err, code)
		return
	}
	var response struct {
		Data struct {
			Perpetuals []struct {
				Index           string `json:"index"`
				Symbol          string `json:"symbol"`
				CollateralName  string `json:"collateralName"`
				OperatorAddress string `json:"operatorAddress"`
				LiquidityPool   struct {
					ID string `json:"id"`
				}
				CreatedAtBlockNumber string `json:"createdAtBlockNumber"`
			} `json:"perpetuals"`
		} `json:"data"`
	}

	err = json.Unmarshal(res, &response)
	if err != nil {
		logger.Infof("Unmarshal error. err:%s", err)
		return
	}

	for _, perp := range response.Data.Perpetuals {
		index, err := strconv.ParseInt(perp.Index, 10, 64)
		if err != nil {
			logger.Errorf("parse perpetual index fail: %s", err)
			continue
		}
		blockNumber, err := strconv.ParseInt(perp.CreatedAtBlockNumber, 10, 64)
		if err != nil {
			logger.Errorf("parse perpetual createdAt blockNumber fail: %s", err)
			continue
		}
		newPerpetual := &model.Perpetual{
			PerpetualIndex:       index,
			LiquidityPoolAddress: perp.LiquidityPool.ID,
			Symbol:               perp.Symbol,
			CollateralSymbol:     perp.CollateralName,
			OperatorAddress:      perp.OperatorAddress,
			IsPublished:          true,
			BlockNumber:          blockNumber,
		}
		err = p.dbDao.CreatePerpetual(newPerpetual)
		if err != nil {
			logger.Errorf("CreatePerpetual fail: %s", err)
		}
		if blockNumber > p.syncedBlockNumber {
			p.syncedBlockNumber = blockNumber
		}
	}
}
