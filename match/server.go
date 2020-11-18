package match

import (
	"context"
	"github.com/mcarloai/mai-v3-broker/common/chain"
	"github.com/mcarloai/mai-v3-broker/common/message"
	"github.com/mcarloai/mai-v3-broker/common/model"
	"github.com/mcarloai/mai-v3-broker/dao"
	logger "github.com/sirupsen/logrus"
	"sync"
)

type Server struct {
	ctx          context.Context
	wsChan       chan interface{}
	msgChan      chan interface{}
	matchErrChan chan error
	subChans     sync.Map
	subCancel    map[string]context.CancelFunc
	chainCli     chain.ChainClient
	dao          dao.DAO
}

func New(ctx context.Context, cli chain.ChainClient, dao dao.DAO, wsChan, matchChan chan interface{}) *Server {
	return &Server{
		ctx:          ctx,
		wsChan:       wsChan,
		msgChan:      matchChan,
		matchErrChan: make(chan error, 1),
		subCancel:    make(map[string]context.CancelFunc),
		chainCli:     cli,
		dao:          dao,
	}
}

func (s *Server) Start() error {
	logger.Infof("Match start")
	perpetuals, err := s.dao.QueryPerpetuals(true)
	if err != nil {
		return err
	}

	for _, perpetual := range perpetuals {
		s.startSubMatch(perpetual)
	}

	go s.startConsumer()

	select {
	case <-s.ctx.Done():
		logger.Infof("Match Server receive context done")
		return nil
	case err := <-s.matchErrChan:
		logger.Infof("Match Server sub match error:%s", err.Error())
		return err
	}
}

func (s *Server) startConsumer() {
	for {
		select {
		case <-s.ctx.Done():
			logger.Infof("Match Server Consumer Exit")
			return
		case msg, ok := <-s.msgChan:
			if !ok {
				return
			}
			s.parseMessage(msg.(message.MatchMessage))
		}
	}
}

func (s *Server) parseMessage(msg message.MatchMessage) {
	if msg.Type == message.MatchTypeNewPerpetual {
		logger.Infof("Match Server start submatch perpetual %s", msg.PerpetualAddress)
		err := s.newPerpetual(msg.PerpetualAddress)
		if err != nil {
			logger.Infof("Match Server start new perpetual match error:%s", err.Error())
		}
		return
	} else if msg.Type == message.MatchTypePerpetualRollBack {
		logger.Infof("Match Server stop submatch perpetual %s", msg.PerpetualAddress)
		if cancel, ok := s.subCancel[msg.PerpetualAddress]; ok {
			cancel()
			delete(s.subCancel, msg.PerpetualAddress)
		}
		v, ok := s.subChans.Load(msg.PerpetualAddress)
		if ok {
			close(v.(chan interface{}))
			s.subChans.Delete(msg.PerpetualAddress)
		}
		return
	}

	v, ok := s.subChans.Load(msg.PerpetualAddress)
	if !ok {
		logger.Infof("Match Server subchan not fund perpetual address:%s, type:%s", msg.PerpetualAddress, msg.Type)
	}
	msgChan := v.(chan interface{})
	msgChan <- msg
}

func (s *Server) newPerpetual(perpetualAddress string) error {
	perpetual, err := s.dao.GetPerpetualByAddress(perpetualAddress)
	if err != nil {
		return err
	}
	s.startSubMatch(perpetual)
	return nil
}

func (s *Server) startSubMatch(perpetual *model.Perpetual) {
	matchChan := make(chan interface{}, 100)
	ctx, cancel := context.WithCancel(s.ctx)
	m := newMatch(ctx, s.chainCli, s.dao, perpetual, s.wsChan, matchChan)
	go func(m *match) {
		if err := m.run(); err != nil {
			s.matchErrChan <- err
		}
	}(m)
	s.subChans.Store(perpetual.PerpetualAddress, matchChan)
	s.subCancel[perpetual.PerpetualAddress] = cancel
}
