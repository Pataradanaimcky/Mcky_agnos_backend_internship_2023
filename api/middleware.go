package api

import (
	"Mcky_agnos_backend_internship_2023/db"
	"time"

	"github.com/gin-gonic/gin"
)

// RequestLogger is a middleware to log incoming requests and their processing time.
func RequestLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		startTime := time.Now()

		c.Next()

		clientIP := c.ClientIP()
		requestMethod := c.Request.Method
		requestPath := c.Request.URL.Path
		requestProto := c.Request.Proto
		statusCode := c.Writer.Status()
		latency := time.Since(startTime)

		// Log the request info in the database
		db.LogRequestInfo(clientIP, requestMethod, requestPath, requestProto, statusCode, latency)
	}
}
