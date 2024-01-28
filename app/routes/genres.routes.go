package routes

import (
	"shive-app/app/controllers"
	authMiddlewares "shive-app/app/middlewares"

	"github.com/gin-gonic/gin"
)

func GenreRoutes(router *gin.Engine) {
	userRoutes := router.Group("/genre")

	userRoutes.POST("/", authMiddlewares.AuthenticateAdmin, controllers.CreateGenre)
	userRoutes.GET("/:genreId", authMiddlewares.AuthenticateAdmin, controllers.GetGenreDetails)
}
