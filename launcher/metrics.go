package launcher

import "github.com/prometheus/client_golang/prometheus"

var (
	mTxPendingDuration = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "launcher_transaction_pending_duration",
			Help: "Pending duration of transaction.",
		}, []string{"replayer"},
	)
)

func init() {
	prometheus.MustRegister(mTxPendingDuration)
}
