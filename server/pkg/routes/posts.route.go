package routes

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/takanome-dev/astro.go.blog/internal/auth"
	"github.com/takanome-dev/astro.go.blog/pkg/controllers"
)

var PostsRoutes = func (router *mux.Router)  {
	router.HandleFunc("/posts", controllers.GetAllPosts).Methods("GET")
	router.HandleFunc("/posts/{id}", controllers.GetPostByID).Methods("GET")
	router.HandleFunc("/posts/users/{userId}", controllers.GetPostsByUserID).Methods("GET")
	router.HandleFunc(
		"/posts/current-user/published",
		auth.Middleware(
			http.HandlerFunc(controllers.GetCurrentUserPosts), 
			auth.AuthMiddleware,
		).ServeHTTP,
	).Methods("GET")
	router.HandleFunc(
		"/posts/current-user/drafts",
		auth.Middleware(
			http.HandlerFunc(controllers.GetCurrentUserDraftPosts), 
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