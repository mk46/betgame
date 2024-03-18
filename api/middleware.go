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

		token := ctx.GetHeader("auth-key")

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
