package db

import (
	"context"
	"os"

	log "github.com/sirupsen/logrus"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var RecipeCollection *mongo.Collection
var UserCollection *mongo.Collection
var ShoppingListCollection *mongo.Collection
var goCookDatabase *mongo.Database

func Init() {
	log.Info("Init MongoDB!")
	clientOptions := options.Client().ApplyURI(os.Getenv("MONGODB_URI"))
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(client)
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Info("Connected to MongoDB!")

	goCookDatabase = client.Database("gocook")
	RecipeCollection = goCookDatabase.Collection("recipes")
	UserCollection = goCookDatabase.Collection("users")
	ShoppingListCollection = goCookDatabase.Collection("shoppingLists")
}
