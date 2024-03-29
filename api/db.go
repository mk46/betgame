package api

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectDB(uri string) *mongo.Client {

	credential := options.Credential{
		Username: "mongoadmin",
		Password: "secret",
	}
	clientOptions := options.Client().ApplyURI(uri).SetAuth(credential)
	client, err := mongo.Connect(context.Background(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Successfully connected to MongoDB")
	return client
}

func CreateUser(ctx context.Context, db *mongo.Database, collectionName string, model interface{}) error {
	collection := db.Collection(collectionName)

	_, err := collection.InsertOne(ctx, model)
	if err != nil {
		return err
	}

	// m.ID = res.InsertedID.(primitive.ObjectID)
	return nil
}

func GetUser(ctx context.Context, db *mongo.Database, collectionName string, filter interface{}, result interface{}) error {
	collection := db.Collection(collectionName)

	err := collection.FindOne(ctx, filter).Decode(result)
	if err != nil {
		return err
	}

	return nil
}

func UpdateUser(ctx context.Context, db *mongo.Database, collectionName string, filter interface{}, update interface{}) error {
	collection := db.Collection(collectionName)

	_, err := collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}

	return nil
}

func DeleteUser(ctx context.Context, db *mongo.Database, collectionName string, filter interface{}) error {
	collection := db.Collection(collectionName)
	_, err := collection.DeleteOne(ctx, filter)
	if err != nil {
		return err
	}

	return nil
}

func CreateGame(ctx context.Context, db *mongo.Database, collectionName string, model interface{}) error {
	collection := db.Collection(collectionName)

	_, err := collection.InsertOne(ctx, model)
	if err != nil {
		return err
	}

	return nil
}

func GetGames(ctx context.Context, db *mongo.Database, collectionName string, filter interface{}, result *[]Game) error {
	collection := db.Collection(collectionName)
	cur, err := collection.Find(ctx, filter)
	if err != nil {
		log.Fatal(err)
	}
	defer cur.Close(ctx)
	for cur.Next(ctx) {
		// To decode into a struct, use cursor.Decode()
		var game Game
		err := cur.Decode(&game)

		if err != nil {
			log.Fatal(err)
		}
		*result = append(*result, game)

	}
	if err := cur.Err(); err != nil {
		return err
	}
	return nil
}

func UpdateGame(ctx context.Context, db *mongo.Database, collectionName string, filter interface{}, updategame interface{}) error {
	collection := db.Collection(collectionName)

	_, err := collection.UpdateOne(ctx, filter, updategame)
	if err != nil {
		return err
	}

	return nil
}

func CreateBet(ctx context.Context, db *mongo.Database, collectionName string, bet interface{}) error {
	collection := db.Collection(collectionName)

	_, err := collection.InsertOne(ctx, bet)
	if err != nil {
		return err
	}
	return nil
}

func DeleteBets(ctx context.Context, db *mongo.Database, collectionName string, filter interface{}) error {
	collection := db.Collection(collectionName)
	_, err := collection.DeleteMany(ctx, filter)
	if err != nil {
		return err
	}

	return nil
}

func AddBetHistory(ctx context.Context, db *mongo.Database, collectionName string, bethistories []interface{}) error {
	collection := db.Collection(collectionName)

	_, err := collection.InsertMany(ctx, bethistories)
	if err != nil {
		return err
	}
	return nil
}

func GetBets(ctx context.Context, db *mongo.Database, collectionName string, filter interface{}, result *[]Bet) error {
	collection := db.Collection(collectionName)
	cur, err := collection.Find(ctx, filter)
	if err != nil {
		return err
	}
	defer cur.Close(ctx)

	if err := cur.All(ctx, result); err != nil {
		return err
	}
	log.Println(result)

	return nil
}
