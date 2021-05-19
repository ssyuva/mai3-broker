package api

import (
	"time"

	"github.com/labstack/echo"
	"github.com/mcarloai/mai-v3-broker/conf"
	"github.com/prometheus/client_golang/prometheus"
)

// GenericHTTPCollector create a new metrics collector with given config
func GenericHTTPCollector() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		if !conf.Conf.EnableMetrics {
			return func(c echo.Context) (err error) {
				return next(c)
			}
		}
		return func(c echo.Context) (err error) {
			start := time.Now()
			err = next(c)
			duration := float64(time.Since(start).Milliseconds())
			httpDuration.WithLabelValues(c.Path()).Observe(duration)
			labels := prometheus.Labels{
				"method": c.Request().Method,
			}
			if err != nil {
				labels["error"] = err.Error()
				labels["status"] = "error"
			} else {
				labels["error"] = ""
				labels["status"] = "success"
			}
			httpReqs.With(labels).Inc()
			return err
		}
	}
}

/// metrics
var (
	httpDuration = prometheus.NewSummaryVec(
		prometheus.SummaryOpts{
			Name:       "http_durations_ms",
			Help:       "api latency distributions.",
			Objectives: map[float64]float64{0.5: 0.05, 0.9: 0.01, 0.99: 0.001},
		},
		[]string{"path"},
	)
	httpReqs = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "http_requests_total",
			Help: "http request by status",
		},
		[]string{"error", "method", "status"},
	)
)

func init() {
	prometheus.MustRegister(httpDuration)
	prometheus.MustRegister(httpReqs)
}
