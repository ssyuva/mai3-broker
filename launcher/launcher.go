package launcher

import (
	"context"
	"fmt"
	"github.com/mcarloai/mai-v3-broker/common/chain"
	"github.com/mcarloai/mai-v3-broker/common/mai3"
	mai3Utils "github.com/mcarloai/mai-v3-broker/common/mai3/utils"
	"github.com/mcarloai/mai-v3-broker/common/model"
	"github.com/mcarloai/mai-v3-broker/conf"
	"github.com/mcarloai/mai-v3-broker/dao"
	"github.com/mcarloai/mai-v3-broker/gasmonitor"
	"github.com/mcarloai/mai-v3-broker/match"
	"github.com/shopspring/decimal"
	logger "github.com/sirupsen/logrus"
	"math/big"
	"time"
)

type Launcher struct {
	ctx        context.Context
	dao        dao.DAO
	chainCli   chain.ChainClient
	gasMonitor *gasmonitor.GasMonitor
	match      *match.Server
	execChan   chan interface{}
	syncChan   chan interface{}
}

func NewLaunch(ctx context.Context, dao dao.DAO, chainCli chain.ChainClient, match *match.Server, gm *gasmonitor.GasMonitor) *Launcher {
	return &Launcher{
		ctx:        ctx,
		dao:        dao,
		chainCli:   chainCli,
		gasMonitor: gm,
		match:      match,
		execChan:   make(chan interface{}, 100),
		syncChan:   make(chan interface{}, 100),
	}
}

func (l *Launcher) Start() error {
	logger.Infof("Launcher start")
	err := l.reloadAccount()
	if err != nil {
		logger.Errorf("reload account error %s", err)
		return err
	}
	// start syncer for sync pending transactions
	syncer := NewSyncer(l.ctx, l.dao, l.chainCli, l.syncChan, l.match)
	go syncer.Run()

	// start executor for execute launch transactions
	executor := NewExecutor(l.ctx, l.dao, l.chainCli, l.execChan, l.gasMonitor)
	go executor.Run()

	for {
		select {
		case <-l.ctx.Done():
			logger.Infof("Launcher receive context done")
			return nil
		case <-time.After(5 * time.Second):
			err := l.checkMatchTransaction()
			if err != nil {
				logger.Errorf("excuteMatchTransaction failed! err:%v", err.Error())
			}
		}
	}
}

func (l *Launcher) reloadAccount() error {
	stores, err := l.dao.List("keystore")
	if err != nil {
		return err
	}
	for _, s := range stores {
		p, err := l.chainCli.DecryptKey(s.Value, conf.Conf.BlockChain.Password)
		if err != nil {
			return err
		}
		err = l.chainCli.AddAccount(p)
		if err != nil {
			return err
		}
	}
	return nil
}

func (l *Launcher) ImportPrivateKey(pk string) (string, error) {
	p, address, err := l.chainCli.HexToPrivate(pk)
	if err != nil {
		return address, err
	}

	b, err := l.chainCli.EncryptKey(p, conf.Conf.BlockChain.Password)
	if err != nil {
		return address, fmt.Errorf("fail to encrypt key", err)
	}
	err = l.dao.Put(&model.KVStore{
		Key:      address,
		Category: "keystore",
		Value:    b,
	})
	if err != nil {
		return address, err
	}

	err = l.chainCli.AddAccount(p)
	return address, err
}

func (l *Launcher) checkMatchTransaction() error {
	transactions, err := l.dao.QueryMatchTransaction("", []model.TransactionStatus{model.TransactionStatusInit})
	if err != nil {
		return fmt.Errorf("QueryUnconfirmedTransactions failed error:%w", err)
	}
	for _, transaction := range transactions {
		if err = l.dao.LoadMatchOrders(transaction.MatchResult.MatchItems); err != nil {
			return fmt.Errorf("LoadMatchOrders:%w", err)
		}
	}

	for _, tx := range transactions {
		if err = l.createLaunchTransaction(tx); err != nil {
			return err
		}
	}
	return nil
}

func (l *Launcher) createLaunchTransaction(matchTx *model.MatchTransaction) error {
	_, err := l.dao.GetTxByID(matchTx.ID)
	if !dao.IsRecordNotFound(err) {
		return fmt.Errorf("Transaction already launched ID:%s", matchTx.ID)
	}

	orderParams := make([]*model.WalletOrderParam, len(matchTx.MatchResult.MatchItems))
	matchAmounts := make([]decimal.Decimal, len(matchTx.MatchResult.MatchItems))
	gasRewards := make([]*big.Int, len(matchTx.MatchResult.MatchItems))
	for _, item := range matchTx.MatchResult.MatchItems {
		param, err := getWalletOrderParam(item.Order)
		if err != nil {
			return err
		}
		orderParams = append(orderParams, param)
		matchAmounts = append(matchAmounts, item.Amount)
		gasReward := l.gasMonitor.GetGasPrice() * 1e9 * conf.Conf.GasStation.GasLimit
		gasRewards = append(gasRewards, big.NewInt(int64(gasReward)))
	}
	inputs, err := l.chainCli.BatchTradeDataPack(orderParams, matchAmounts, gasRewards)
	if err != nil {
		return err
	}

	signAccount, err := l.chainCli.GetSignAccount()
	if err != nil {
		return err
	}

	tx := &model.LaunchTransaction{
		TxID:        matchTx.ID,
		Type:        model.TxNormal,
		FromAddress: signAccount,
		ToAddress:   matchTx.BrokerAddress,
		Inputs:      inputs,
		Status:      model.TxInitial,
		CommitTime:  time.Now(),
	}

	err = l.dao.Transaction(func(dao dao.DAO) error {
		if err := dao.CreateTx(tx); err != nil {
			return fmt.Errorf("create transaction failed error:%w", err)
		}
		matchTx.Status = model.TransactionStatusPending
		if err := dao.UpdateMatchTransaction(matchTx); err != nil {
			return fmt.Errorf("update match transaction status failed error:%w", err)
		}
		return nil
	})

	if err == nil {
		l.execChan <- nil
	}
	return err
}

func getWalletOrderParam(order *model.Order) (*model.WalletOrderParam, error) {
	if order == nil {
		return nil, fmt.Errorf("getWalletOrderParam:nil order")
	}
	signature, err := mai3Utils.Hex2Bytes(order.Signature)
	if err != nil {
		return nil, fmt.Errorf("getWalletOrderParam:%w", err)
	}
	orderData, err := mai3.GetOrderData(
		order.ExpiresAt.UTC().Unix(),
		order.Version,
		int8(order.Type),
		order.IsCloseOnly,
		order.Salt,
	)
	param := &model.WalletOrderParam{
		Trader:    order.TraderAddress,
		Broker:    order.BrokerAddress,
		Relayer:   order.RelayerAddress,
		Perpetual: order.PerpetualAddress,
		Referrer:  order.ReferrerAddress,
		Price:     order.Price,
		Amount:    order.Amount,
		OrderData: orderData,
		ChainID:   uint64(order.ChainID),
		Signature: signature,
	}

	return param, nil
}
