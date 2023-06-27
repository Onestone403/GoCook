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

var (
	recipeStore map[uint]*model.Recipe
	actRecipeId uint = 1
)

func init() {
	recipeStore = make(map[uint]*model.Recipe)
}

func CreateRecipe(ctx context.Context, recipe *model.Recipe) error {
	if ctx.Value("istest") == nil {
		decisionRequest, err := PrepareDecsisionRequest(ctx, recipe)
		if err != nil {
			return err
		}
		allowed, err := authorization.New().IsAllowed(decisionRequest)
		if err != nil {
			log.Errorf("Error while checking authorization: %v", err)
			return err
		}
		if !allowed {
			entry := log.WithFields(log.Fields{
				"recipe": recipe,
				"user":   decisionRequest.User,
			})
			entry.Warn("Not authorized to create recipe")
			return errors.New("Not authorized to create recipe")
		}
	}
	user, err := GetUserByID(ctx)
	if err != nil {
		return err
	}
	recipe.CookID = user.ID
	insertResult, err := db.RecipeCollection.InsertOne(ctx, recipe)
	if err != nil {
		log.Errorf("Could not store new recipe in database: %v", err)
		return errors.New("Could not store new recipe in database")
	}
	log.Infof("Successfully stored new recipe with ID %v in database", insertResult.InsertedID)
	recipe.ID = insertResult.InsertedID.(primitive.ObjectID)
	return nil
}

func GetRecipe(ctx context.Context, id primitive.ObjectID) (*model.Recipe, error) {
	var recipe *model.Recipe
	err := db.RecipeCollection.FindOne(ctx, bson.M{"_id": id}).Decode(&recipe)
	if err != nil {
		log.Errorf("Could not find recipe with ID %v in database: %v", id, err)
		return nil, nil
	}
	return recipe, nil
}

func UpdateRecipe(ctx context.Context, id primitive.ObjectID, recipe *model.Recipe) (*model.Recipe, error) {
	existingRecipe, err := GetRecipe(ctx, id)
	if existingRecipe == nil || err != nil {
		return nil, err
	}
	existingRecipe.Name = recipe.Name
	existingRecipe.Ingredients = recipe.Ingredients
	decisionRequest, err := PrepareDecsisionRequest(ctx, existingRecipe)
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
			"recipe": existingRecipe,
			"user":   decisionRequest.User,
		})
		entry.Warn("Not authorized to change the recipe")
		return nil, errors.New("Not authorized to change the recipe")
	}
	_, err = db.RecipeCollection.UpdateOne(ctx, bson.M{"_id": id}, bson.M{"$set": bson.M{"name": recipe.Name, "ingredients": recipe.Ingredients}})
	if err != nil {
		log.Errorf("Could not update recipe %s in database: %v", existingRecipe.Name, err)
		return nil, err
	}
	log.Infof("Successfully updated recipe %s in database", existingRecipe.Name)
	return existingRecipe, nil
}

func DeleteRecipe(ctx context.Context, id primitive.ObjectID) (*model.Recipe, error) {
	existingRecipe, err := GetRecipe(ctx, id)
	if existingRecipe == nil || err != nil {
		log.Errorf("Could not find recipe with ID %v in database", id)
		return nil, err
	}
	decisionRequest, err := PrepareDecsisionRequest(ctx, existingRecipe)
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
			"recipe": existingRecipe,
			"user":   decisionRequest.User,
		})
		entry.Warn("Not authorized to delete the recipe")
		return nil, errors.New("Not authorized to delete the recipe")
	}
	_, err = db.RecipeCollection.DeleteOne(ctx, bson.M{"_id": id})
	if err != nil {
		log.Errorf("Could not delete recipe %s from database: %v", existingRecipe.Name, err)
		return nil, err
	}
	log.Infof("Successfully deleted recipe %s from database", existingRecipe.Name)
	return existingRecipe, nil
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
		log.Error("Could not find recipes with ingredient %s in database", ingredient)
		return nil, err
	}
	if err = recipeCursor.All(ctx, &recipes); err != nil {
		log.Error("Could not retrieve recipes from cursor: %v", err)
		return nil, err
	}
	return recipes, nil
}
