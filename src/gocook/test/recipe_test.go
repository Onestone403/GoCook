package test

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"gocook/db"
	"gocook/handler"
	"gocook/model"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type RecipeTestSuite struct {
	suite.Suite
	CloseTestDB func()
	recipe      *model.Recipe
}

func (suite *RecipeTestSuite) SetupSuite() {
	fmt.Println("-- From SetupSuite")
	fmt.Print("Setting up recipe test suite")

	//Start Mongo Container
	suite.CloseTestDB = db.SetupTestDB(suite.T())

}
func (suite *RecipeTestSuite) TearDownSuite() {
	fmt.Println("-- From TearDownSuite")
	suite.CloseTestDB()

}

func (suite *RecipeTestSuite) Test_01_CreateRecipe() {
	fmt.Println("-- From TestCreateRecipe")

	json_data := `{"name":"Pommes","ingredients":[{"name":"Potato","neededAmount":2.0,"unit":"KILOGRAMS"},
	{"name":"FryFat","neededAmount":4.0,"unit":"LITER"}],"CookId":"000000000000000000000001"}`
	rr := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodPost, "/dummy-url", bytes.NewBufferString(json_data))
	req.Header.Set("Content-Type", "application/json")
	ctx := context.WithValue(req.Context(), "istest", true)

	req = req.WithContext(ctx)
	handler := http.HandlerFunc(handler.CreateRecipe)

	handler.ServeHTTP(rr, req)

	if status := rr.Result().StatusCode; status != http.StatusOK {
		msg := fmt.Sprintf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
		assert.Fail(suite.T(), msg)
		return
	}
	fmt.Printf("Test passed with recipe: %v", rr.Body.String())

}
func (suite *RecipeTestSuite) Test_02_GetRecipes() {
	fmt.Println("-- From TestGetRecipes")
	rr := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/dummy-url", nil)

	handler := http.HandlerFunc(handler.GetRecipes)

	handler.ServeHTTP(rr, req)
	if status := rr.Result().StatusCode; status != http.StatusOK {
		msg := fmt.Sprintf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
		assert.Fail(suite.T(), msg)
		return
	}
	fmt.Printf("Test passed with recipes received: %v", rr.Body.String())
	var recipes []model.Recipe
	err := json.Unmarshal([]byte(rr.Body.String()), &recipes)
	if err != nil {
		fmt.Print(err)
		assert.Fail(suite.T(), "Error unmarshalling recipe")
	}
	suite.recipe = &recipes[0]
	assert.Equal(suite.T(), "Pommes", suite.recipe.Name)
}

func (suite *RecipeTestSuite) Test_03_GetRecipe() {
	fmt.Println("-- From TestGetRecipe")
	rr := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/recipe/:id", nil)
	fmt.Printf("Recipe ID: %v", suite.recipe.ID)
	req = req.WithContext(context.WithValue(req.Context(), "id", suite.recipe.ID))

	handler := http.HandlerFunc(handler.GetRecipe)

	handler.ServeHTTP(rr, req)
	if status := rr.Result().StatusCode; status != http.StatusOK {
		msg := fmt.Sprintf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
		assert.Fail(suite.T(), msg)
		return
	}
	fmt.Printf("Test passed with recipe received: %v", rr.Body.String())
	var recipe model.Recipe
	err := json.Unmarshal([]byte(rr.Body.String()), &recipe)
	if err != nil {
		fmt.Print(err)
		assert.Fail(suite.T(), "Error unmarshalling recipe")
	}
	assert.Equal(suite.T(), "Pommes", recipe.Name)
}

func TestRecipeTestSuite(t *testing.T) {
	suite.Run(t, new(RecipeTestSuite))
}
