package handlers

import (
	"math/rand"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func LatencySimulator(c *gin.Context) {
	rand.Seed(time.Now().UnixNano())

	delay := rand.Intn(900) + 100
	time.Sleep(time.Duration(delay) * time.Millisecond)

	if rand.Float64() < 0.1 {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Internal Server Error (simulated)",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":  "Request completed successfully",
		"delaysMs": delay,
	})
}
