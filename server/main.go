package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/rs/cors"
	"github.com/takanome-dev/blog-with-astro-golang/internal/auth"
	"github.com/takanome-dev/blog-with-astro-golang/pkg/routes"
)

func main() {
	godotenv.Load()
	r := mux.NewRouter()

	c := cors.New(cors.Options{
    AllowedOrigins: []string{"*"},
		AllowedHeaders: []string{"*"},
    AllowCredentials: true,
		AllowedMethods: []string{http.MethodGet, http.MethodPost, http.MethodPut, http.MethodDelete, http.MethodOptions},
    // TODO: disable debug mode in production
    Debug: false,
	})
	handler := c.Handler(r)
	r.Use(auth.LoggingMiddleware)
	routes.UsersRoute(r)
	routes.PostsRoutes(r)
	routes.AuthRoute(r)
	routes.CommentsRoutes(r)


	port := os.Getenv("PORT")
	log.Printf("🚀 server listening at localhost %v 🚀\n", port)

	err := http.ListenAndServe(":" + port, handler)
	if err != nil {
		log.Fatal(err)
	}
}