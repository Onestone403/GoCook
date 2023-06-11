package service

import (
	authorization "GoCook/authorization"
	"GoCook/db"
	"GoCook/model"
	"context"
	"errors"
	"log"

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
		log.Printf("Error while checking authorization: %v", err)
		return nil, err
	}
	if !allowed {
		log.Printf("Not authorized to add a rating")
		return nil, errors.New("Not authorized to add a rating")
	}
	_, err = db.RecipeCollection.UpdateOne(ctx, bson.M{"_id": recipeId}, bson.M{"$push": bson.M{"ratings": rating}})
	if err != nil {
		log.Printf("Error while updating recipe: %v", err)
		return nil, err
	}
	log.Printf("Successfully added rating to recipe: %v", recipteToRate.Name)
	return rating, nil

}
