package routes

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/takanome-dev/blog-with-astro-golang/internal/auth"
	"github.com/takanome-dev/blog-with-astro-golang/pkg/controllers"
)

var UsersRoute = func (router *mux.Router) {
	router.HandleFunc(
		"/users",
		auth.Middleware(
			http.HandlerFunc(controllers.GetAllUsers), 
			auth.AuthMiddleware,
		).ServeHTTP,
		).Methods("GET")
	router.HandleFunc(
		"/users/{id}", 
		auth.Middleware(
			http.HandlerFunc(controllers.GetUserById), 
			auth.AuthMiddleware,
		).ServeHTTP,
	).Methods("GET")
	router.HandleFunc(
		"/users/username/{username}", 
		auth.Middleware(
			http.HandlerFunc(controllers.GetUserByUsername), 
			auth.AuthMiddleware,
		).ServeHTTP,
	).Methods("GET")
	router.HandleFunc(
		"/users/{id}", 
		auth.Middleware(
			http.HandlerFunc(controllers.UpdateUser), 
			auth.AuthMiddleware,
		).ServeHTTP,
	).Methods("PUT")
	router.HandleFunc(
		"/users/{id}", 
		auth.Middleware(
			http.HandlerFunc(controllers.DeleteUser), 
			auth.AuthMiddleware,
		).ServeHTTP,
	).Methods("DELETE")
}