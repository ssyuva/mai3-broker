package auth

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/mcarloai/mai-v3-broker/common/mai3"
	"github.com/mcarloai/mai-v3-broker/conf"
)

var (
	MaiAuthTimeout   = time.Minute * 5
	JwtSigningMethod = "HS256"
	JwtExpiration    = time.Hour * 24
)

// get time from a Mai-Authentication
func getMaiAuthTimestamp(token string) (time.Time, error) {
	var timestamp time.Time
	fields := strings.Split(token, "@")
	if len(fields) != 2 {
		return timestamp, errors.New("Token should be MAI-AUTHENTICATION@{time}")
	}
	// parse rfc
	t, err := time.Parse(time.RFC3339Nano, fields[1])
	if err == nil {
		return t, err
	}
	n, err := strconv.ParseInt(fields[1], 10, 64)
	if err == nil {
		return time.Unix(n/1000, n%1000*1e6), nil
	}
	return timestamp, errors.New("Time field should be a unix timestamp or an rfc3339 time string")
}

// validate Mai-Authentication
func ValidateMaiAuth(token string) (string, error) {
	maiAuthTokens := strings.Split(token, "#")
	if len(maiAuthTokens) != 3 {
		return "", fmt.Errorf("Mai-Authentication should be like {address}#MAI-AUTHENTICATION@{time}#{signature}")
	}
	signTime, err := getMaiAuthTimestamp(maiAuthTokens[1])
	if err != nil {
		return "", fmt.Errorf("Unable to find valid Mai-Authentication {time} field:%w", err)
	}
	now := time.Now()
	if signTime.Before(now.Add(-MaiAuthTimeout)) {
		return "", fmt.Errorf("Mai-Authentication token has expired.")
	}
	if signTime.After(now.Add(MaiAuthTimeout)) {
		return "", fmt.Errorf("Timestamp of Mai-Authentication is in the future, check your local time.")
	}
	valid, err := mai3.IsValidSignature(maiAuthTokens[0], maiAuthTokens[1], maiAuthTokens[2], mai3.EthSign)
	if !valid || err != nil {
		return "", fmt.Errorf("Token is invalid or expired, please check your authentication")
	}
	return maiAuthTokens[0], nil
}

func jwtKeyfunc(t *jwt.Token) (interface{}, error) {
	// Check the signing method
	if t.Method.Alg() != JwtSigningMethod {
		return nil, fmt.Errorf("unexpected jwt signing method=%v", t.Header["alg"])
	}
	return []byte(conf.Conf.JwtSecret), nil
}

func ValidateJwt(token string) (string, error) {
	claims := make(jwt.MapClaims)

	jwtToken, err := jwt.ParseWithClaims(token, claims, jwtKeyfunc)
	if err != nil {
		return "", fmt.Errorf("Jwt-Authentication contains invalid claims:%w", err)
	}
	if !jwtToken.Valid {
		return "", fmt.Errorf("Invalid Jwt-Authentication, please check your request")
	}
	address, ok := claims["address"].(string)
	if !ok {
		return "", fmt.Errorf("Invalid Jwt-Authentication, please check your request")
	}
	return address, nil
}

func SignJwt(address string) (string, error) {
	method := jwt.GetSigningMethod(JwtSigningMethod)
	token := jwt.NewWithClaims(method, jwt.MapClaims{
		"address": address,
		"exp":     time.Now().Add(JwtExpiration).Unix(),
	})
	return token.SignedString([]byte(conf.Conf.JwtSecret))
}
