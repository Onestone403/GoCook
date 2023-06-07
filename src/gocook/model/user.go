package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	FirstName string             `bson:"FirstName,omitempty"`
	LastName  string             `bson:"LastName,omitempty"`
	IsCook    bool
}
