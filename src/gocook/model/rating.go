package model

type Rating struct {
	Score   int    `bson:"Score,omitempty"`
	Comment string `bson:"Comment,omitempty"`
	UserId  uint   `bson:"UserId,omitempty"`
}
