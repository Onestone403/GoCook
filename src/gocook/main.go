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

	//Recipe routes (private as user must be logged in to access)
	router.HandleFunc("/recipe", verifyJWT(handler.CreateRecipe)).Methods("POST")
	router.HandleFunc("/recipe/{id}", verifyJWT(handler.GetRecipe)).Methods("GET")
	router.HandleFunc("/recipe/{id}", verifyJWT(handler.UpdateRecipe)).Methods("PUT")
	router.HandleFunc("/recipe/{id}", verifyJWT(handler.DeleteRecipe)).Methods("DELETE")
	router.HandleFunc("/recipes", verifyJWT(handler.GetRecipes)).Methods("GET")
	router.HandleFunc("/recipes/{ingredient}", verifyJWT(handler.GetRecipesByIngredient)).Methods("GET")

	//Rating route (private as user must be logged in to access)
	router.HandleFunc("/recipe/{id}/rating", verifyJWT(handler.AddRating)).Methods("POST")

	//User routes

	//Shopping list routes

	if err := http.ListenAndServe(":8000", router); err != nil {
		log.Fatal(err)
	}
}
