package api

import (
	"github.com/mcarloai/mai-v3-broker/dao"
)

func (s *Server) GetPerpetual(p Param) (interface{}, error) {
	params := p.(*GetPerpetualReq)
	perpetual, err := s.dao.GetPerpetualByAddress(params.PerpetualAddress)
	if err != nil {
		if dao.IsRecordNotFound(err) {
			return nil, PerpetualNotFoundError(params.PerpetualAddress)
		}
		return nil, InternalError(err)
	}
	return &GetPerpetualResp{
		Perpetual: perpetual,
	}, nil
}

func (s *Server) GetBrokerRelay(p Param) (interface{}, error) {
	params := p.(*GetBrokerRelayReq)
	perpetual, err := s.dao.GetPerpetualByAddress(params.PerpetualAddress)
	if err != nil {
		if dao.IsRecordNotFound(err) {
			return nil, PerpetualNotFoundError(params.PerpetualAddress)
		}
		return nil, InternalError(err)
	}

	//TODO relayer address
	return &GetBrokerRelayResp{
		BrokerAddress:  perpetual.BrokerAddress,
		RelayerAddress: "0xd595f7c2c071d3fd8f5587931edf34e92f9ad39f",
		Version:        perpetual.Version,
	}, nil
}
