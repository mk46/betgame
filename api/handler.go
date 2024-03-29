package api

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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
			ID:    primitive.NewObjectID(),
		}

		err = CreateUser(context.Background(), db, "users", &user)
		if err != nil {
			ctx.JSON(http.StatusForbidden, JsonResponse{Status: http.StatusForbidden, Message: "failed to insert User in DB", Data: err.Error()})
			return
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

		// // Send OTP
		// resp, err := app.twilioSendOTP(user.Phone)
		// if err != nil {
		// 	ctx.JSON(http.StatusForbidden, JsonResponse{Status: http.StatusForbidden, Message: "failed to send otp while updating user", Data: err.Error()})
		// }

		// Insert the same data on DB
		updateuser := bson.M{"$set": bson.M{"email": user.Email, "name": user.Name}}
		err := UpdateUser(context.Background(), db, "users", bson.M{"phone": user.Phone}, updateuser)
		if err != nil {
			ctx.JSON(http.StatusForbidden, JsonResponse{Status: http.StatusForbidden, Message: "failed to update User in DB", Data: err.Error()})
			return
		}
		log.Printf("User updated in db: %v\n", user)

		// Inform the ack to caller
		ctx.JSON(http.StatusAccepted, JsonResponse{Status: http.StatusAccepted, Message: "User updated successfully", Data: updateuser})

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
			ctx.JSON(http.StatusForbidden, JsonResponse{Status: http.StatusForbidden, Message: "failed to get User from DB", Data: err.Error()})
			return
		}
		log.Printf("User fetched: %v\n", user)
		ctx.JSON(http.StatusAccepted, JsonResponse{Status: http.StatusAccepted, Message: "Your authorized to access your data", Data: user})

	}
}

func AddGameController(app *Config) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		_, cancel := context.WithTimeout(context.Background(), appTimeout)
		defer cancel()
		var game Game
		if err := ctx.BindJSON(&game); err != nil {
			ctx.JSON(http.StatusForbidden, JsonResponse{Status: http.StatusForbidden, Message: "failed to parse Game data", Data: err.Error()})
			return
		}

		game.ID = primitive.NewObjectID()
		err := CreateGame(context.Background(), db, "games", &game)
		if err != nil {
			ctx.JSON(http.StatusForbidden, JsonResponse{Status: http.StatusForbidden, Message: "failed to insert Game in DB", Data: err.Error()})
			return
		}

		log.Printf("Game inserted to db: %v\n", game)

		ctx.JSON(http.StatusAccepted, JsonResponse{Status: http.StatusAccepted, Message: "Game created successfully", Data: game})
	}
}

func GetGamesController(app *Config) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		_, cancel := context.WithTimeout(context.Background(), appTimeout)
		defer cancel()
		var games []Game

		if err := GetGames(context.Background(), db, "games", bson.D{}, &games); err != nil {
			ctx.JSON(http.StatusForbidden, JsonResponse{Status: http.StatusForbidden, Message: "failed to parse Games data from DB", Data: err.Error()})

			return
		}
		ctx.JSON(http.StatusAccepted, JsonResponse{Status: http.StatusAccepted, Message: "Game fetched successfully", Data: games})

	}
}

func RescheduleGameController(app *Config) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		_, cancel := context.WithTimeout(context.Background(), appTimeout)
		defer cancel()
		var rgame RescheduleGame

		if err := ctx.BindJSON(&rgame); err != nil {
			ctx.JSON(http.StatusForbidden, JsonResponse{Status: http.StatusForbidden, Message: "failed to parse request body while rescheduling game", Data: err.Error()})
			return
		}

		updategame := bson.M{"$set": bson.M{"start": rgame.Start, "end": rgame.End}}
		err := UpdateGame(context.Background(), db, "games", bson.M{"_id": rgame.ID}, updategame)
		if err != nil {
			ctx.JSON(http.StatusForbidden, JsonResponse{Status: http.StatusForbidden, Message: "failed to update Game in DB", Data: err.Error()})
			return
		}
		log.Printf("Game updated in db: %v\n", rgame)

		// Inform the ack to caller
		ctx.JSON(http.StatusAccepted, JsonResponse{Status: http.StatusAccepted, Message: "Game rescheduled successfully", Data: updategame})

	}
}

func AddBetController(app *Config) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		_, cancel := context.WithTimeout(context.Background(), appTimeout)
		defer cancel()
		var bet Bet
		if err := ctx.BindJSON(&bet); err != nil {
			ctx.JSON(http.StatusForbidden, JsonResponse{Status: http.StatusForbidden, Message: "failed to parse request body while placing bet", Data: err.Error()})
			return
		}

		var user User
		id, err := primitive.ObjectIDFromHex(ctx.Param("userid"))
		if err != nil {
			ctx.JSON(http.StatusForbidden, JsonResponse{Status: http.StatusForbidden, Message: "failed to parse userid from request body while placing bet", Data: err.Error()})
			return
		}

		// Check user exists or not
		err = GetUser(context.Background(), db, "users", bson.M{"_id": id}, &user)
		if err != nil {
			ctx.JSON(http.StatusForbidden, JsonResponse{Status: http.StatusForbidden, Message: "failed to parse user from provided id ", Data: err.Error()})
			return
		}

		// check balance for bet
		if user.Balance < bet.Amount {
			ctx.JSON(http.StatusForbidden, JsonResponse{Status: http.StatusForbidden, Message: "insufficient fund to place order for uid: " + id.Hex(), Data: "Please add cash to wallet"})
			return
		}
		user.Balance -= bet.Amount
		bet.UserId = user.ID
		bet.ID = primitive.NewObjectID()
		if bet.PlacedAt.IsZero() {
			bet.PlacedAt = time.Now().UTC()
		}

		err = CreateBet(context.Background(), db, "bets", &bet)
		if err != nil {
			ctx.JSON(http.StatusForbidden, JsonResponse{Status: http.StatusForbidden, Message: "failed to update Bet in DB for uid: " + id.Hex(), Data: err.Error()})
			return
		}
		log.Printf("Bet placed for user:%v in db\n", user)

		updatebal := bson.M{"$set": bson.M{"balance": user.Balance}}
		err = UpdateUser(context.Background(), db, "users", bson.M{"_id": user.ID}, updatebal)
		if err != nil {
			ctx.JSON(http.StatusForbidden, JsonResponse{Status: http.StatusForbidden, Message: "failed to update Balance in DB for uid: " + id.Hex(), Data: err.Error()})
			return
		}

		log.Printf("Balance updated for user:%v in db\n", user)

		ctx.JSON(http.StatusOK, JsonResponse{Status: http.StatusOK, Message: "Bet placed successfully", Data: user})

	}
}

func DeclairWinnerController(app *Config) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		_, cancel := context.WithTimeout(context.Background(), appTimeout)
		defer cancel()
		var winner DeclareWinner
		if err := ctx.BindJSON(&winner); err != nil {
			ctx.JSON(http.StatusForbidden, JsonResponse{Status: http.StatusForbidden, Message: "failed to parse request body while declaring winner", Data: err.Error()})
			return
		}

		// Get all bets associated with gameId

		var bets []Bet

		if err := GetBets(context.Background(), db, "bets", bson.M{"gameid": winner.GameID}, &bets); err != nil {
			ctx.JSON(http.StatusForbidden, JsonResponse{Status: http.StatusForbidden, Message: "failed to parse Games data from DB", Data: err.Error()})

			return
		}

		var bethistories []interface{}

		for _, bet := range bets {
			bethistory := BetHistory{
				ID:         primitive.NewObjectID(),
				ResultTime: time.Now().UTC(),
				PlacedBet:  bet,
			}
			if winner.Result == bet.Number {
				bethistory.Winner = true

			}
			// Processing winning amount to user

			// Get user from bet
			var user User
			err := GetUser(context.Background(), db, "users", bson.M{"_id": bet.UserId}, &user)
			if err != nil {
				ctx.JSON(http.StatusForbidden, JsonResponse{Status: http.StatusForbidden, Message: "failed to parse user from provided id while declairing winner", Data: err.Error()})
				return
			}

			if bethistory.Winner {
				user.Balance += 3 * bet.Amount
				updatebal := bson.M{"$set": bson.M{"balance": user.Balance}}
				err = UpdateUser(context.Background(), db, "users", bson.M{"_id": user.ID}, updatebal)
				if err != nil {
					ctx.JSON(http.StatusForbidden, JsonResponse{Status: http.StatusForbidden, Message: "failed to update Balance in DB while declairing winner for uid: " + user.ID.Hex(), Data: err.Error()})
					return
				}

				log.Printf("Balance updated for user:%v in db while declairing winner\n", user)

			}

			bethistories = append(bethistories, bethistory)
		}

		// Remove all bets and move to BetHistory

		// Remove bets from DB
		if err := DeleteBets(context.Background(), db, "bets", bson.M{"gameid": winner.GameID}); err != nil {
			ctx.JSON(http.StatusForbidden, JsonResponse{Status: http.StatusForbidden, Message: "failed to remove Bets from DB", Data: err.Error()})
			return
		}

		// Insert BetHistory to DB
		if err := AddBetHistory(context.Background(), db, "bethistories", bethistories); err != nil {
			ctx.JSON(http.StatusForbidden, JsonResponse{Status: http.StatusForbidden, Message: "failed to insert Bets to BetHistories in DB", Data: err.Error()})
			return
		}

		// Done
		ctx.JSON(http.StatusAccepted, JsonResponse{Status: http.StatusAccepted, Message: "Winner declaired successfully", Data: bets})
	}
}

func AddCashController(app *Config) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var addcash Cash
		if err := ctx.BindJSON(&addcash); err != nil {
			ctx.JSON(http.StatusForbidden, JsonResponse{Status: http.StatusForbidden, Message: "failed to parse request body while adding cash", Data: err.Error()})
			return
		}

		var user User
		id, err := primitive.ObjectIDFromHex(ctx.Param("userid"))
		if err != nil {
			ctx.JSON(http.StatusForbidden, JsonResponse{Status: http.StatusForbidden, Message: "failed to parse userid from request body while adding cash", Data: err.Error()})
			return
		}

		err = GetUser(context.Background(), db, "users", bson.M{"_id": id}, &user)
		if err != nil {
			ctx.JSON(http.StatusForbidden, JsonResponse{Status: http.StatusForbidden, Message: "failed to parse user from provided id while adding cash", Data: err.Error()})
			return
		}

		user.Balance += addcash.Amount

		updateuser := bson.M{"$set": bson.M{"balance": user.Balance}}
		err = UpdateUser(context.Background(), db, "users", bson.M{"_id": user.ID}, updateuser)
		if err != nil {
			ctx.JSON(http.StatusForbidden, JsonResponse{Status: http.StatusForbidden, Message: "failed to add cash in DB for uid: " + id.Hex(), Data: err.Error()})
			return
		}

		log.Printf("Cash added for user:%v in db\n", user)

		ctx.JSON(http.StatusOK, JsonResponse{Status: http.StatusOK, Message: "Cash added successfully", Data: user})
	}
}

func WithdrawCashController(app *Config) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var withdrawcash Cash
		if err := ctx.BindJSON(&withdrawcash); err != nil {
			ctx.JSON(http.StatusForbidden, JsonResponse{Status: http.StatusForbidden, Message: "failed to parse request body while withdrawing cash", Data: err.Error()})
			return
		}

		var user User
		id, err := primitive.ObjectIDFromHex(ctx.Param("userid"))
		if err != nil {
			ctx.JSON(http.StatusForbidden, JsonResponse{Status: http.StatusForbidden, Message: "failed to parse userid from request body while withdrawing cash", Data: err.Error()})
			return
		}

		err = GetUser(context.Background(), db, "users", bson.M{"_id": id}, &user)
		if err != nil {
			ctx.JSON(http.StatusForbidden, JsonResponse{Status: http.StatusForbidden, Message: "failed to parse user from provided id while withdrawing cash", Data: err.Error()})
			return
		}

		if user.Balance < withdrawcash.Amount {
			ctx.JSON(http.StatusForbidden, JsonResponse{Status: http.StatusForbidden, Message: "insufficient fund while withdrawing cash", Data: user.Balance})
			return
		}

		user.Balance -= withdrawcash.Amount

		updateuser := bson.M{"$set": bson.M{"balance": user.Balance}}
		err = UpdateUser(context.Background(), db, "users", bson.M{"_id": user.ID}, updateuser)
		if err != nil {
			ctx.JSON(http.StatusForbidden, JsonResponse{Status: http.StatusForbidden, Message: "failed to withdraw cash in DB for uid: " + id.Hex(), Data: err.Error()})
			return
		}

		log.Printf("Cash added for user:%v in db\n", user)

		ctx.JSON(http.StatusOK, JsonResponse{Status: http.StatusOK, Message: "Cash withdrew successfully", Data: user})

	}
}
