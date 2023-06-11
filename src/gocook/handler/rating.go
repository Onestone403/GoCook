package handler

import (
	"GoCook/model"
	"GoCook/service"
	"encoding/json"
	"log"
	"net/http"
)

func AddRating(w http.ResponseWriter, r *http.Request) {
	rating, err := getRatingFromRequest(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	recipeId, err := getId(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	rating, err = service.AddRating(r.Context(), recipeId, rating)
	if err != nil {
		log.Printf("Failure adding rating to recipe with ID %v: %v", recipeId, err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	sendJson(w, rating)
}

func getRatingFromRequest(r *http.Request) (*model.Rating, error) {
	var rating model.Rating
	err := json.NewDecoder(r.Body).Decode(&rating)
	if err != nil {
		log.Printf("Can't decode request body to recipe struct: %v", err)
		return nil, err
	}
	return &rating, nil
}
