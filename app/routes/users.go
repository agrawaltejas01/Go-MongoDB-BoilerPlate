package routes

import (
	"github.com/agrawaltejas01/Go-MongoDB-BoilerPlate/app/controllers"
	"github.com/gorilla/mux"
)

func UserRoutes(router *mux.Router) {
	router.HandleFunc("/user", controllers.GetUsers).Methods("GET")
	router.HandleFunc("/user/{id}", controllers.GetUser).Methods("GET")
	router.HandleFunc("/user", controllers.AddUser).Methods("POST")
	router.HandleFunc("/user/{id}", controllers.UpdateUser).Methods("PATCH")
	router.HandleFunc("/user/{id}", controllers.DeleteUser).Methods("DELETE")
}
