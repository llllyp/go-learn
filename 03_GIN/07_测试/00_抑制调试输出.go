package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"encoding/json"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

type User struct {
	Username string `json:"username"`
	Gender string `json:"gender"`
}

func setupRouter() *gin.Engine {
	router := gin.Default()
	router.GET("/ping", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "pong")
	})
	return router
}

func postUser(router *gin.Engine) *gin.Engine {
	router.POST("/user/add", func(ctx *gin.Context) {
		var user User
		ctx.BindJSON(&user)
		ctx.JSON(http.StatusOK, user)
	})
	return router
}

func main() {
	router := setupRouter()
	router = postUser(router)
	router.Run()
}

func TestPingRoute(test *testing.T) {
	router := setupRouter()

	record := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/ping", nil)

	router.ServeHTTP(record, req)

	assert.Equal(test, 200, record.Code)
	assert.Equal(test, "pong", record.Body.String())
}

// Test for POST /user/add
func TestPostUser(t *testing.T) {
  router := setupRouter()
  router = postUser(router)

  w := httptest.NewRecorder()

  // Create an example user for testing
  exampleUser := User{
    Username: "test_name",
    Gender:   "male",
  }
  userJson, _ := json.Marshal(exampleUser)
  req, _ := http.NewRequest("POST", "/user/add", strings.NewReader(string(userJson)))
  router.ServeHTTP(w, req)

  assert.Equal(t, 200, w.Code)
  // Compare the response body with the json data of exampleUser
  assert.Equal(t, string(userJson), w.Body.String())
}