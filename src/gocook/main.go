package main

import (
	"GoCook/handler"
	"log"
	"net/http"

	"GoCook/db"

	"github.com/gorilla/mux"
)

func init() {
	defer db.Init()
}

func main() {

	log.Printf("Starting Server")
	router := mux.NewRouter()

	//Recipe routes
	router.HandleFunc("/recipe", handler.CreateRecipe).Methods("POST")
	router.HandleFunc("/recipe/{id}", handler.GetRecipe).Methods("GET")
	router.HandleFunc("/recipe/{id}", handler.UpdateRecipe).Methods("PUT")
	router.HandleFunc("/recipe/{id}", handler.DeleteRecipe).Methods("DELETE")
	router.HandleFunc("/recipes", handler.GetRecipes).Methods("GET")
	router.HandleFunc("/recipes/{ingredient}", handler.GetRecipesByIngredient).Methods("GET")

	//Rating route
	router.HandleFunc("/recipe/{id}/rating", handler.AddRating).Methods("POST")

	//User routes

	//Shopping list routes

	if err := http.ListenAndServe(":8000", router); err != nil {
		log.Fatal(err)
	}
}
