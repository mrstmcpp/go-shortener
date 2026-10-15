package config

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

var DB *mongo.Database

func ConnectDB() {
	ctx, cancel := context.WithTimeout(
		context.Background(),
		10*time.Second,
	)

	defer cancel()

	client, err := mongo.Connect(
		options.Client().ApplyURI("mongodb://localhost:27017"),
	)

	if err != nil {
		log.Fatal("MongoDB Connection setup failed", err)
	}

	if err := client.Ping(ctx, nil); err != nil {
		log.Fatal("MongoDB Ping failed", err)
	}

	DB = client.Database("go_shortener")

	indexModel := mongo.IndexModel{
		Keys:    bson.D{{Key: "email", Value: 1}},
		Options: options.Index().SetUnique(true),
	}

	_, err = DB.Collection("users").Indexes().CreateOne(ctx, indexModel)
	if err != nil {
		log.Fatal("Failed to create email index", err)
	}

	log.Println("MongoDB Connected Successfully.")
}
