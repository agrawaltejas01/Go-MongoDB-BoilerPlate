package routes

import "github.com/gin-gonic/gin"

func Routes() *gin.Engine {
	router := gin.Default()

	UserRoutes(router)
	GenreRoutes(router)

	return router

}
