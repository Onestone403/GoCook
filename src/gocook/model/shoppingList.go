package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type ShoppingList struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"`
	UserID      primitive.ObjectID `bson:"userID,omitempty"`
	Title       string             `bson:"title,omitempty"`
	Ingredients []Ingredient       `bson:"ingredients,omitempty"`
}
