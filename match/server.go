package match

import (
	"context"
	"github.com/mcarloai/mai-v3-broker/common/message"
	"github.com/mcarloai/mai-v3-broker/dao"
	"github.com/micro/go-micro/v2/logger"
	"sync"
)

type Server struct {
	ctx      context.Context
	wsChan   chan interface{}
	msgChan  chan interface{}
	subChans sync.Map
	dao      dao.DAO
}

func New(ctx context.Context, dao dao.DAO, wsChan, matchChan chan interface{}) *Server {
	return &Server{
		ctx:     ctx,
		wsChan:  wsChan,
		msgChan: matchChan,
		dao:     dao,
	}
}

func (s *Server) Start() error {
	perpetuals, err := s.dao.QueryPerpetuals(true)
	if err != nil {
		return err
	}

	matchErrChan := make(chan error, 1)

	for _, perpetual := range perpetuals {
		matchChan := make(chan interface{}, 100)
		m := newMatch(s.ctx, s.dao, perpetual, s.wsChan, matchChan)
		go func(m *match) {
			if err := m.run(); err != nil {
				matchErrChan <- err
			}
		}(m)
		s.subChans.Store(perpetual.PerpetualAddress, matchChan)
	}

	go s.startConsumer()

	select {
	case <-s.ctx.Done():
		logger.Infof("Match Server recieve context done")
		return nil
	case err := <-matchErrChan:
		logger.Infof("Match Server sub match error:%s", err.Error())
		return err
	}
}

func (s *Server) startConsumer() {
	for {
		select {
		case <-s.ctx.Done():
			logger.Infof("Match Consumer Exit")
			return
		case msg, ok := <-s.msgChan:
			if !ok {
				return
			}
			wsMsg := msg.(message.MatchMessage)
			if wsMsg.Type == message.MatchTypeNewPerpetual {

			}
		}
	}
}
