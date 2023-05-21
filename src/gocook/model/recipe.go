package model

type Recipe struct {
	Id          uint
	Name        string
	Ingredients []Ingredient
	CookId      int
	Ratings     []Rating
}
