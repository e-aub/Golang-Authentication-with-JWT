package routes

import (
	"github.com/gin-gonic/gin"
)

func authRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.POST("/users/signup", controller.)
	incomingRoutes.POST("/users/login", controller.)
}
