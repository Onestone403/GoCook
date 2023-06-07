package model

type Rating struct {
	Score   int    `bson:"score,omitempty"`
	Comment string `bson:"comment,omitempty"`
	UserID  uint   `bson:"userID,omitempty"`
}
