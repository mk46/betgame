package api

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

var appTimeout = 10 * time.Second
var db = ConnectDB("mongodb://localhost:27017").Database("test_db")

func LoginUserController(app *Config) gin.HandlerFunc {
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

		user := User{
			Phone: phone.Phone,
		}

		err = CreateUser(context.Background(), db, "users", &user)
		if err != nil {
			panic(err)
		}

		log.Printf("User inserted to db: %v\n", user)

	}
}

func UpdateUserController(app *Config) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		_, cancel := context.WithTimeout(context.Background(), appTimeout)
		defer cancel()

		var user User

		if err := ctx.BindJSON(&user); err != nil {
			ctx.JSON(http.StatusForbidden, JsonResponse{Status: http.StatusForbidden, Message: "failed to parse request body while updating User", Data: err.Error()})
		}

		// Send OTP
		resp, err := app.twilioSendOTP(user.Phone)
		if err != nil {
			ctx.JSON(http.StatusForbidden, JsonResponse{Status: http.StatusForbidden, Message: "failed to send otp while updating user", Data: err.Error()})
		}

		// Insert the same data on DB
		updateuser := bson.M{"$set": bson.M{"email": user.Email, "name": user.Name}}
		err = UpdateUser(context.Background(), db, "users", bson.M{"phone": user.Phone}, updateuser)
		if err != nil {
			panic(err)
		}
		log.Printf("User updated in db: %v\n", user)

		// Inform the ack to caller
		ctx.JSON(http.StatusAccepted, JsonResponse{Status: http.StatusAccepted, Message: "User updated successfully", Data: resp})

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

func GetUserController(app *Config) gin.HandlerFunc {
	return func(ctx *gin.Context) {

		phone, _ := ctx.Get("phone")

		var user User
		err := GetUser(context.Background(), db, "users", bson.M{"phone": phone}, &user)
		if err != nil {
			panic(err)
		}
		log.Printf("User fetched: %v\n", user)
		ctx.JSON(http.StatusAccepted, JsonResponse{Status: http.StatusAccepted, Message: "Your authorized to access your data", Data: user})

	}
}
