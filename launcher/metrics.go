package launcher

import "github.com/prometheus/client_golang/prometheus"

var (
	txConfirmDuration = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "transaction_confirm_duration",
			Help: "Confirm duration of transaction.",
		}, []string{"replayer"},
	)

	txPendingDuration = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "transaction_pending_duration",
			Help: "Pending duration of transaction.",
		}, []string{"replayer"},
	)
)

func init() {
	prometheus.MustRegister(txConfirmDuration)
	prometheus.MustRegister(txPendingDuration)
}
