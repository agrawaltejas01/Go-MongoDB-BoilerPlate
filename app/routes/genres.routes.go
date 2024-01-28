package routes

import (
	"shive-app/app/controllers"
	authMiddlewares "shive-app/app/middlewares"

	"github.com/gin-gonic/gin"
)

func GenreRoutes(router *gin.Engine) {
	genreRoutes := router.Group("/genre")

	genreRoutes.Use(authMiddlewares.AuthenticateAdmin)

	genreRoutes.POST("/", controllers.CreateGenre)
	genreRoutes.GET("/:genreId", controllers.GetGenreDetails)
	genreRoutes.GET("/all", controllers.GetAllGenreDetails)
	genreRoutes.PUT("/", controllers.UpdateGenreDetails)
	genreRoutes.DELETE("/:genreId", controllers.DeleteGenre)
}
