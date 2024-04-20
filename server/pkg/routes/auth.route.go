package routes

import (
	"github.com/gorilla/mux"
	"github.com/takanome-dev/astro.go.blog/pkg/controllers"
)

var AuthRoute = func (router *mux.Router) {
	router.HandleFunc("/auth/login", controllers.Login).Methods("POST")
	router.HandleFunc("/auth/register", controllers.Register).Methods("POST")
  router.HandleFunc("/auth/reset-password", controllers.ResetPassword).Methods("POST")
}
