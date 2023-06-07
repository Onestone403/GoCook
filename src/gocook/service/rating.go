package service

import (
	"GoCook/db"
	"GoCook/model"
	"errors"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func AddRating(recipeId primitive.ObjectID, rating *model.Rating) error {
	recipe, err := GetRecipe(recipeId)
	if err != nil {
		return err
	}
	//Check if user already rated
	if recipe.Ratings != nil {
		for _, r := range recipe.Ratings {
			if r.UserID == rating.UserID {
				log.Printf("User already rated this recipe")
				return errors.New("User already rated this recipe")
			}
		}
	}
	db.RecipeCollection.UpdateOne(db.Ctx, bson.M{"_id": recipeId}, bson.M{"$push": bson.M{"ratings": rating}})
	return nil

}
