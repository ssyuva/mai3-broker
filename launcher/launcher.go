package launcher

import (
	"context"
	"fmt"
	"github.com/mcarloai/mai-v3-broker/common/chain"
	"github.com/mcarloai/mai-v3-broker/common/mai3"
	"github.com/mcarloai/mai-v3-broker/common/mai3/utils"
	"github.com/mcarloai/mai-v3-broker/common/model"
	"github.com/mcarloai/mai-v3-broker/dao"
	"github.com/shopspring/decimal"
	logger "github.com/sirupsen/logrus"
	"math/big"
	"time"
)

type Launcher struct {
	ctx       context.Context
	addr      string
	dao       dao.DAO
	chainCli  chain.ChainClient
	wsChan    chan interface{}
	matchChan chan interface{}
	execChan  chan interface{}
	syncChan  chan interface{}
}

func NewLaunch(ctx context.Context, addr string, dao dao.DAO, chainCli chain.ChainClient, wsChan, matchChan chan interface{}) *Launcher {
	return &Launcher{
		ctx:       ctx,
		addr:      addr,
		dao:       dao,
		chainCli:  chainCli,
		wsChan:    wsChan,
		matchChan: matchChan,
		execChan:  make(chan interface{}, 100),
		syncChan:  make(chan interface{}, 100),
	}
}

func (l *Launcher) Start() error {
	logger.Infof("Launcher start")
	//TODO private key aes crypto
	err := l.chainCli.AddAccount("")
	if err != nil {
		return err
	}
	// start syncer for sync pending transactions
	syncer := NewSyncer(l.ctx, l.dao, l.chainCli, l.matchChan, l.syncChan)
	go syncer.Run()

	// start executor for execute launch transactions
	executor := NewExecutor(l.ctx, l.dao, l.chainCli, l.execChan)
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

func (l *Launcher) checkMatchTransaction() error {
	l.dao.ForUpdate()
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
	//TODO gas price
	gasRewards := make([]*big.Int, len(matchTx.MatchResult.MatchItems))
	for _, item := range matchTx.MatchResult.MatchItems {
		param, err := getWalletOrderParam(item.Order)
		if err != nil {
			return err
		}
		orderParams = append(orderParams, param)
		amount := item.Amount
		if item.Order.Side == model.SideSell {
			amount = amount.Neg()
		}
		matchAmounts = append(matchAmounts, amount)
		gasRewards = append(gasRewards, big.NewInt(1000000))
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
		ToAddress:   l.addr,
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
	signature, err := utils.Hex2Bytes(order.Signature)
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
	amount := order.Amount
	if order.Side == model.SideSell {
		amount = amount.Neg()
	}
	param := &model.WalletOrderParam{
		Trader:    order.TraderAddress,
		Broker:    order.BrokerAddress,
		Relayer:   order.RelayerAddress,
		Perpetual: order.PerpetualAddress,
		Referrer:  order.ReferrerAddress,
		Price:     order.Price,
		Amount:    amount,
		OrderData: orderData,
		ChainID:   uint64(order.ChainID),
		Signature: signature,
	}

	return param, nil
}
