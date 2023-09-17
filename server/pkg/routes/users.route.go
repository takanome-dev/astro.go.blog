package routes

import (
	"github.com/gorilla/mux"
	"github.com/takanome-dev/blog-with-astro-golang/pkg/controllers"
)

var UsersRoute = func (router *mux.Router) {
	router.HandleFunc("/users", controllers.GetAllUsers).Methods("GET")
	router.HandleFunc("users/{id}", controllers.GetUserById).Methods("GET")
	router.HandleFunc("users", controllers.CreateUser).Methods("POST")
	router.HandleFunc("users/{id}", controllers.UpdateUser).Methods("PUT")
	router.HandleFunc("users/{id}", controllers.DeleteUser).Methods("DELETE")
}