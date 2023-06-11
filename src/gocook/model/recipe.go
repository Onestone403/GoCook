package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Recipe struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"`
	Name        string             `bson:"name"`
	Ingredients []Ingredient       `bson:"ingredients"`
	CookID      primitive.ObjectID `bson:"cookId"`
	Ratings     []Rating           `bson:"ratings,omitempty"`
}
