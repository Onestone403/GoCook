package service

import (
	"GoCook/db"
	"GoCook/model"
	"context"
	"errors"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func AddRating(ctx context.Context, recipeId primitive.ObjectID, rating *model.Rating) error {
	recipe, err := GetRecipe(ctx, recipeId)
	if err != nil {
		return err
	}
	//TODO verify by OPA
	if recipe.Ratings != nil {
		for _, r := range recipe.Ratings {
			if r.UserID == rating.UserID {
				log.Printf("User already rated this recipe")
				return errors.New("User already rated this recipe")
			}
		}
	}
	db.RecipeCollection.UpdateOne(ctx, bson.M{"_id": recipeId}, bson.M{"$push": bson.M{"ratings": rating}})
	return nil

}
