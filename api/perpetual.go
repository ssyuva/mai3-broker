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
