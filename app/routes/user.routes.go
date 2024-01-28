package routes

import (
	"shive-app/app/controllers"
	authMiddlewares "shive-app/app/middlewares"

	"github.com/gin-gonic/gin"
)

func UserRoutes(router *gin.Engine) {
	userRoutes := router.Group("/user")

	userRoutes.POST("/signup", controllers.Signup)
	userRoutes.POST("/login", controllers.Login)

	userRoutes.Use(authMiddlewares.Authenticate)
	userRoutes.GET("/:userId", controllers.GetDetails)
	userRoutes.GET("/all", controllers.GetAllUsers)
}
