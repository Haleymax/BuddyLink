package middleware

import (
	"buddylink/pkg/jwtutil"
	"github.com/gin-gonic/gin"
	"net/http"
)

func UserAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("Authorization")
		if token == "" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code":    http.StatusUnauthorized,
				"message": "no token found",
				"status":  false,
				"data":    nil,
			})
			return
		}

		jwtUtil := jwtutil.NewJWTUtil("secret")

		claims, err := jwtUtil.ParseToken(token)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code":    http.StatusUnauthorized,
				"message": "invalid token",
				"status":  false,
				"data":    nil,
			})
			return
		}
		c.Set("username", claims.Username)
		c.Set("id", claims.Id)
		c.Set("email", claims.Email)
		c.Next()
	}
}
