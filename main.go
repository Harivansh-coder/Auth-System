package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// album represents data about a record album.
type user struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func main() {
	router := gin.Default()
	router.GET("/", defaultRoute)

	router.Run("localhost:8080")
}

func defaultRoute(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, gin.H{"message": "This is the default route for Authentication API"})
}
