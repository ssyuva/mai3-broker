package match

import (
	"context"
	"fmt"
	"github.com/mcarloai/mai-v3-broker/common/chain"
	"github.com/mcarloai/mai-v3-broker/common/model"
	"github.com/mcarloai/mai-v3-broker/dao"
	"github.com/mcarloai/mai-v3-broker/gasmonitor"
	"github.com/shopspring/decimal"
	logger "github.com/sirupsen/logrus"
	"sync"
)

type Server struct {
	ctx             context.Context
	mu              sync.Mutex
	matchHandlerMap map[string]*match
	wsChan          chan interface{}
	matchErrChan    chan error
	chainCli        chain.ChainClient
	gasMonitor      *gasmonitor.GasMonitor
	dao             dao.DAO
}

func New(ctx context.Context, cli chain.ChainClient, dao dao.DAO, wsChan chan interface{}, gm *gasmonitor.GasMonitor) (*Server, error) {
	server := &Server{
		ctx:             ctx,
		wsChan:          wsChan,
		matchHandlerMap: make(map[string]*match),
		matchErrChan:    make(chan error, 1),
		chainCli:        cli,
		gasMonitor:      gm,
		dao:             dao,
	}
	perpetuals, err := dao.QueryPerpetuals(true)
	if err != nil {
		logger.Errorf("New Match Server QueryPerpetuals:%w", err)
		return nil, err
	}

	for _, perpetual := range perpetuals {
		if err := server.newMatch(perpetual); err != nil {
			logger.Errorf("New SubMatch Server newMatch:%w", err)
			return nil, err
		}
	}
	return server, nil
}

func (s *Server) newMatch(perpetual *model.Perpetual) error {
	m, err := newMatch(s.ctx, s.chainCli, s.dao, perpetual, s.wsChan, s.gasMonitor)
	if err != nil {
		return err
	}
	s.setMatchHandler(perpetual.PerpetualAddress, m)
	return nil
}

func (s *Server) NewOrder(order *model.Order) string {
	handler := s.getMatchHandler(order.PerpetualAddress)
	if handler == nil {
		perpetual, err := s.dao.GetPerpetualByAddress(order.PerpetualAddress, true)
		if err != nil {
			return model.MatchInternalErrorID
		}
		err = s.newMatch(perpetual)
		return model.MatchInternalErrorID
	}
	return handler.NewOrder(order)
}

func (s *Server) CancelOrder(perpetualAddress, orderHash string) error {
	handler := s.getMatchHandler(perpetualAddress)
	if handler == nil {
		return fmt.Errorf("CancelOrder error: perpetual[%s] is not open.", perpetualAddress)
	}
	return handler.CancelOrder(orderHash, model.CancelReasonUserCancel, true, decimal.Zero)
}

func (s *Server) CancelAllOrders(perpetualAddress, trader string) error {
	handler := s.getMatchHandler(perpetualAddress)
	if handler == nil {
		return fmt.Errorf("CancelAllOrders error: perpetual[%s] is not open.", perpetualAddress)
	}
	return handler.CancelAllOrders(perpetualAddress, trader)
}

func (s *Server) ClosePerpetual(perpetualAddress string) error {
	handler := s.getMatchHandler(perpetualAddress)
	if handler == nil {
		return fmt.Errorf("ClosePerpetual error: perpetual[%s] is not open.", perpetualAddress)
	}
	perpetual, err := s.dao.GetPerpetualByAddress(perpetualAddress, true)
	if err != nil {
		return err
	}
	perpetual.IsPublished = false
	err = s.dao.UpdatePerpetual(perpetual)
	if err != nil {
		return err
	}
	handler.Close()
	return s.deleteMatchHandler(perpetualAddress)
}

func (s *Server) UpdateOrdersStatus(txID string, status model.TransactionStatus, transactionHash, blockHash string, blockNumber, blockTime uint64) error {
	matchTx, err := s.dao.GetMatchTransaction(txID)
	if err != nil {
		return err
	}
	handler := s.getMatchHandler(matchTx.PerpetualAddress)
	if handler == nil {
		return fmt.Errorf("UpdateOrdersStatus error: perpetual[%s] is not open.", matchTx.PerpetualAddress)
	}
	err = handler.UpdateOrdersStatus(txID, status, transactionHash, blockHash, blockNumber, blockTime)
	return err
}

func (s *Server) RollbackOrdersStatus(txID string, status model.TransactionStatus, transactionHash, blockHash string, blockNumber, blockTime uint64) error {
	matchTx, err := s.dao.GetMatchTransaction(txID)
	if err != nil {
		return err
	}
	handler := s.getMatchHandler(matchTx.PerpetualAddress)
	if handler == nil {
		return fmt.Errorf("RollBackOrdersStatus error: perpetual[%s] is not open.", matchTx.PerpetualAddress)
	}
	err = handler.RollbackOrdersStatus(txID, status, transactionHash, blockHash, blockNumber, blockTime)
	return err
}

func (s *Server) getMatchHandler(perpetualAddress string) *match {
	s.mu.Lock()
	defer s.mu.Unlock()
	h, ok := s.matchHandlerMap[perpetualAddress]
	if !ok {
		return nil
	}
	return h
}

func (s *Server) setMatchHandler(perpetualAddress string, h *match) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	_, ok := s.matchHandlerMap[perpetualAddress]
	if ok {
		return fmt.Errorf("setMatchHandler error:perpetualAddress[%s] exists", perpetualAddress)
	}
	s.matchHandlerMap[perpetualAddress] = h
	return nil
}

func (s *Server) deleteMatchHandler(perpetualAddress string) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	_, ok := s.matchHandlerMap[perpetualAddress]
	if ok {
		delete(s.matchHandlerMap, perpetualAddress)
	} else {
		return fmt.Errorf("deleteMatchHandler:perpetualAddress[%s] do not exists", perpetualAddress)
	}
	return nil
}
