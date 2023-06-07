package db

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var RecipeCollection *mongo.Collection
var goCookDatabase *mongo.Database
var Ctx = context.TODO()

func Init() {
	log.Println("Init MongoDB!")
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(client)
	err = client.Ping(Ctx, nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Connected to MongoDB!")

	goCookDatabase = client.Database("gocook")
	RecipeCollection = goCookDatabase.Collection("recipes")
}