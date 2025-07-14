package main

import (
	"encoding/json"
	"fmt"
	"latencytracker/handlers"
	"latencytracker/middleware"
	"latencytracker/prometheus"
	"math/rand"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	go func() {
		for {
			logEntry := middleware.GenerateLogs()

			data, _ := json.Marshal(logEntry)
			fmt.Println(string(data))

			time.Sleep(time.Duration(rand.Intn(1000)+500) * time.Millisecond)
		}
	}()

	server := gin.Default()
	server.Use(prometheus.PrometheusMiddleware())

	server.GET("/metrics", gin.WrapH(promhttp.Handler()))

	handlers.RegisterRoutes(server)
	server.Run(":8080")

}
