package database

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func initDB() *mongo.Client {
	MongoDB := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
	err := MongoDB.Connect(context.TODO())
	if err != nil {
		log.Fatal(err)
	}
	return MongoDB
}

func initCollection(MongoDB *mongo.Client) *mongo.Collection
