package middleware

import (
	"fmt"
	"github.com/bandanascripts/phantompay/pkg/client/token"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func AuthToken(c *gin.Context) (string, error) {

	var authHeader = c.GetHeader("Authorization")

	if authHeader == "" {
		return "", fmt.Errorf("missing token")
	}

	if !strings.HasPrefix(authHeader, "Bearer ") {
		return "", fmt.Errorf("invalid token format")
	}

	var authToken = strings.TrimPrefix(authHeader, "Bearer ")

	return authToken, nil
}

func ExtractUserId(c *gin.Context) (string, error) {

	userCtx , exists := c.Get("userClaim")

	if !exists {
		return "" , fmt.Errorf("user claim not set in context")
	}

	userClaim, ok := userCtx.(*token.UserClaim)

	if !ok {
		return "", fmt.Errorf("context does not contain user claim")
	}

	return userClaim.UserId, nil
}

func TokenMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		authToken, err := AuthToken(c)

		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			return
		}

		userClaim, err := token.ValidateToken(c.Request.Context(), authToken)

		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			return
		}

		c.Set("userClaim", userClaim)
		c.Next()
	}
}
