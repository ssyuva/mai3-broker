package syncer

import (
	"fmt"
	"github.com/mcarloai/mai-v3-broker/common/chain"
	"github.com/mcarloai/mai-v3-broker/common/model"
	"github.com/mcarloai/mai-v3-broker/dao"
	logger "github.com/sirupsen/logrus"
)

type CreatePerpetualSyncer struct {
	factoryAddress string
	chainCli       chain.ChainClient
}

func NewCreatePerpetualSyncer(factoryAddress string, chainCli chain.ChainClient) *CreatePerpetualSyncer {
	return &CreatePerpetualSyncer{
		factoryAddress: factoryAddress,
		chainCli:       chainCli,
	}
}

func (c *CreatePerpetualSyncer) Rollback(syncCtx *SyncBlockContext) error {
	_, err := syncCtx.Dao.RollbackPerpetual(syncCtx.RollbackBeginHeight, syncCtx.RollbackEndHeight)
	if err != nil {
		return fmt.Errorf("watchCreatePerpetual.rollback failed:%w", err)
	}
	return nil
}

func (c *CreatePerpetualSyncer) Forward(syncCtx *SyncBlockContext) error {
	perpetualEvents, err := c.chainCli.FilterCreatePerpetual(syncCtx.Context, c.factoryAddress, uint64(syncCtx.RollbackBeginHeight), uint64(syncCtx.LatestBlockNumber))
	if err != nil {
		return fmt.Errorf("watcher filter create perpetual event failed:%w", err)
	}
	for _, event := range perpetualEvents {
		_, err := syncCtx.Dao.GetPerpetualByAddress(event.PerpetualAddress, true)
		if err != nil && !dao.IsRecordNotFound(err) {
			logger.Errorf("watcher GetPerpetualByAddress failed:%w", err)
			continue
		} else if err == nil {
			logger.Errorf("watcher perpetual already exists:%s", event.PerpetualAddress)
			continue
		}
		symbol, err := c.chainCli.GetTokenSymbol(syncCtx.Context, event.CollateralAddress)
		if err != nil {
			logger.Errorf("watcher GetTokenSymbol failed:%w", err)
			continue
		}

		dbPerpetual := &model.Perpetual{
			PerpetualAddress:  event.PerpetualAddress,
			GovernorAddress:   event.GovernorAddress,
			ShareToken:        event.ShareToken,
			OperatorAddress:   event.OperatorAddress,
			CollateralAddress: event.CollateralAddress,
			CollateralSymbol:  symbol,
			OracleAddress:     event.OracleAddress,
			BlockNumber:       event.BlockNumber,
			IsPublished:       true,
		}
		err = syncCtx.Dao.CreatePerpetual(dbPerpetual)
		if err != nil {
			logger.Errorf("watcher CreatePerpetual failed:%w", err)
			continue
		}
	}
	return nil
}
