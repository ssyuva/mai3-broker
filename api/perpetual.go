package api

import (
	"github.com/mcarloai/mai-v3-broker/common/mai3"
	"github.com/mcarloai/mai-v3-broker/conf"
	"github.com/mcarloai/mai-v3-broker/dao"
)

func (s *Server) GetPerpetual(p Param) (interface{}, error) {
	params := p.(*GetPerpetualReq)
	perpetual, err := s.dao.GetPerpetualByPoolAddressAndIndex(params.LiquidityPoolAddress, params.PerpetualIndex, true)
	if err != nil {
		if dao.IsRecordNotFound(err) {
			return nil, PerpetualNotFoundError(params.LiquidityPoolAddress, params.PerpetualIndex)
		}
		return nil, InternalError(err)
	}
	return &GetPerpetualResp{
		Perpetual: perpetual,
	}, nil
}

func (s *Server) GetBrokerRelay(p Param) (interface{}, error) {
	params := p.(*GetBrokerRelayReq)
	_, err := s.dao.GetPerpetualByPoolAddressAndIndex(params.LiquidityPoolAddress, params.PerpetualIndex, true)
	if err != nil {
		if dao.IsRecordNotFound(err) {
			return nil, PerpetualNotFoundError(params.LiquidityPoolAddress, params.PerpetualIndex)
		}
		return nil, InternalError(err)
	}

	relayer, err := s.chainCli.GetSignAccount()
	if err != nil {
		return nil, InternalError(err)
	}

	return &GetBrokerRelayResp{
		BrokerAddress:  conf.Conf.BrokerAddress,
		RelayerAddress: relayer,
		Version:        int(mai3.ProtocolV3),
	}, nil
}
