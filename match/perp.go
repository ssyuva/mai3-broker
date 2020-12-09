package match

import (
	"context"
	"fmt"
	"github.com/mcarloai/mai-v3-broker/common/model"
	"github.com/mcarloai/mai-v3-broker/conf"
	logger "github.com/sirupsen/logrus"
	"time"
)

type PerpetualContext struct {
	GovParams   *model.GovParams
	PerpStorage *model.PerpetualStorage
	Amm         *model.AccountStorage
}

func (m *match) getPerpetualContext() (*PerpetualContext, error) {
	ctxTimeout, ctxTimeoutCancel := context.WithTimeout(m.ctx, conf.Conf.BlockChain.Timeout.Duration)
	defer ctxTimeoutCancel()
	storage, err := m.chainCli.GetPerpetualStorage(ctxTimeout, m.perpetual.PerpetualAddress)
	if err != nil {
		logger.Errorf("GetPerpetualStorage fail! err:%s", err.Error())
		return nil, fmt.Errorf("getPerpetualContext:%w", err)
	}

	gov, err := m.chainCli.GetPerpetualGovParams(ctxTimeout, m.perpetual.PerpetualAddress)
	if err != nil {
		logger.Errorf("GetPerpetualGovParams fail! err:%s", err.Error())
		return nil, fmt.Errorf("getPerpetualContext:%w", err)
	}

	amm, err := m.chainCli.GetMarginAccount(ctxTimeout, m.perpetual.PerpetualAddress, m.perpetual.PerpetualAddress)
	if err != nil {
		logger.Errorf("GetMarginAccount fail! err:%s", err.Error())
		return nil, fmt.Errorf("getPerpetualContext:%w", err)
	}
	return &PerpetualContext{
		GovParams:   gov,
		PerpStorage: storage,
		Amm:         amm,
	}, nil
}

func (m *match) updatePerpetualContext() {
	for {
		select {
		case <-m.ctx.Done():
			return
		case <-time.After(10 * time.Second):
			perpContext, err := m.getPerpetualContext()
			if err != nil {
				logger.Errorf("updatePerpetualContext: %s", err)
			} else {
				m.mu.Lock()
				defer m.mu.Unlock()
				m.perpetualContext = perpContext
			}
		}
	}
}
