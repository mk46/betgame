package api

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

var appTimeout = 10 * time.Second
var db = connectPostgresDB()

func LoginUser(app *Config) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		_, cancel := context.WithTimeout(context.Background(), appTimeout)
		defer cancel()

		var phone ParsePhone

		if err := ctx.BindJSON(&phone); err != nil {
			ctx.JSON(http.StatusForbidden, JsonResponse{Status: http.StatusForbidden, Message: "failed to parse request body while login", Data: err.Error()})
			return
		}

		// Send OTP
		resp, err := app.twilioSendOTP(phone.Phone)
		if err != nil {
			ctx.JSON(http.StatusForbidden, JsonResponse{Status: http.StatusForbidden, Message: "failed to send otp while login", Data: err.Error()})
			return
		}

		ctx.JSON(http.StatusAccepted, JsonResponse{Status: http.StatusAccepted, Message: "login success. Please verify OTP", Data: resp})
		Insert(db, User{Phone: phone.Phone})
	}
}

func RegisterUser(app *Config) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		_, cancel := context.WithTimeout(context.Background(), appTimeout)
		defer cancel()

		var user User

		if err := ctx.BindJSON(&user); err != nil {
			ctx.JSON(http.StatusForbidden, JsonResponse{Status: http.StatusForbidden, Message: "failed to parse request body while signup", Data: err.Error()})
		}

		// Send OTP
		resp, err := app.twilioSendOTP(user.Phone)
		if err != nil {
			ctx.JSON(http.StatusForbidden, JsonResponse{Status: http.StatusForbidden, Message: "failed to send otp while signup", Data: err.Error()})
		}

		ctx.JSON(http.StatusAccepted, JsonResponse{Status: http.StatusAccepted, Message: "success", Data: resp})

		Insert(db, user)
	}
}

func VerifyOTP(app *Config) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		_, cancel := context.WithTimeout(context.Background(), appTimeout)
		defer cancel()

		var otp OTP
		if err := ctx.BindJSON(&otp); err != nil {
			ctx.JSON(http.StatusForbidden, JsonResponse{Status: http.StatusForbidden, Message: "failed to parse OTP while otp verification", Data: err.Error()})
			return
		}

		// Verify OTP
		err := app.twilioVerifyOTP(otp.Phone, otp.Code)
		if err != nil {
			ctx.JSON(http.StatusForbidden, JsonResponse{Status: http.StatusForbidden, Message: "failed to verify otp with twilio", Data: err.Error()})
			return
		}

		token, _ := generateJWT(otp.Phone)

		ctx.JSON(http.StatusAccepted, JsonResponse{Status: http.StatusAccepted, Message: "OTP verified successfully", Data: token})

	}
}

