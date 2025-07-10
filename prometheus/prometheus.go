package prometheus

import (
	"latencytracker/metrics"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func PrometheusMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		path := c.FullPath()

		c.Next()

		duration := time.Since(start).Seconds()
		status := strconv.Itoa((c.Writer.Status()))

		metrics.HttpRequestTotal.WithLabelValues(path, c.Request.Method, status).Inc()
		metrics.HttpRequestDuration.WithLabelValues(path).Observe(duration)
	}
}
