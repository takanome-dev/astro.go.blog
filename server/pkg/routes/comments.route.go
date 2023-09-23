package routes

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/takanome-dev/blog-with-astro-golang/internal/auth"
	"github.com/takanome-dev/blog-with-astro-golang/pkg/controllers"
)

var CommentsRoutes = func (router *mux.Router)  {
	router.HandleFunc("/comments", controllers.GetAllComments).Methods("GET")
	router.HandleFunc("/comments/{id}", controllers.GetCommentByID).Methods("GET")
	router.HandleFunc(
		"/comments", 
		auth.Middleware(
			http.HandlerFunc(controllers.CreateComment), 
			auth.AuthMiddleware,
		).ServeHTTP,
	).Methods("POST")
	router.HandleFunc(
		"/comments/{id}", 
		auth.Middleware(
			http.HandlerFunc(controllers.UpdateComment), 
			auth.AuthMiddleware,
		).ServeHTTP,
	).Methods("PUT")
	router.HandleFunc(
		"/comments/{id}", 
		auth.Middleware(
			http.HandlerFunc(controllers.DeleteComment), 
			auth.AuthMiddleware,
		).ServeHTTP,
	).Methods("DELETE")
}