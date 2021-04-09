package match

import "github.com/prometheus/client_golang/prometheus"

var (
	mTxPendingDuration = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "match_transaction_pending_duration",
			Help: "Pending duration of match.",
		}, []string{"perpetual"},
	)

	activeOrderCount = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "total_active_orders",
			Help: "count of active orders per perpetual",
		},
		[]string{"perpetual"},
	)
)

func init() {
	prometheus.MustRegister(mTxPendingDuration)
	prometheus.MustRegister(activeOrderCount)
}
