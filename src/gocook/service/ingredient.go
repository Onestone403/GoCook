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

func AddIngredient(ctx context.Context, shoppingListID primitive.ObjectID, ingredient *model.Ingredient) (*model.Ingredient, error) {
	shoppingList, err := GetShoppingList(ctx, shoppingListID)
	if err != nil {
		return nil, err
	}

	decisionRequest, err := PrepareDecsisionRequest(ctx, shoppingList)
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
			"shoppingList": shoppingList,
			"user":         decisionRequest.User,
		})
		entry.Warn("User not authorized to add a ingredients")
		return nil, errors.New("Not authorized to add a ingredient")
	}
	_, err = db.ShoppingListCollection.UpdateOne(ctx, bson.M{"_id": shoppingList.ID}, bson.M{"$push": bson.M{"ingredients": ingredient}})
	if err != nil {
		log.Errorf("Error while updating shopping list: %v", err)
		return nil, err
	}
	log.Infof("Successfully added ingredient %v to shopping list: %v", ingredient.Name, shoppingList.Title)
	return ingredient, nil
}
