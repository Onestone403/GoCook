#!/bin/bash
set -e

mongosh <<EOF
use gocook

db.createCollection('users');
db.users.insertMany([
    {
        "_id": ObjectId("000000000000000000000001"),
        "firstName": "Tim",
        "lastName": "Koch",
        "isCook":true
    },
    {
        "_id": ObjectId("000000000000000000000002"),
        "firstName": "Elon",
        "lastName": "Muskat",
        "isCook":false
    },
    {
        "_id": ObjectId("000000000000000000000003"),
        "firstName": "Mark",
        "lastName": "Zuckerberg",
        "isCook":true
    }]
)

db.createCollection('recipes');
db.recipes.insertMany([
    {
        "name":"Apple Pie",
        "ingredients":[
            {"name":"Apple","neededAmount":4.0,"unit":"KILOGRAMS"},
            {"name":"Sugar","neededAmount":8,"unit":"TABLESPOON"},
            {"name":"Flour","neededAmount":2,"unit":"KILOGRAMS"},
            {"name":"Butter","neededAmount":1,"unit":"KILOGRAMS"},
            {"name":"Cinnamon","neededAmount":1,"unit":"TEASPOON"},
            {"name":"Salt","neededAmount":1,"unit":"TEASPOON"},
            {"name":"Water","neededAmount":1,"unit":"LITER"}],
        "cookId":"000000000000000000000001"
}]
)
EOF