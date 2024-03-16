package api

import (
	"github.com/gin-gonic/gin"
)

type Config struct {
	Router *gin.Engine
}

func (app *Config) Routes() {
	app.Router.POST("/loginuser", LoginUser(app))
	app.Router.POST("/updateuser", RegisterUser(app))
	app.Router.POST("/verifyOTP", VerifyOTP(app))
	app.Router.GET("/getuser", ValidateJWT(app))
	app.Router.POST("/addgame")
	app.Router.POST("/reschedulegame")
	app.Router.POST("/deactivategame")
	app.Router.POST("/addbet")
	app.Router.POST("/declairewinner")
	app.Router.POST("/addcash")
	app.Router.POST("/withdrawmoney")

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
