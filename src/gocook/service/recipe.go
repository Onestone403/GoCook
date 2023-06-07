package service

import (
	"GoCook/db"
	"GoCook/model"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var (
	recipeStore map[uint]*model.Recipe
	actRecipeId uint = 1
)

func init() {
	recipeStore = make(map[uint]*model.Recipe)
}

func CreateRecipe(recipe *model.Recipe) error {
	insertResult, err := db.RecipeCollection.InsertOne(db.Ctx, recipe)
	if err != nil {
		log.Printf("Could not store new recipe in database: %v", err)
	}
	log.Printf("Successfully stored new recipe with ID %v in database.", insertResult.InsertedID)
	recipe.ID = insertResult.InsertedID.(primitive.ObjectID)
	return nil
}

func GetRecipe(id primitive.ObjectID) (*model.Recipe, error) {
	var recipe *model.Recipe
	err := db.RecipeCollection.FindOne(db.Ctx, bson.M{"_id": id}).Decode(&recipe)
	if err != nil {
		return nil, nil
	}
	return recipe, nil
}

func UpdateRecipe(id primitive.ObjectID, recipe *model.Recipe) (*model.Recipe, error) {
	existingRecipe, err := GetRecipe(id)
	if existingRecipe == nil || err != nil {
		return existingRecipe, err
	}
	existingRecipe.Name = recipe.Name
	existingRecipe.Ingredients = recipe.Ingredients
	db.RecipeCollection.UpdateOne(db.Ctx, bson.M{"_id": id}, bson.M{"$set": bson.M{"name": recipe.Name, "ingredients": recipe.Ingredients}})
	return existingRecipe, nil
}

func DeleteRecipe(id primitive.ObjectID) (*model.Recipe, error) {
	recipe, err := GetRecipe(id)
	if recipe == nil || err != nil {
		return recipe, err
	}
	// delete(recipeStore, id)
	return recipe, nil
}

func GetRecipes() ([]*model.Recipe, error) {
	// recipes := make([]*model.Recipe, 0, len(recipeStore))
	// for _, recipe := range recipeStore {
	// 	recipes = append(recipes, recipe)
	// }

	var recipes []*model.Recipe
	recipeCursor, err := db.RecipeCollection.Find(db.Ctx, bson.M{})
	if err != nil {
		panic(err)
	}

	if err = recipeCursor.All(db.Ctx, &recipes); err != nil {
		panic(err)
	}

	return recipes, nil
}

func GetRecipesByIngredient(ingredient string) ([]*model.Recipe, error) {
	var recipes []*model.Recipe
	recipeCursor, err := db.RecipeCollection.Find(db.Ctx, bson.M{"ingredients.name": ingredient})
	if err != nil {
		panic(err)
	}

	if err = recipeCursor.All(db.Ctx, &recipes); err != nil {
		panic(err)
	}

	return recipes, nil
}
