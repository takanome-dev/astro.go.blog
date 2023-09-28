package routes

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/takanome-dev/blog-with-astro-golang/internal/auth"
	"github.com/takanome-dev/blog-with-astro-golang/pkg/controllers"
)

var PostsRoutes = func (router *mux.Router)  {
	router.HandleFunc("/posts", controllers.GetAllPosts).Methods("GET")
	router.HandleFunc("/posts/{id}", controllers.GetPostByID).Methods("GET")
	// TODO: remove middleware here
	router.HandleFunc("/posts/users/{userId}", auth.Middleware(
		http.HandlerFunc(controllers.GetPostsByUserID), 
		auth.AuthMiddleware,
	).ServeHTTP,).Methods("GET")
	router.HandleFunc(
		"/posts/current-user",
		auth.Middleware(
			http.HandlerFunc(controllers.GetPostsForLoggedInUser), 
			auth.AuthMiddleware,
		).ServeHTTP,
	).Methods("GET")
	router.HandleFunc(
		"/posts", 
		auth.Middleware(
			http.HandlerFunc(controllers.CreatePost), 
			auth.AuthMiddleware,
		).ServeHTTP,
	).Methods("POST")
	router.HandleFunc(
		"/posts/{id}", 
		auth.Middleware(
			http.HandlerFunc(controllers.UpdatePost), 
			auth.AuthMiddleware,
		).ServeHTTP,
	).Methods("PUT")
	router.HandleFunc(
		"/posts/{id}", 
		auth.Middleware(
			http.HandlerFunc(controllers.DeletePost), 
			auth.AuthMiddleware,
		).ServeHTTP,
	).Methods("DELETE")
}