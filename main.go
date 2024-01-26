package main

import (
	"fmt"
	"os"
	"shive-app/app/routes"
	"shive-app/database"

	"github.com/gin-gonic/gin"
)

func main() {

	// Load Env file
	// envVariableErr := godotenv.Load(".env")
	// if envVariableErr != nil {
	// 	panic("Error in Loading Env Variable")
	// }

	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}

	// DB
	defer database.DisconnectDB()

	router := routes.Routes()

	//Log events
	router.Use(gin.Logger())

	router.Run(port)
	fmt.Printf("Server Running on Port: %s", port)
}
