package service

import "GoCook/model"

func AddRating(recipeId uint, rating *model.Rating) error {
	recipe, err := GetRecipe(recipeId)
	if err != nil {
		return err
	}
	recipe.Ratings = append(recipe.Ratings, *rating)
	return nil
}
