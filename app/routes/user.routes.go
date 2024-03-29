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

	userRoutes.GET("/:userId", authMiddlewares.Authenticate, controllers.GetUserDetails)
	userRoutes.GET("/all", authMiddlewares.AuthenticateAdmin, controllers.GetAllUsers)
}
