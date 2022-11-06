package main

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"harry/auth_system/routers"
)

// album represents data about a record album.

func main() {
	router := gin.Default()
	router.GET("/", defaultRoute)

	routers.UserRouters(router)

	// router.GET("/users", getUsers)
	// router.GET("/login/:id", getUsers)

	router.Run("localhost:8080")
}

func defaultRoute(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, gin.H{"message": "This is the default route for Authentication API"})
}
