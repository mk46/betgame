package api

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID      primitive.ObjectID `json:"id,omitempty" validate:"required" bson:"_id"`
	Name    string             `json:"name,omitempty" validate:"required" bson:"name"`
	Phone   string             `json:"phone,omitempty" validate:"required" bson:"phone"`
	Email   string             `json:"email,omitempty" validate:"required" bson:"email"`
	Balance int                `json:"balance,omitempty" validate:"required" bson:"balance"`
}

type ParsePhone struct {
	Phone string `json:"phone,omitempty" validate:"required"`
}

type JsonResponse struct {
	Status  int    `json:"status,omitempty" validate:"required"`
	Message string `json:"message,omitempty" validate:"required"`
	Data    any    `json:"data,omitempty" validate:"required"`
}

type OTP struct {
	Phone string `json:"phone,omitempty" validate:"required"`
	Code  string `json:"code,omitempty" validate:"required"`
}

type Game struct {
	ID     primitive.ObjectID `json:"id,omitempty" validate:"required" bson:"_id"`
	Name   string             `json:"name,omitempty" validate:"required" bson:"name"`
	Start  time.Time          `json:"start,omitempty" validate:"required" bson:"start"`
	End    time.Time          `json:"end,omitempty" validate:"required" bson:"end"`
	Result int                `json:"result,omitempty" validate:"required" bson:"result"`
}

type RescheduleGame struct {
	ID    primitive.ObjectID `json:"id,omitempty" validate:"required"`
	Start time.Time          `json:"start,omitempty" validate:"required"`
	End   time.Time          `json:"end,omitempty" validate:"required"`
}

type Bet struct {
	ID       primitive.ObjectID `json:"id,omitempty" validate:"required" bson:"_id"`
	UserId   primitive.ObjectID `json:"userid,omitempty" validate:"required" bson:"userid"`
	Amount   int                `json:"amount,omitempty" validate:"required" bson:"amount"`
	Number   int                `json:"number,omitempty" validate:"required" bson:"number"`
	GameID   primitive.ObjectID `json:"gameid,omitempty" validate:"required" bson:"gameid"`
	PlacedAt time.Time          `json:"placed_at,omitempty" validate:"required" bson:"placed_at"`
}

type BetHistory struct {
	ID         primitive.ObjectID `json:"id,omitempty" validate:"required" bson:"_id"`
	PlacedBet  Bet                `json:"placed_bet,omitempty" validate:"required" bson:"placed_bet"`
	ResultTime time.Time          `json:"result_time,omitempty" validate:"required" bson:"result_time"`
	Winner     bool               `json:"winner,omitempty" validate:"required" bson:"winner"`
}

type Cash struct {
	Amount int `json:"amount,omitempty" validate:"required"`
}

type DeclareWinner struct {
	GameID primitive.ObjectID `json:"gameid,omitempty" validate:"required" bson:"gameid"`
	Result int                `json:"result,omitempty" validate:"required" bson:"result"`
}
