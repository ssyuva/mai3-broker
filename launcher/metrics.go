package launcher

import "github.com/prometheus/client_golang/prometheus"

var (
	mSyncingDuration = prometheus.NewSummary(prometheus.SummaryOpts{
		Name: "syncer_duration",
		Help: "Duration of syncing block.",
	})

	mTxPendingDuration = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "launcher_transaction_pending_duration",
			Help: "Pending duration of transaction.",
		}, []string{"replayer"},
	)
)

func init() {
	prometheus.MustRegister(mSyncingDuration)
	prometheus.MustRegister(mTxPendingDuration)
}
