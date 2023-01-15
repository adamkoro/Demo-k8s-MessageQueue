package api

// Import required packages
import (
	"net/http"

	"demo-k8s-messagequeue/model"

	"github.com/gin-gonic/gin"
)

// Gin ping endpoint
func Ping(c *gin.Context) {
	var msg model.PingResponse
	msg.Message = "pong"
	c.JSON(http.StatusOK, msg)
}

// Gin health endpoint
func Health(c *gin.Context) {
	var msg model.HealthResponse
	msg.Status = "UP"
	msg.Message = "Service is running"
	c.JSON(http.StatusOK, msg)
}

// Gin push endpoint
func Push(c *gin.Context) {
	var msg model.MqMessage
	c.BindJSON(&msg)
	c.JSON(http.StatusOK, msg)
}
