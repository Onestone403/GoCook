package service

import (
	"GoCook/model"
	"log"
)

var (
	recipeStore map[uint]*model.Recipe
	actRecipeId uint = 1
)

func init() {
	recipeStore = make(map[uint]*model.Recipe)
}

func CreateRecipe(recipe *model.Recipe) error {
	recipe.Id = actRecipeId
	recipeStore[actRecipeId] = recipe
	actRecipeId++
	log.Printf("Successfully stored new campaign with ID %v in database.", recipe.Id)
	log.Printf("Stored: %v", recipe)
	return nil
}

func GetRecipe(id uint) (*model.Recipe, error) {
	recipe, ok := recipeStore[id]
	if !ok {
		return nil, nil
	}
	return recipe, nil
}

func UpdateRecipe(id uint, recipe *model.Recipe) (*model.Recipe, error) {
	existingRecipe, err := GetRecipe(id)
	if existingRecipe == nil || err != nil {
		return existingRecipe, err
	}

	existingRecipe.Name = recipe.Name
	existingRecipe.Ingredients = recipe.Ingredients

	recipeStore[id] = existingRecipe
	return existingRecipe, nil
}

func DeleteRecipe(id uint) (*model.Recipe, error) {
	recipe, err := GetRecipe(id)
	if recipe == nil || err != nil {
		return recipe, err
	}
	delete(recipeStore, id)
	return recipe, nil
}

func GetRecipes() ([]*model.Recipe, error) {
	recipes := make([]*model.Recipe, 0, len(recipeStore))
	for _, recipe := range recipeStore {
		recipes = append(recipes, recipe)
	}
	return recipes, nil
}
