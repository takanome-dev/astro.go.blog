package routes

import (
	"github.com/gorilla/mux"
	"github.com/takanome-dev/blog-with-astro-golang/pkg/controllers"
)

var PostsRoutes = func (router *mux.Router)  {
	router.HandleFunc("/posts", controllers.GetAllPosts).Methods("GET")
	router.HandleFunc("/posts/{id}", controllers.GetPostByID).Methods("GET")
	router.HandleFunc("/posts", controllers.CreatePost).Methods("POST")
	router.HandleFunc("/posts/{id}", controllers.UpdatePost).Methods("PUT")
	router.HandleFunc("/posts/{id}", controllers.DeletePost).Methods("DELETE")
}