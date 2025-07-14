package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var (
	HttpRequestTotal = promauto.NewCounterVec(
		prometheus.CounterOpts{
			Name: "http_requests_total",
			Help: "Total number of HTTP requests",
		},
		[]string{"path", "method", "status"},
	)

	HttpRequestDuration = promauto.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "http_request-duration-seconds",
			Help:    "Duration of HTTP requests",
			Buckets: []float64{0.1, 0.3, 0.5, 1, 2.5, 5, 10},
		},
		[]string{"path"},
	)

	TheCounter = promauto.NewCounter(
		prometheus.CounterOpts{
			Name: "my_application_events_total",
			Help: "Total number of events processed by my application",
		},
	)

	LogCounter = promauto.NewCounterVec(
		prometheus.CounterOpts{
			Name: "my_application_logs_total",
			Help: "Total number of logs created by my application",
		},
		[]string{"level"},
	)

	ErrorCounter = promauto.NewCounterVec(
		prometheus.CounterOpts{
			Name: "my_application_simulated_errors_total",
			Help: "Total number of logs created by my application",
		},
		[]string{"level"},
	)

	SuccessCounter = promauto.NewCounterVec(
		prometheus.CounterOpts{
			Name: "my_application_success_total",
			Help: "Total number of logs created by my application",
		},
		[]string{"level"},
	)
)
