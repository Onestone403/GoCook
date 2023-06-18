package handler

import (
	"gocook/service"
	"net/http"

	log "github.com/sirupsen/logrus"
)

func CreateShoppingList(w http.ResponseWriter, r *http.Request) {
	title, err := getTitle(r)
	shoppingList, err := service.CreateShoppingList(r.Context(), title)
	if err != nil {
		log.Errorf("Failure creating shopping list: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	sendJson(w, shoppingList)
}
func GetShoppingList(w http.ResponseWriter, r *http.Request) {
	id, err := getId(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	shoppingList, err := service.GetShoppingList(r.Context(), id)
	if err != nil {
		log.Errorf("Failure retrieving shopping list with ID %v: %v", id, err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if shoppingList == nil {
		http.Error(w, "404 shopping list not found", http.StatusNotFound)
		return
	}
	sendJson(w, shoppingList)
}

func GetShoppingLists(w http.ResponseWriter, r *http.Request) {
	shoppingLists, err := service.GetShoppingLists(r.Context())
	if err != nil {
		log.Errorf("Failure retrieving shopping list: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if shoppingLists == nil {
		http.Error(w, "404 no shopping list for user", http.StatusNotFound)
		return
	}
	sendJson(w, shoppingLists)
}

func DeleteShoppingList(w http.ResponseWriter, r *http.Request) {
	id, err := getId(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	shoppingList, err := service.DeleteShoppingList(r.Context(), id)
	if err != nil {
		log.Errorf("Failure deleting shopping list with ID %v: %v", id, err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if shoppingList == nil {
		http.Error(w, "404 shopping list not found", http.StatusNotFound)
		return
	}
}
