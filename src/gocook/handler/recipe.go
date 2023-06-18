package handler

import (
	"encoding/json"
	"gocook/model"
	"gocook/service"
	"net/http"

	log "github.com/sirupsen/logrus"
)

func CreateRecipe(w http.ResponseWriter, r *http.Request) {
	recipe, err := getRecipeFromRequest(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err := service.CreateRecipe(r.Context(), recipe); err != nil {
		log.Errorf("Error calling service CreateRecipe: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(recipe); err != nil {
		log.Errorf("Failure encoding value to JSON: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func GetRecipe(w http.ResponseWriter, r *http.Request) {
	id, err := getId(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	recipe, err := service.GetRecipe(r.Context(), id)
	if err != nil {
		log.Errorf("Failure retrieving recipe with ID %v: %v", id, err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if recipe == nil {
		http.Error(w, "404 recipe not found", http.StatusNotFound)
		return
	}
	sendJson(w, recipe)
}

func UpdateRecipe(w http.ResponseWriter, r *http.Request) {
	id, err := getId(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	recipe, err := getRecipeFromRequest(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	recipe, err = service.UpdateRecipe(r.Context(), id, recipe)
	if err != nil {
		log.Errorf("Failure updating recipe with ID %v: %v", id, err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if recipe == nil {
		http.Error(w, "404 recipe not found", http.StatusNotFound)
		return
	}
	sendJson(w, recipe)
}

func DeleteRecipe(w http.ResponseWriter, r *http.Request) {
	id, err := getId(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	recipe, err := service.DeleteRecipe(r.Context(), id)
	if err != nil {
		log.Errorf("Failure deleting recipe with ID %v: %v", id, err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if recipe == nil {
		http.Error(w, "404 recipe not found", http.StatusNotFound)
		return
	}
	sendJson(w, "")
}

func GetRecipes(w http.ResponseWriter, r *http.Request) {
	recipes, err := service.GetRecipes(r.Context())
	if err != nil {
		log.Errorf("Failure retrieving recipes: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if recipes == nil {
		http.Error(w, "404 recipes not found", http.StatusNotFound)
		return
	}
	sendJson(w, recipes)
}

func GetRecipesByIngredient(w http.ResponseWriter, r *http.Request) {
	ingredient, err := getIngredient(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	recipes, err := service.GetRecipesByIngredient(r.Context(), ingredient)
	if err != nil {
		log.Errorf("Failure retrieving recipes: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if recipes == nil {
		http.Error(w, "404 No recipe with Ingredient found", http.StatusNotFound)
		return
	}
	sendJson(w, recipes)
}

func getRecipeFromRequest(r *http.Request) (*model.Recipe, error) {
	var recipe model.Recipe
	err := json.NewDecoder(r.Body).Decode(&recipe)
	if err != nil {
		log.Errorf("Can't decode request body to recipe struct: %v", err)
		return nil, err
	}
	return &recipe, nil
}
