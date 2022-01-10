package metrics

import "github.com/prometheus/client_golang/prometheus"

// Registry stored metrics using lasso controller framework
var (
	Registry = prometheus.NewRegistry()
)
