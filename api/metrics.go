package api

import (
	"context"
	"net/http"
	"time"

	"github.com/labstack/echo"
	"github.com/mcarloai/mai-v3-broker/conf"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	logger "github.com/sirupsen/logrus"
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

func StartMetricsServer(ctx context.Context) error {
	e := echo.New()
	e.GET("/metrics", echo.WrapHandler(promhttp.Handler()))
	s := &http.Server{
		Addr:         conf.Conf.MetricsAddr,
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
	}
	srvFail := make(chan error, 1)
	go func() {
		if err := e.StartServer(s); err != nil {
			srvFail <- err
		}
	}()

	select {
	case <-ctx.Done():
		logger.Infof("metrics server shutdown")
		e.Listener.Close()
		// now close the server gracefully ("shutdown")
		graceTime := 10 * time.Second
		timeoutCtx, cancel := context.WithTimeout(context.Background(), graceTime)
		if err := e.Shutdown(timeoutCtx); err != nil {
			logger.Errorf("shutdown metrics server error:%s", err.Error())
		}
		cancel()
	case err := <-srvFail:
		return err
	}
	return nil
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
