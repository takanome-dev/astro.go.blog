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
    AllowedOrigins: []string{"http://localhost:4321"},
    AllowCredentials: true,
		AllowedMethods: []string{http.MethodGet, http.MethodPost, http.MethodPut, http.MethodDelete, http.MethodOptions},
		AllowedHeaders: []string{"authorization", "content-type"},
    // Enable Debugging for testing, consider disabling in production
    Debug: true,
	})
	handler := c.Handler(r)
// r.Use(mux.CORSMethodMiddleware(r))
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