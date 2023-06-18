package service

import (
	"context"
	"errors"
	authorization "gocook/authorization"
	"gocook/db"
	"gocook/model"

	log "github.com/sirupsen/logrus"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func AddRating(ctx context.Context, recipeId primitive.ObjectID, rating *model.Rating) (*model.Rating, error) {
	recipteToRate, err := GetRecipe(ctx, recipeId)
	if err != nil {
		return nil, err
	}
	user, err := GetUserByID(ctx)
	if err != nil {
		return nil, err
	}
	rating.UserID = user.ID
	decisionRequest, err := PrepareDecsisionRequest(ctx, recipteToRate)
	if err != nil {
		return nil, err
	}
	allowed, err := authorization.New().IsAllowed(decisionRequest)
	if err != nil {
		log.Errorf("Error while checking authorization: %v", err)
		return nil, err
	}
	if !allowed {
		entry := log.WithFields(log.Fields{
			"recipe": recipteToRate,
			"user":   decisionRequest.User,
		})
		entry.Warn("User not authorized to add a ratings")
		return nil, errors.New("Not authorized to add a rating")
	}
	_, err = db.RecipeCollection.UpdateOne(ctx, bson.M{"_id": recipeId}, bson.M{"$push": bson.M{"ratings": rating}})
	if err != nil {
		log.Errorf("Error while updating recipe: %v", err)
		return nil, err
	}
	log.Infof("Successfully added rating to recipe: %v", recipteToRate.Name)
	return rating, nil

}
