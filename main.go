package main

import (
	"fmt"
	"os"
	"shive-app/database"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {

	// Load Env file
	envVariableErr := godotenv.Load(".env")
	if envVariableErr != nil {
		panic("Error in Loading Env Variable")
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}
	router := gin.Default()

	// Connect DB
	database.ConnectDB()
	defer database.DisconnectDB()

	//Log events
	router.Use(gin.Logger())

	router.Run(port)
	fmt.Printf("Server Running on Port: %s", port)
}
