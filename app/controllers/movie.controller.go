package controllers

import (
	movieService "shive-app/app/service/movie"
	"shive-app/database/models"
	serverResponse "shive-app/lib/server-response"
	"strconv"

	"github.com/gin-gonic/gin"
)

func returnMovieData(movieObject models.Movie) map[string]interface{} {
	var userData = make(map[string]interface{})
	userData["name"] = movieObject.Name
	userData["movie_id"] = movieObject.Movie_id
	userData["topic"] = movieObject.Topic
	userData["movie_url"] = movieObject.Movie_url

	return userData
}

func CreateMovie(context *gin.Context) {
	var movie models.Movie
	jsonBindErr := context.BindJSON(&movie)

	if jsonBindErr != nil {
		serverResponse.BadRequestServerError(context, "Bad Data received for movie object "+jsonBindErr.Error())
		return
	}

	id, creationErr := movieService.CreateMovie(movie)

	if creationErr != nil {
		serverResponse.InternalServerError(context, "Error in Creating Movie - "+creationErr.Error())
		return
	}

	var result = make(map[string]interface{})
	result["id"] = id

	serverResponse.SuccessResponse(context, result, 0)
}

func GetAllMovieDetails(context *gin.Context) {
	movies, err := movieService.GetAllMovies()

	if err != nil {
		serverResponse.InternalServerError(context, "Error in Getting Movie Data - "+err.Error())
		return
	}

	var movieSlice []interface{}
	for _, user := range movies {
		movieSlice = append(movieSlice, returnMovieData(user))
	}

	var result = make(map[string]interface{})
	result["data"] = movieSlice

	serverResponse.SuccessResponse(context, result, 0)
}

func GetMovieDetails(context *gin.Context) {
	movieId := context.Param("movieId")
	movieIdInt, strToIntErr := strconv.Atoi(movieId)

	if strToIntErr != nil {
		serverResponse.BadRequestServerError(context, "Invalid Movie Id passed")
		return
	}

	movieObject, err := movieService.GetMovieDataFromMovieId(movieIdInt)

	if err != nil {
		serverResponse.InternalServerError(context, "Error in Getting Movie Data - "+err.Error())
		return
	}

	var result = make(map[string]interface{})
	result["movie"] = returnMovieData(movieObject)
	serverResponse.SuccessResponse(context, result, 0)
}

func UpdateMovieDetails(context *gin.Context) {
	var movie models.Movie
	jsonBindErr := context.BindJSON(&movie)

	if jsonBindErr != nil {
		serverResponse.BadRequestServerError(context, "Bad Data received for movie object "+jsonBindErr.Error())
		return
	}

	err := movieService.UpdateMovie(movie)

	if err != nil {
		serverResponse.InternalServerError(context, "Error in Updating Movie Data - "+err.Error())
		return
	}

	var result = make(map[string]interface{})
	result["success"] = true
	serverResponse.SuccessResponse(context, result, 0)
}

func DeleteMovie(context *gin.Context) {
	movieId := context.Param("movieId")
	movieIdInt, strToIntErr := strconv.Atoi(movieId)

	if strToIntErr != nil {
		serverResponse.BadRequestServerError(context, "Invalid Movie Id passed")
		return
	}

	err := movieService.DeleteMovie(movieIdInt)

	if err != nil {
		serverResponse.InternalServerError(context, "Error in Deleting Movie Data - "+err.Error())
		return
	}

	var result = make(map[string]interface{})
	result["success"] = true
	serverResponse.SuccessResponse(context, result, 0)
}

func SearchMoviesByName(context *gin.Context) {
	name := context.Query("name")

	movies, err := movieService.SearchMoviesByNamePattern(name)

	if err != nil {
		serverResponse.InternalServerError(context, "Error in Getting Movie Data - "+err.Error())
		return
	}

	var movieSlice []interface{}
	for _, user := range movies {
		movieSlice = append(movieSlice, returnMovieData(user))
	}

	var result = make(map[string]interface{})
	result["data"] = movieSlice
	serverResponse.SuccessResponse(context, result, 0)
}

func SearchMoviesByGenreId(context *gin.Context) {
	movieId := context.Param("genreId")
	movieIdInt, strToIntErr := strconv.Atoi(movieId)

	if strToIntErr != nil {
		serverResponse.BadRequestServerError(context, "Invalid Movie Id passed")
		return
	}

	movies, err := movieService.SearchMovieByGenreId(movieIdInt)

	if err != nil {
		serverResponse.InternalServerError(context, "Error in Getting Movie Data - "+err.Error())
		return
	}

	var movieSlice []interface{}
	for _, user := range movies {
		movieSlice = append(movieSlice, returnMovieData(user))
	}

	var result = make(map[string]interface{})
	result["data"] = movieSlice
	serverResponse.SuccessResponse(context, result, 0)
}
