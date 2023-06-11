package handler

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func sendJson(w http.ResponseWriter, value interface{}) {
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(value); err != nil {
		log.Printf("Failure encoding value to JSON: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func getId(r *http.Request) (primitive.ObjectID, error) {
	vars := mux.Vars(r)
	id, err := primitive.ObjectIDFromHex(vars["id"])
	if err != nil {
		log.Printf("ID from request not a valid ObjectID!")
	}
	return id, nil
}

func getIngredient(r *http.Request) (string, error) {
	vars := mux.Vars(r)
	ingredient := vars["ingredient"]
	if len([]rune(ingredient)) < 3 {
		log.Printf("Ingredient from request not a valid ingredient!")
		return ingredient, errors.New("Ingredient from request not a valid ingredient! At least 3 characters required")
	}
	return ingredient, nil
}
