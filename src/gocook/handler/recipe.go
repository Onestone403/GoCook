package handler

import (
	"GoCook/model"
	"GoCook/service"
	"encoding/json"
	"log"
	"net/http"
)

func getRecipeFromRequest(r *http.Request) (*model.Recipe, error) {
	var recipe model.Recipe
	err := json.NewDecoder(r.Body).Decode(&recipe)
	if err != nil {
		log.Printf("Can't decode request body to recipe struct: %v", err)
		return nil, err
	}
	return &recipe, nil
}

func CreateRecipe(w http.ResponseWriter, r *http.Request) {
	recipe, err := getRecipeFromRequest(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err := service.CreateRecipe(recipe); err != nil {
		log.Printf("Error calling service CreateRecipe: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(recipe); err != nil {
		log.Printf("Failure encoding value to JSON: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func GetRecipe(w http.ResponseWriter, r *http.Request) {
	id, err := getId(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	recipe, err := service.GetRecipe(id)
	if err != nil {
		log.Printf("Failure retrieving recipe with ID %v: %v", id, err)
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
	recipe, err = service.UpdateRecipe(id, recipe)
	if err != nil {
		log.Printf("Failure updating campaign with ID %v: %v", id, err)
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
	recipe, err := service.DeleteRecipe(id)
	if err != nil {
		log.Printf("Failure deleting recipe with ID %v: %v", id, err)
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
	recipes, err := service.GetRecipes()
	if err != nil {
		log.Printf("Failure retrieving recipes: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if recipes == nil {
		http.Error(w, "404 recipes not found", http.StatusNotFound)
		return
	}
	sendJson(w, recipes)
}