package authz

import future.keywords.in
import future.keywords.every

default allow = false

#Recipes authorization
#Only cooks are allowed to create a recipe
allow {
    input.method == "POST"
    input.path = ["recipe"]
    input.user.IsCook == true
}
#Cooks can only manage their own recipes
allow {
    input.method == "PUT"
    input.path = ["recipe",input.ressource.ID]
    input.ressource.CookID == input.user.ID
}

allow {
    input.method == "DELETE"
    input.path = ["recipe",input.ressource.ID]
    input.ressource.CookID == input.user.ID
}

#Rating authorization
#Only non-cook users can rate recipes
#A recipe can only be rated once by a user
allow {
    input.method == "POST"
    input.path = ["recipe",input.ressource.ID,"rating"]
    input.user.IsCook == false
    every rating in input.ressource.Ratings{
    rating.UserID != input.user.ID
    }
}

#Shopping list authorization
#Only owner of shopping list can add ingredients
allow{
    input.method == "POST"
    input.path = ["shoppingList",input.ressource.ID,"ingredient"]
    input.ressource.UserID == input.user.ID
}

#Only owner of shopping list can load shopping list
allow{
    input.method == "GET"
    input.path = ["shoppingList",input.ressource.ID]
    input.ressource.UserID == input.user.ID
}

#Only owner of shopping list can delete shopping list
allow{
    input.method == "DELETE"
    input.path = ["shoppingList",input.ressource.ID]
    input.ressource.UserID == input.user.ID
}

