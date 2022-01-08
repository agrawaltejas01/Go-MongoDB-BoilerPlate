package main

import (
	"log"
	"net/http"

	"github.com/agrawaltejas01/Go-MongoDB-BoilerPlate/app/middleware"
	"github.com/agrawaltejas01/Go-MongoDB-BoilerPlate/app/routes"
	"github.com/agrawaltejas01/Go-MongoDB-BoilerPlate/db"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.Use(middleware.SetWriteHeader)
	router.Use(middleware.PrintReq)

	routes.UserRoutes(router)

	defer db.DisconnectDB()

	log.Fatal(http.ListenAndServe(":8764", router))

}
