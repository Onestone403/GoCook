db = db.getSiblingDB('gocook');
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
        "firstName": "Warren",
        "lastName": "Buffet",
        "isCook":true
    }]
)

db.createCollection('recipes');
db.recipes.insertMany([
    {
        "name":"Apple Pie",
        "ingredients":[
            {
                "name":"Apple",
                "neededAmount":2,
                "unit":"kg"
            },
            {
                "name":"Sugar",
                "neededAmount":1,
                "unit":"kg"
            },
            {
                "name":"Flour",
                "neededAmount":1,
                "unit":"kg"
            },
            {
                "name":"Butter",
                "neededAmount":500,
                "unit":"g"
            }
        ],
        "cookID":ObjectId("000000000000000000000001")
    }, 
    {
        "name":"Pizza",
        "ingredients":[
            {
                "name":"Tomato",
                "neededAmount":2,
                "unit":"kg"
            },
            {
                "name":"Flour",
                "neededAmount":1,
                "unit":"kg"
            },
            {
                "name":"Cheese",
                "neededAmount":1,
                "unit":"kg"
            },
        ],
        "cookID":ObjectId("000000000000000000000003")
    }]);



db.createCollection('shoppingLists');