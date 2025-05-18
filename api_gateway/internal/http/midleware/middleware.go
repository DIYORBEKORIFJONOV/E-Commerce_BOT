package middleware

import (
	"github.com/gin-gonic/gin"
	"time"
)


func CorsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "http://localhost:5173")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		if c.Request.Method == "OPTIONS" {
			// Завершаем обработку и возвращаем OK для preflight
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}



// func IPFilterMiddleware(allowedIPs []string) gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		clientIP := c.ClientIP()
// 		allowed := false

// 		for _, ip := range allowedIPs {
// 			if clientIP == ip {
// 				allowed = true
// 				break
// 			}
// 		}

// 		if !allowed {
// 			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{
// 				"error": "Access denied for IP: " + clientIP,
// 			})
// 			return
// 		}

// 		c.Next()
// 	}
// }

func TimingMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		c.Next()
		duration := time.Since(start)
		c.Writer.Header().Set("X-Response-Time", duration.String())
	}
}
