package routes

import (
	"shive-app/app/controllers"
	authMiddlewares "shive-app/app/middlewares"

	"github.com/gin-gonic/gin"
)

func MovieRoutes(router *gin.Engine) {
	movieRoutes := router.Group("/movie")

	movieRoutes.POST("/", authMiddlewares.AuthenticateAdmin, controllers.CreateMovie)
	movieRoutes.PUT("/", authMiddlewares.AuthenticateAdmin, controllers.UpdateMovieDetails)
	movieRoutes.DELETE("/:movieId", authMiddlewares.AuthenticateAdmin, controllers.DeleteMovie)

	movieRoutes.GET("/:movieId", authMiddlewares.Authenticate, controllers.GetMovieDetails)
	movieRoutes.GET("/all", authMiddlewares.Authenticate, controllers.GetAllMovieDetails)
	movieRoutes.GET("/searchByName", authMiddlewares.Authenticate, controllers.SearchMoviesByName)
	movieRoutes.GET("/searchByGenreId/:genreId", authMiddlewares.Authenticate, controllers.SearchMoviesByGenreId)

}
