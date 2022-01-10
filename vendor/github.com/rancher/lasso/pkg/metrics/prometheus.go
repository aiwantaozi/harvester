package metrics

import "github.com/prometheus/client_golang/prometheus/collectors"

func init() {
	Registry.MustRegister(
		// expose process metrics.
		collectors.NewProcessCollector(collectors.ProcessCollectorOpts{}),
		// expose Go runtime metrics.
		collectors.NewGoCollector(),
	)
}
