package api

import (
	"fmt"
	"strconv"
	"time"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/mcarloai/mai-v3-broker/conf"
	"github.com/ulule/limiter/v3"
)

// Ratelimit Default settings && Variables
var (
	DefaultIPRatelimiterConfig = IPRatelimiterConfig{
		Period: time.Second * 1,
		Limit:  1000,
	}
	RateLimitEnable       = true
	RateLimitCacheService limiter.Store
)

// IPRatelimiterConfig represents the configuration of ip rale-limiter
type IPRatelimiterConfig struct {
	Skipper middleware.Skipper
	Period  time.Duration
	Limit   int64
}

// IPRatelimiter limits access rate by ip address.
// Default rate-limit is 1000 and can be set through passing in HSK_API_RATE_LIMIT env variable.
// Set HSK_API_RATE_LIMIT to 0 will diable limiter
func IPRatelimiter() echo.MiddlewareFunc {
	config := DefaultIPRatelimiterConfig
	if conf.Conf.APIRateLimit > 0 {
		config.Limit = conf.Conf.APIRateLimit
	} else {
		RateLimitEnable = false
	}
	return IPRatelimiterWithConfig(config)
}

// IPRatelimiterWithConfig creates a ip rate-limiter instance
func IPRatelimiterWithConfig(config IPRatelimiterConfig) echo.MiddlewareFunc {
	var (
		rate = limiter.Rate{
			Period: config.Period,
			Limit:  config.Limit,
		}
		// store   = memory.NewStore()
		limiter = limiter.New(RateLimitCacheService, rate)
	)

	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) (err error) {
			if limiter.Store == nil || !RateLimitEnable {
				return next(c)
			}

			ip := c.RealIP()
			limiterCtx, err := limiter.Get(c.Request().Context(), ip)
			if err != nil {
				return RateLimitError("Limiter internal error", err)
			}

			h := c.Response().Header()
			h.Set("X-RateLimit-Limit", strconv.FormatInt(limiterCtx.Limit, 10))
			h.Set("X-RateLimit-Remaining", strconv.FormatInt(limiterCtx.Remaining, 10))
			h.Set("X-RateLimit-Reset", strconv.FormatInt(limiterCtx.Reset, 10))

			if limiterCtx.Reached {
				return RateLimitError(fmt.Sprintf("Rate-limit %v exceeded for %v", config.Limit, ip), err)
			}
			return next(c)
		}
	}
}
