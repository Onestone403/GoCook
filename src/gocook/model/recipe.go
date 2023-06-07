package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Recipe struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"`
	Name        string             `bson:"name,omitempty"`
	Ingredients []Ingredient       `bson:"ingredients,omitempty"`
	CookID      int                `bson:"cookId,omitempty"`
	Ratings     []Rating           `bson:"ratings,omitempty"`
}
