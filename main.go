package main

import (
	"latencytracker/handlers"
	"latencytracker/prometheus"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	server := gin.Default()
	server.Use(prometheus.PrometheusMiddleware())

	server.GET("/metrics", gin.WrapH(promhttp.Handler()))

	handlers.RegisterRoutes(server)
	server.Run(":8080")
}
