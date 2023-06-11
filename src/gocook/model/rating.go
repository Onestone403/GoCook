package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Rating struct {
	Score   int                `bson:"score,omitempty"`
	Comment string             `bson:"comment,omitempty"`
	UserID  primitive.ObjectID `bson:"userID,omitempty"`
}
