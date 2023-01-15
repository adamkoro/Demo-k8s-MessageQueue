package api

import "github.com/gin-gonic/gin"

// Create Gin router group to api/v1 endpoint
func ApiV1() *gin.Engine {
	router := gin.Default()
	v1 := router.Group("/api/v1")
	{
		v1.GET("/ping", Ping)
	}
	return router
}
