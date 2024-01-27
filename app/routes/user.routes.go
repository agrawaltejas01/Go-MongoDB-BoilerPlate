package routes

import (
	"shive-app/app/controllers"

	"github.com/gin-gonic/gin"
)

func UserRoutes(router *gin.Engine) {
	userRoutes := router.Group("/user")

	userRoutes.POST("/signup", controllers.Signup)
	userRoutes.POST("/login", controllers.Login)
}
