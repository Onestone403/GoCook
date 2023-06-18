package service

import (
	"context"
	"errors"
	"gocook/db"
	"gocook/model"

	log "github.com/sirupsen/logrus"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetUserByID(ctx context.Context) (*model.User, error) {
	var user *model.User
	userID := ctx.Value("userID")
	if userID == nil {
		log.Error("Error getting userID from context")
		return nil, errors.New("Error getting userID from context")
	}
	userID, err := primitive.ObjectIDFromHex(userID.(string))
	if err != nil {
		log.Errorf("Error converting userID to ObjectID: %v", err)
		return nil, errors.New("Error converting userID to ObjectID")
	}
	err = db.UserCollection.FindOne(ctx, bson.M{"_id": userID}).Decode(&user)
	if err != nil {
		log.Errorf("Error getting user from database:%v ", err)
		return nil, errors.New("User doesn't exist")
	}
	return user, nil
}
