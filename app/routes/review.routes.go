package routes

import (
	"shive-app/app/controllers"
	authMiddlewares "shive-app/app/middlewares"

	"github.com/gin-gonic/gin"
)

func ReviewRoutes(router *gin.Engine) {
	reviewRoutes := router.Group("/review")

	reviewRoutes.Use(authMiddlewares.Authenticate)

	reviewRoutes.POST("/", controllers.CreateReview)
	reviewRoutes.DELETE("/:reviewId", controllers.DeleteReview)

	reviewRoutes.GET("/:movieId", controllers.GetReviewsForMovieId)
	reviewRoutes.GET("/user", controllers.GetAllUserReviews)

}
