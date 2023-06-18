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
        "firstName": "Mark",
        "lastName": "Zuckerberg",
        "isCook":true
    }]
)

db.createCollection('recipes');
db.createCollection('shoppingLists');