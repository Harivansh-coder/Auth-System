package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	config "harry/auth_system/configs"
)

// album represents data about a record album.
type user struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

// albums slice to seed record album data.
var users = []user{
	{ID: "1", Name: "John", Email: "", Password: ""},
	{ID: "1", Name: "John", Email: "", Password: ""},
	{ID: "1", Name: "John", Email: "", Password: ""}}

func main() {
	router := gin.Default()
	router.GET("/", defaultRoute)
	router.GET("/users", getUsers)
	router.GET("/login/:id", getUsers)

	router.Run("localhost:8080")
}

func defaultRoute(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, gin.H{"message": "This is the default route for Authentication API"})
}

func getUsers(c *gin.Context) {
	fmt.Println(config.EnvMongoURI())
	c.IndentedJSON(http.StatusOK, users)
}
