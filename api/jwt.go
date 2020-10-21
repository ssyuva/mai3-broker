package api

import (
	"github.com/mcarloai/mai-v3-broker/common/auth"
	"time"
)

func GetJwtAuth(param Param) (interface{}, error) {
	address := param.GetAddress()
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
