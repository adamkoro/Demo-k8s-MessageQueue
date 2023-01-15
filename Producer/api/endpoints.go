package api

// Import required packages
import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Gin ping endpoint
func Ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}
