package routes

import (
	"github.com/gorilla/mux"
	"github.com/takanome-dev/blog-with-astro-golang/pkg/controllers"
)

var AuthRoute = func (router *mux.Router) {
	router.HandleFunc("/auth/login", controllers.Login).Methods("POST")
	router.HandleFunc("/auth/register", controllers.Register).Methods("POST")
}