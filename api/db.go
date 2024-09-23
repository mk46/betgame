package api

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDB() *gorm.DB {
	dbURL := "postgres://pg:pass@localhost:5432/game_db"

	db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}

	if err = db.AutoMigrate(&User{}, &Game{}, &BetHistory{}, &Bet{}); err != nil {
		log.Fatalln("Failed to migrate db: ", err.Error())
	}

	return db
}

func CreateUser(db *gorm.DB, user User) error {
	// Append to the User in User table
	if result := db.Create(&user); result.Error != nil {
		return result.Error
	}

	return nil
}

func GetUserByPhone(db *gorm.DB, phone string, user *User) error {

	if result := db.Where("phone=?", phone).First(user); result.Error != nil {
		return result.Error
	}
	return nil
}

func GetUserByID(db *gorm.DB, id int, user *User) error {

	if result := db.First(user, id); result.Error != nil {
		return result.Error
	}
	return nil
}

func UpdateUserByPhone(db *gorm.DB, updateuser User) error {
	var user User
	if err := GetUserByPhone(db, updateuser.Phone, &user); err != nil {
		return err
	}
	user.Name = updateuser.Name
	user.Email = updateuser.Email
	user.Balance = updateuser.Balance
	if result := db.Save(&user); result.Error != nil {
		return result.Error
	}
	return nil
}

func UpdateUserByID(db *gorm.DB, userid int, updateuser User) error {
	var user User
	if err := GetUserByID(db, userid, &user); err != nil {
		return err
	}

	user.Balance = updateuser.Balance
	if result := db.Save(&user); result.Error != nil {
		return result.Error
	}
	return nil
}

func DeleteUser(db *gorm.DB, phone string) error {
	var user User
	if err := GetUserByPhone(db, phone, &user); err != nil {
		return err
	}
	if result := db.Delete(&user); result.Error != nil {
		return result.Error
	}
	return nil
}

func CreateGame(db *gorm.DB, game Game) error {
	// Append to the Game in Game table
	if result := db.Create(&game); result.Error != nil {
		return result.Error
	}
	return nil
}

func GetGames(db *gorm.DB, games *[]Game) error {

	if result := db.Find(games); result.Error != nil {
		return result.Error
	}
	return nil
}

func UpdateGame(db *gorm.DB, gameid int, updategame Game) error {
	var game Game
	if result := db.First(&game, gameid); result.Error != nil {
		return result.Error
	}

	game.Start = updategame.Start
	game.End = updategame.End

	if result := db.Save(&game); result.Error != nil {
		return result.Error
	}
	return nil
}

func CreateBet(db *gorm.DB, bet Bet) error {
	if result := db.Create(&bet); result.Error != nil {
		return result.Error
	}
	return nil
}

func DeleteBets(db *gorm.DB, gameid int) error {
	var bet Bet
	if result := db.Where("game_id=?", gameid).First(&bet); result.Error != nil {
		return result.Error
	}

	log.Println("Fetched bet for delete", bet)

	if result := db.Delete(&bet); result.Error != nil {
		return result.Error
	}
	return nil
}

func AddBetHistory(db *gorm.DB, bethistory []BetHistory) error {

	if result := db.Create(&bethistory); result.Error != nil {
		return result.Error
	}
	return nil
}

func GetBets(db *gorm.DB, gameid int, bets *[]Bet) error {

	if result := db.Where("game_id=?", gameid).Find(&bets); result.Error != nil {
		return result.Error
	}

	return nil
}
