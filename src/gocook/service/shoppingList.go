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

func CreateShoppingList(ctx context.Context, title string) (*model.ShoppingList, error) {
	var shoppingList = &model.ShoppingList{Title: title}
	user, err := GetUserByID(ctx)
	if err != nil {
		return nil, err
	}
	shoppingList.UserID = user.ID
	insertResult, err := db.ShoppingListCollection.InsertOne(ctx, shoppingList)
	if err != nil {
		log.Errorf("Could not create shopping list of user with ID %v in database: %v", shoppingList.UserID, err)
		return nil, err
	}
	shoppingList.ID = insertResult.InsertedID.(primitive.ObjectID)
	return shoppingList, nil
}

func GetShoppingList(ctx context.Context, id primitive.ObjectID) (*model.ShoppingList, error) {
	var shoppingList *model.ShoppingList
	err := db.ShoppingListCollection.FindOne(ctx, bson.M{"_id": id}).Decode(&shoppingList)
	if err != nil {
		log.Errorf("Could not find shopping list with ID %v in database: %v", id, err)
		return nil, nil
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
		entry.Printf("Not authorized to manage Shopping list")
		return nil, errors.New("Not authorized to manage shopping list")
	}
	return shoppingList, nil
}

func GetShoppingLists(ctx context.Context) ([]*model.ShoppingList, error) {
	user, err := GetUserByID(ctx)
	if err != nil {
		return nil, err
	}
	var shoppingLists []*model.ShoppingList
	shoppingListCursor, err := db.ShoppingListCollection.Find(ctx, bson.M{"userID": user.ID})
	if err != nil {
		log.Warn("Couldn't find a shopping list of user with ID %v in database: %v", user.ID, err)
		return nil, err
	}

	if err = shoppingListCursor.All(ctx, &shoppingLists); err != nil {
		panic(err)
	}

	return shoppingLists, nil
}

func DeleteShoppingList(ctx context.Context, id primitive.ObjectID) (*model.ShoppingList, error) {
	shoppingList, err := GetShoppingList(ctx, id)
	if err != nil {
		return nil, err
	}
	if shoppingList == nil {
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
		entry.Printf("Not authorized to delete shopping list")
		return nil, errors.New("Not authorized to delete shopping list")
	}
	_, err = db.ShoppingListCollection.DeleteOne(ctx, bson.M{"_id": id})
	if err != nil {
		log.Errorf("Could not delete shopping list with ID %v in database: %v", id, err)
		return nil, err
	}
	log.Infof("Successfully deleted shopping list with id %v from database", shoppingList.ID)
	return shoppingList, nil
}
