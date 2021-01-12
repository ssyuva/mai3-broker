package api

import (
	"github.com/mcarloai/mai-v3-broker/common/auth"
	"github.com/mcarloai/mai-v3-broker/conf"
	"strings"
	"time"
)

func GetJwtAuth(param Param) (interface{}, error) {
	address := param.GetAddress()
	// white list
	find := false
	for _, item := range conf.Conf.WhiteList {
		if strings.ToLower(address) == strings.ToLower(item) {
			find = true
			break
		}
	}
	if !find {
		return nil, AddressNotInWhiteListError()
	}
	jwt, err := auth.SignJwt(address)
	if err != nil {
		return "", err
	}
	expires := auth.JwtExpiration / time.Second * 1000
	return map[string]interface{}{
		"jwt":     jwt,
		"expires": expires,
	}, nil
}
