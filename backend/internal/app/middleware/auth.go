package middleware

import "github.com/gin-gonic/gin"

func UserAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("Authorization")
	}
}
