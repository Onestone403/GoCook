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
    input.path = ["recipe",input.recipe.ID]
    input.recipe.CookID == input.user.ID
}

allow {
    input.method == "DELETE"
    input.path = ["recipe",input.recipe.ID]
    input.recipe.CookID == input.user.ID
}

#Rating authorization
#Only non-cook users can rate recipes
#A recipe can only be rated once by a user
allow {
    input.method == "POST"
    input.path = ["recipe",input.recipe._id,"rating"]
    input.user.IsCook == false
    every rating in input.recipe.ratings{
    rating.userID != input.user.ID
    }
    
}
