package handler

import (
	"encoding/json"
	"gocook/model"
	"gocook/service"
	"net/http"

	log "github.com/sirupsen/logrus"
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
		log.Errorf("Failure adding rating to recipe with ID %v: %v", recipeId, err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	sendJson(w, rating)
}

func getRatingFromRequest(r *http.Request) (*model.Rating, error) {
	var rating model.Rating
	err := json.NewDecoder(r.Body).Decode(&rating)
	if err != nil {
		log.Errorf("Can't decode request body to recipe struct: %v", err)
		return nil, err
	}
	return &rating, nil
}
