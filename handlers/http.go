package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func getHealthy(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "The pod is healthy"})
}

func getLatency(c *gin.Context) {
	c.JSON(http.StatusCreated, gin.H{"message": "Latency ping created!"})
}

func RegisterRoutes(server *gin.Engine) {
	latencyRoutes := server.Group("/tracking")
	{
		latencyRoutes.GET("/healthy", getHealthy)
		latencyRoutes.GET("/latency-ping", getLatency)
		latencyRoutes.GET("/simulate-latency", LatencySimulator)
	}
}
