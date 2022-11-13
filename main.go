package main

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"harry/auth_system/routers"

	config "harry/auth_system/configs"
)

// album represents data about a record album.

func main() {
	router := gin.Default()

	// database
	config.ConnectDB()

	// routers
	router.GET("/", defaultRoute)

	routers.UserRouters(router)

	routers.LoginRouters(router)

	router.Run("localhost:8080")

}

func defaultRoute(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "This is the default route for Authentication API"})
}
