package test

import (
	"bytes"
	"context"
	"fmt"
	"gocook/db"
	"gocook/handler"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type RecipeTestSuite struct {
	suite.Suite
	CloseTestDB func()
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

func (suite *RecipeTestSuite) TestCreateRecipe() {
	fmt.Println("-- From TestCreateRecipe")

	json_data := `{"name":"Pommes","ingredients":[{"name":"Potato","neededAmount":2.0,"unit":"KILOGRAMS"},
	{"name":"FryFat","neededAmount":4.0,"unit":"LITER"}],"CookId":"000000000000000000000001"}`
	rr := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodPost, "/dummy-url", bytes.NewBufferString(json_data))
	req.Header.Set("Content-Type", "application/json")
	ctx := context.WithValue(req.Context(), "userID", "000000000000000000000001")
	ctx = context.WithValue(ctx, "path", []string{"recipe"})
	ctx = context.WithValue(ctx, "method", req.Method)

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

func (suite *RecipeTestSuite) TestGetRecipes() {
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

}

func TestRecipeTestSuite(t *testing.T) {
	suite.Run(t, new(RecipeTestSuite))
}
