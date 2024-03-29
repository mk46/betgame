1. Update hanlder logic sequence


Bets - gameid - UserId(update)

Bets{
    Amount   int                `json:"amount,omitempty" validate:"required" bson:"amount"`
	Number   int                `json:"number,omitempty" validate:"required" bson:"number"`
	GameID   primitive.ObjectID `json:"gameid,omitempty" validate:"required" bson:"gameid"`
	PlacedAt time.Time          `json:"placed_at,omitempty" validate:"required" bson:"placed_at"`
}



BetsHistory {
    PlacedBet  Bet       `json:"placed_bet,omitempty" validate:"required" bson:"placed_bet"`
	ResultTime time.Time `json:"result_time,omitempty" validate:"required" bson:"result_time"`
	Winner     bool      `json:"winner,omitempty" validate:"required" bson:"winner"`
}
