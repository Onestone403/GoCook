package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type ShoppingList struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"`
	UserId      primitive.ObjectID `bson:"UserId,omitempty"`
	Ingredients []Ingredient       `bson:"Ingredients,omitempty"`
}
