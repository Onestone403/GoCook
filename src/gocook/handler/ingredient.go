package handler

import (
	"encoding/json"
	"gocook/model"
	"gocook/service"
	"net/http"

	log "github.com/sirupsen/logrus"
)

func AddIngredient(w http.ResponseWriter, r *http.Request) {
	ingredient, err := getIngredientFromRequest(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	shoppingListId, err := getId(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	ingredient, err = service.AddIngredient(r.Context(), shoppingListId, ingredient)
	if err != nil {
		log.Errorf("Failure adding ingredient: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	sendJson(w, ingredient)
}

func getIngredientFromRequest(r *http.Request) (*model.Ingredient, error) {
	var ingredient model.Ingredient
	err := json.NewDecoder(r.Body).Decode(&ingredient)
	if err != nil {
		log.Errorf("Can't decode request body to ingredient struct: %v", err)
		return nil, err
	}
	return &ingredient, nil
}
