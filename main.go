package main

import (
	"fmt"
	"os"
	"shive-app/app/routes"
	"shive-app/database"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func init() {
	// Load Env file
	envVariableErr := godotenv.Load(".env")
	if envVariableErr != nil {
		panic("Error in Loading Env Variable")
	}
}

func main() {

	port := os.Getenv("PORT")
	if port == "" {
		panic("PORT not found to start server on")
	}

	// DB
	defer database.DisconnectDB()

	router := routes.Routes()

	//Log events
	router.Use(gin.Logger())

	router.Run(port)
	fmt.Printf("Server Running on Port: %s", port)
}
