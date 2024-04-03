package main

import (
	"os"

	"github.com/e-aub/Golang-Authentication-with-JWT.git/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	router := gin.New()
	router.Use(gin.Logger())
	routes.authRoutes(router)
	routes.userRoutes(router)
	router.GET("/api-1", func(c *gin.Context) {
		c.JSON(200, gin.H{"Success": "Access granted for api-1"})
	})
	router.GET("/api-2", func(c *gin.Context) {
		c.JSON(200, gin.H{"success": "Access granted for api-2"})
	})
	router.Run(":" + port)
}
