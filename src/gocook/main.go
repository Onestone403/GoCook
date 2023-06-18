package main

import (
	"gocook/handler"
	"net/http"
	"os"

	log "github.com/sirupsen/logrus"

	"gocook/db"

	"github.com/gorilla/mux"
)

func init() {
	defer db.Init()

	// init logger
	log.SetFormatter(&log.TextFormatter{
		ForceColors: true,
	})
	log.SetReportCaller(true)
	level, err := log.ParseLevel(os.Getenv("LOG_LEVEL"))
	if err != nil {
		log.Info("Log level not specified, using default log level: INFO")
		log.SetLevel(log.InfoLevel)
		return
	}
	log.SetLevel(level)
}

func main() {

	log.Printf("Starting Server")
	router := mux.NewRouter()

	//Recipe routes
	router.HandleFunc("/recipe", verifyJWT(handler.CreateRecipe)).Methods("POST")
	router.HandleFunc("/recipe/{id}", verifyJWT(handler.GetRecipe)).Methods("GET")
	router.HandleFunc("/recipe/{id}", verifyJWT(handler.UpdateRecipe)).Methods("PUT")
	router.HandleFunc("/recipe/{id}", verifyJWT(handler.DeleteRecipe)).Methods("DELETE")
	router.HandleFunc("/recipes", verifyJWT(handler.GetRecipesByIngredient)).Queries("ingredient", "{ingredient}").Methods("GET")
	router.HandleFunc("/recipes", verifyJWT(handler.GetRecipes)).Methods("GET")

	//Rating route
	router.HandleFunc("/recipe/{id}/rating", verifyJWT(handler.AddRating)).Methods("POST")

	//Shopping list routes
	router.HandleFunc("/shoppingList", verifyJWT(handler.CreateShoppingList)).Queries("title", "{title}").Methods("POST")
	router.HandleFunc("/shoppingList/{id}", verifyJWT(handler.GetShoppingList)).Methods("GET")
	router.HandleFunc("/shoppingList/{id}/ingredient", verifyJWT(handler.AddIngredient)).Methods("POST")
	router.HandleFunc("/shoppingLists", verifyJWT(handler.GetShoppingLists)).Methods("GET")
	router.HandleFunc("/shoppingList/{id}", verifyJWT(handler.DeleteShoppingList)).Methods("DELETE")

	if err := http.ListenAndServe(":8000", router); err != nil {
		log.Fatal(err)
	}
}
