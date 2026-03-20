package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func RequireAuthHeader() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := strings.TrimSpace(c.GetHeader("Authorization"))
		if authHeader == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "authorization header is required",
			})
			return
		}

		c.Next()
	}
}
