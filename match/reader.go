package match

import (
	"context"
	"github.com/mcarloai/mai-v3-broker/common/model"
	"github.com/mcarloai/mai-v3-broker/conf"
	logger "github.com/sirupsen/logrus"
	"time"
)

type PerpetualContext struct {
	GovParams   *model.GovParams
	PerpStorage *model.PerpetualStorage
	AMM         *model.AccountStorage
}

func (m *match) updatePerpContext() error {
	m.mu.Lock()
	defer m.mu.Unlock()

	ctxTimeout, ctxTimeoutCancel := context.WithTimeout(m.ctx, conf.Conf.BlockChain.Timeout.Duration)
	defer ctxTimeoutCancel()
	//TODO
	storage, err := m.chainCli.GetPerpetualStorage(ctxTimeout, m.perpetual.LiquidityPoolAddress)
	if err != nil {
		logger.Errorf("GetPerpetualStorage fail! err:%s", err.Error())
		return err
	}
	m.perpetualContext.PerpStorage = storage

	gov, err := m.chainCli.GetPerpetualGovParams(ctxTimeout, m.perpetual.LiquidityPoolAddress)
	if err != nil {
		logger.Errorf("GetPerpetualGovParams fail! err:%s", err.Error())
		return err
	}
	m.perpetualContext.GovParams = gov

	amm, err := m.chainCli.GetMarginAccount(ctxTimeout, m.perpetual.LiquidityPoolAddress, m.perpetual.LiquidityPoolAddress)
	if err != nil {
		logger.Errorf("GetMarginAccount fail! err:%s", err.Error())
		return err
	}

	m.perpetualContext.AMM = amm

	return nil
}

func (m *match) checkPerpetualContext() {
	for {
		select {
		case <-m.ctx.Done():
			return
		case <-time.After(10 * time.Second):
			m.updatePerpContext()
		}
	}
}
