package metrics

import (
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
)

// requestCount Counter for requests count
var requestCount = prometheus.NewCounterVec(
	prometheus.CounterOpts{
		Namespace: "example",
		Name:      "requests_count",
		Help:      "Request counter per method, url and status code",
	},
	[]string{"method", "url", "status"},
)

// latencyHistogram Histogram for the request latency
var latencyHistogram = prometheus.NewHistogramVec(
	prometheus.HistogramOpts{
		Namespace: "example",
		Name:      "request_latency_ms",
		Help:      "Histogram for request latency in millisecond",
		Buckets:   []float64{100, 200, 300, 400, 500, 600, 700, 800, 900, 1000},
	},
	[]string{"method", "url"},
)

// registerRequestMetrics registers the prometheus metrics into the registry
func registerRequestMetrics(registry prometheus.Registerer) {
	registry.MustRegister(requestCount, latencyHistogram)
}

func MetricsMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// Record the start time
		startTime := time.Now()

		// Process the request
		ctx.Next()

		// Increase requests counter
		requestCount.With(
			prometheus.Labels{
				"method": ctx.Request.Method,
				"url":    ctx.FullPath(),
				"status": strconv.Itoa(ctx.Writer.Status()),
			},
		).Inc()

		// Observe the latency value of the request
		latencyHistogram.With(
			prometheus.Labels{
				"method": ctx.Request.Method,
				"url":    ctx.FullPath(),
			},
		).Observe(float64(time.Since(startTime).Milliseconds()))
	}
}
