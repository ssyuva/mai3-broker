package syncer

import (
	"fmt"
	"github.com/mcarloai/mai-v3-broker/common/chain"
	"github.com/mcarloai/mai-v3-broker/common/message"
	"github.com/mcarloai/mai-v3-broker/common/model"
	"github.com/mcarloai/mai-v3-broker/dao"
	logger "github.com/sirupsen/logrus"
)

type CreatePerpetualSyncer struct {
	factoryAddress string
	chainCli       chain.ChainClient
	matchChan      chan interface{}
}

func NewCreatePerpetualSyncer(factoryAddress string, chainCli chain.ChainClient, matchChan chan interface{}) *CreatePerpetualSyncer {
	return &CreatePerpetualSyncer{
		factoryAddress: factoryAddress,
		chainCli:       chainCli,
		matchChan:      matchChan,
	}
}

func (c *CreatePerpetualSyncer) Rollback(syncCtx *SyncBlockContext) error {
	perpetuals, err := syncCtx.Dao.RollbackPerpetual(syncCtx.RollbackBeginHeight, syncCtx.RollbackEndHeight)
	if err != nil {
		return fmt.Errorf("watchCreatePerpetual.rollback failed:%w", err)
	}

	for _, perpetual := range perpetuals {
		matchMsg := message.MatchMessage{
			Type:             message.MatchTypePerpetualRollBack,
			PerpetualAddress: perpetual.PerpetualAddress,
		}
		c.matchChan <- matchMsg
	}

	return nil
}

func (c *CreatePerpetualSyncer) Forward(syncCtx *SyncBlockContext) error {
	perpetualEvents, err := c.chainCli.FilterCreatePerpetual(syncCtx.Context, c.factoryAddress, uint64(syncCtx.RollbackBeginHeight), uint64(syncCtx.LatestBlockNumber))
	if err != nil {
		return fmt.Errorf("watcher filter create perpetual event failed:%w", err)
	}
	for _, event := range perpetualEvents {
		_, err := syncCtx.Dao.GetPerpetualByAddress(event.PerpetualAddress)
		if err != nil && !dao.IsRecordNotFound(err) {
			logger.Errorf("watcher GetPerpetualByAddress failed:%w", err)
			continue
		} else if err == nil {
			logger.Errorf("watcher perpetual already exists:%s", event.PerpetualAddress)
			continue
		}
		//TODO CreatePerpetual event
		dbPerpetual := &model.Perpetual{
			PerpetualAddress: event.PerpetualAddress,
			OracleAddress:    event.OracleAddress,
			BlockNumber:      event.BlockNumber,
		}
		err = syncCtx.Dao.CreatePerpetual(dbPerpetual)
		if err != nil {
			logger.Errorf("watcher CreatePerpetual failed:%w", err)
			continue
		}

		matchMsg := message.MatchMessage{
			Type:             message.MatchTypeNewPerpetual,
			PerpetualAddress: event.PerpetualAddress,
		}
		c.matchChan <- matchMsg
	}
	return nil
}
