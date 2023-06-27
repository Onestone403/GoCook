package db

import (
	"context"
	"fmt"
	"testing"

	"github.com/ory/dockertest/v3"
	"github.com/ory/dockertest/v3/docker"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func SetupTestDB(t *testing.T) func() {
	pool, err := dockertest.NewPool("")
	if err != nil {
		t.Fatalf("Could not connect to docker: %s", err)
	}

	runDockerOpt := &dockertest.RunOptions{
		Repository: "mongo",
		Tag:        "4.4.6",
		PortBindings: map[docker.Port][]docker.PortBinding{
			"27017/tcp": {{HostIP: "localhost", HostPort: "27017"}},
		},
	}

	fnConfig := func(config *docker.HostConfig) {
		config.AutoRemove = true
		config.RestartPolicy = docker.NeverRestart()
	}

	resource, err := pool.RunWithOptions(runDockerOpt, fnConfig)
	if err != nil {
		t.Fatalf("Could not start test DB: %s", err)
	}

	err = pool.Retry(func() error {
		clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
		client, err := mongo.Connect(context.Background(), clientOptions)
		if err != nil {
			return err
		}
		err = client.Ping(context.TODO(), nil)
		return err
	})
	if err != nil {
		t.Fatalf("Could not connect to test DB: %s", err)
	}

	fmt.Println("Test DB started")

	err = initTestDatabase()
	if err != nil {
		t.Fatalf("Could not init test DB: %s", err)
	}

	return func() { resource.Close() }
}

func initTestDatabase() error {
	fmt.Print("Initializing TestDB with users...\n")
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		fmt.Print(err)
		return err
	}

	err = client.Ping(context.Background(), nil)
	if err != nil {
		fmt.Print(err)
		return err
	}
	fmt.Print("Connected to MongoDB!\n")

	goCookDatabase = client.Database("gocook")
	RecipeCollection = goCookDatabase.Collection("recipes")
	UserCollection = goCookDatabase.Collection("users")
	ShoppingListCollection = goCookDatabase.Collection("shoppingLists")

	id1, _ := primitive.ObjectIDFromHex("000000000000000000000001")
	id2, _ := primitive.ObjectIDFromHex("000000000000000000000002")
	id3, _ := primitive.ObjectIDFromHex("000000000000000000000003")

	_, err = UserCollection.InsertMany(context.Background(), []interface{}{
		bson.M{
			"_id":       id1,
			"firstName": "Tim",
			"lastName":  "Koch",
			"isCook":    true,
		},
		bson.M{
			"_id":       id2,
			"firstName": "Elon",
			"lastName":  "Muskat",
			"isCook":    false,
		},
		bson.M{
			"_id":       id3,
			"firstName": "Warren",
			"lastName":  "Buffet",
			"isCook":    true,
		},
	})
	if err != nil {
		fmt.Printf("Insert error %v \n", err)
		return err
	}
	return nil
}
