# GoCook

GoCook is a recipe app for hobby cooks and professionals, to provide their recipes to a broad audience. 
Users can search for recipes and rate them. Additionaly it's possible to create shopping list with ingredients of the recipes.


The API requires a authenticated user to acces the individula endpoints.
As part of this project, only the verification of a JWT is implemented.

Some of the endpoints require a fine grained authorization, which is implemented via Open Policy Agent. 

A dedicated mongoDB container is used to persist the data of the API.

All API endpoints can be tested via the postman request collection.
Each request can be sent with one of three possible JWTs to simulate the different users. 

To start the service just download the code, cd into the root directory of the repo and start the docker containers:

`docker-compose up --build`

