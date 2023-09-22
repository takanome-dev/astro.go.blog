package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/takanome-dev/blog-with-astro-golang/internal/auth"
	"github.com/takanome-dev/blog-with-astro-golang/pkg/routes"
)

func main() {
	godotenv.Load()
	r := mux.NewRouter()
	
	r.Use(mux.CORSMethodMiddleware(r))
	r.Use(auth.LoggingMiddleware)
	
	routes.UsersRoute(r)
	routes.PostsRoutes(r)
	routes.AuthRoute(r)

	port := os.Getenv("PORT")
	log.Printf("🚀 server listening at localhost %v 🚀\n", port)

	err := http.ListenAndServe(":" + port, r)
	if err != nil {
		log.Fatal(err)
	}
}