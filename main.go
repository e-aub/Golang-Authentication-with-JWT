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
}
