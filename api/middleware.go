package api

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ValidateJWT(app *Config) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		_, cancel := context.WithTimeout(context.Background(), appTimeout)
		defer cancel()

		token := ctx.GetHeader("Authorization")

		if token == "" {
			ctx.JSON(http.StatusForbidden, JsonResponse{Status: http.StatusForbidden, Message: "failed to extract jwt token from request"})
			return
		}

		phone, err := validateToken(token)

		if err != nil {
			ctx.JSON(http.StatusForbidden, JsonResponse{Status: http.StatusForbidden, Message: "Jwt token is invalid", Data: err})
			return
		}

		ctx.Set("phone", phone)

		ctx.Next()

	}
}

// Add middleware for cors handling
func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	}
}
