package match

import (
	"context"
	"github.com/mcarloai/mai-v3-broker/common/message"
	"github.com/mcarloai/mai-v3-broker/common/model"
	"github.com/mcarloai/mai-v3-broker/dao"
	"github.com/micro/go-micro/v2/logger"
	"sync"
)

type Server struct {
	ctx          context.Context
	wsChan       chan interface{}
	msgChan      chan interface{}
	matchErrChan chan error
	subChans     sync.Map
	dao          dao.DAO
}

func New(ctx context.Context, dao dao.DAO, wsChan, matchChan chan interface{}) *Server {
	return &Server{
		ctx:          ctx,
		wsChan:       wsChan,
		msgChan:      matchChan,
		matchErrChan: make(chan error, 1),
		dao:          dao,
	}
}

func (s *Server) Start() error {
	perpetuals, err := s.dao.QueryPerpetuals(true)
	if err != nil {
		return err
	}

	for _, perpetual := range perpetuals {
		s.startMatch(perpetual)
	}

	go s.startConsumer()

	select {
	case <-s.ctx.Done():
		logger.Infof("Match Server recieve context done")
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
		err := s.newPerpetual(msg.PerpetualAddress)
		if err != nil {
			logger.Infof("Match Server start new perpetual match error:%s", err.Error())
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
	s.startMatch(perpetual)
	return nil
}

func (s *Server) startMatch(perpetual *model.Perpetual) {
	matchChan := make(chan interface{}, 100)
	m := newMatch(s.ctx, s.dao, perpetual, s.wsChan, matchChan)
	go func(m *match) {
		if err := m.run(); err != nil {
			s.matchErrChan <- err
		}
	}(m)
	s.subChans.Store(perpetual.PerpetualAddress, matchChan)
}
