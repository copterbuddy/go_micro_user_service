package intercepter

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

// GeneralInterceptor - call this methos to add interceptor
func GeneralInterceptor(c *gin.Context) {
	token := c.Request.Header.Get("Authorization")

	if len(token) <= len("Bearer ") {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid token"})
		c.Abort()
	}

	parts := strings.SplitN(token, " ", 2)
	if parts[0] != "Bearer" {
		c.JSON(401, &gin.H{
			"code":  401,
			"error": "Invalid Authorization header",
		})
		c.Abort()
		return
	}
	userId, isOk := JwtVerify(parts[1])
	if !isOk || userId == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid token"})
		c.Abort()
	}

	c.Set("Issuer", userId)

	c.Next()
}
