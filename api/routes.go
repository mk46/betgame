package api

import (
	"github.com/gin-gonic/gin"
)

type Config struct {
	Router *gin.Engine
}

func (app *Config) Routes() {
	app.Router.POST("/loginuser", LoginUserController(app))
	app.Router.POST("/updateuser", UpdateUserController(app))
	app.Router.POST("/verifyOTP", VerifyOTP(app))
	app.Router.GET("/getuser", ValidateJWT(app), GetUserController(app))
	app.Router.POST("/addgame", AddGameController(app))
	app.Router.POST("/getgames", GetGamesController(app))
	app.Router.POST("/reschedulegame/", RescheduleGameController(app))
	// app.Router.POST("/deactivategame")
	app.Router.POST("/addbet/:userid",AddBetController(app))
	app.Router.POST("/declairewinner", DeclairWinnerController(app))
	app.Router.POST("/addcash/:userid",AddCashController(app))
	app.Router.POST("/withdrawcash/:userid",WithdrawCashController(app))

	/*
		User{Name, phone, balance,photo, email(u), bets[cb]}

		Bets {gameid,uid,amount,date,numberMap[]}

		Payments{uid, amount, date , w/d,}

		Games{Name, number, Schedule End, Schedule start(Time), result date}
		/login(phone)->OTP->Verify(OTP) -> Show profile

		/updateProfile Name,Photo,
		/addgame
		/reschedulegame
		/deactivategame
		/addbets
		/declarewinner
		/addcash
		/withdrawcash





	*/
}

//Mongo db
