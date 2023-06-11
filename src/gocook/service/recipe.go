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

var (
	recipeStore map[uint]*model.Recipe
	actRecipeId uint = 1
)

func init() {
	recipeStore = make(map[uint]*model.Recipe)
}

func CreateRecipe(ctx context.Context, recipe *model.Recipe) error {

	decisionRequest, err := PrepareDecsisionRequest(ctx, recipe)
	log.Printf("DecisionRequest: %v", decisionRequest)
	if err != nil {
		return err
	}
	allowed, err := authorization.New().IsAllowed(decisionRequest)
	log.Printf("Allowed: %v", allowed)
	if err != nil {
		return err
	}
	if !allowed {
		return errors.New("Not allowed to create recipe")
	}
	insertResult, err := db.RecipeCollection.InsertOne(ctx, recipe)
	if err != nil {
		log.Printf("Could not store new recipe in database: %v", err)
	}
	log.Printf("Successfully stored new recipe with ID %v in database.", insertResult.InsertedID)
	recipe.ID = insertResult.InsertedID.(primitive.ObjectID)
	return nil
}

func GetRecipe(ctx context.Context, id primitive.ObjectID) (*model.Recipe, error) {
	var recipe *model.Recipe
	err := db.RecipeCollection.FindOne(ctx, bson.M{"_id": id}).Decode(&recipe)
	if err != nil {
		return nil, nil
	}
	return recipe, nil
}

func UpdateRecipe(ctx context.Context, id primitive.ObjectID, recipe *model.Recipe) (*model.Recipe, error) {
	existingRecipe, err := GetRecipe(ctx, id)
	if existingRecipe == nil || err != nil {
		return existingRecipe, err
	}
	existingRecipe.Name = recipe.Name
	existingRecipe.Ingredients = recipe.Ingredients
	db.RecipeCollection.UpdateOne(ctx, bson.M{"_id": id}, bson.M{"$set": bson.M{"name": recipe.Name, "ingredients": recipe.Ingredients}})
	return existingRecipe, nil
}

func DeleteRecipe(ctx context.Context, id primitive.ObjectID) (*model.Recipe, error) {
	recipe, err := GetRecipe(ctx, id)
	if recipe == nil || err != nil {
		return recipe, err
	}
	// delete(recipeStore, id)
	return recipe, nil
}

func GetRecipes(ctx context.Context) ([]*model.Recipe, error) {
	var recipes []*model.Recipe
	recipeCursor, err := db.RecipeCollection.Find(ctx, bson.M{})
	if err != nil {
		panic(err)
	}

	if err = recipeCursor.All(ctx, &recipes); err != nil {
		panic(err)
	}

	return recipes, nil
}

func GetRecipesByIngredient(ctx context.Context, ingredient string) ([]*model.Recipe, error) {
	var recipes []*model.Recipe
	recipeCursor, err := db.RecipeCollection.Find(ctx, bson.M{"ingredients.name": ingredient})
	if err != nil {
		panic(err)
	}

	if err = recipeCursor.All(ctx, &recipes); err != nil {
		panic(err)
	}

	return recipes, nil
}
