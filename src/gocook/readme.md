# **GoCook**

### **Summary**:
GoCook is a recipe app for hobby cooks and professionals, to provide their recipes to a broad audience. 
Users can search for recipes and rate them. Additionaly it's possible to create shopping list with ingredients of the recipes.

The final application consists of four docker containers:

- GoCook 
- MongoDB
- Open Policy Agent
- Bundle-Server


### **Authentication**:
The API requires a authenticated user to acces the individula endpoints.
As part of this project, only the verification of a JWT is implemented.

### **Authorization**:
Some of the endpoints require a fine grained authorization, which is implemented via Open Policy Agent. [Info to authorization in this app](../authorization/readme.md)

### **Persistence**:
A dedicated mongoDB container is used to persist the data of the API.

### **Endpoints**:

All API endpoints can be tested via the postman request collection.
Each request can be sent with one of three possible JWTs to simulate the different users. 

The following endppoints are provided:

### Recipes:
|Method|URL|Description|
|---|---|---|
|POST|/recipe| creates a new recipe|
|GET|/recipe{id}| returns the specified recipe|
|PUT|/recipe{id}| updates the recipe via the recipe provided by the request body|
|DELETE|/recipe{id}| deletes the specified recipe|
|GET|/recipes| without param ingredient: returns all recipes </br> with param ingredient: returns all recipes where ingredients contain the specified ingredient (case-sensitiv) |

### Shopping List:
|Method|URL|Description|
|---|---|---|
|POST|/shoppingList|creates a shopping list with a specific title (must be provided via param title)|
|GET|/shoppingList{id}| returns the specified shopping list|
|POST|/shoppingList/{id}/ingredient| adds the ingredient provided via the request body to the shopping list|
|DELETE|/shoppingList/{id}| deletes the specified shopping list|
|GET|/shoppingLists| returns all shopping lists for the authenticated user |



## **Get started**

To start the application just download the code, cd into the root directory of the repo and start the docker containers:

`docker-compose up --build`