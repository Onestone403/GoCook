package model

type Recipe struct {
	Id          int
	Name        string
	Ingredients []Ingredient
	CookId      int
	Ratings     []int
}
