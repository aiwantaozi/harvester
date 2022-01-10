package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
)

const (
	ControllerSubsystem        = "controller"
	ReconcileResultSuccess     = "success"
	ReconcileResultError       = "error"
	ReconcileResultErrorIgnore = "error_ignore"
)

var (
	// ReconcileTotal is a prometheus counter metrics exposes the total number of reconciliations per controller.
	// controller label refers to the controller name and result label refers to the reconcile result.
	ReconcileTotal = prometheus.NewCounterVec(prometheus.CounterOpts{
		Subsystem: ControllerSubsystem,
		Name:      "reconcile_total",
		Help:      "Total number of reconciliations per controller",
	}, []string{"controller", "result"})

	// ReconcileTime is a prometheus histogram metric exposes the duration of reconciliations per controller.
	// controller label refers to the controller name
	ReconcileTime = prometheus.NewHistogramVec(prometheus.HistogramOpts{
		Subsystem: ControllerSubsystem,
		Name:      "reconcile_time_seconds",
		Help:      "Durations per reconciliation per controller",
	}, []string{"controller"})
)

func init() {
	Registry.MustRegister(
		ReconcileTotal,
		ReconcileTime,
	)
}
