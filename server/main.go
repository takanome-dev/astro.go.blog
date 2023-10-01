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
<<<<<<< HEAD
	err := godotenv.Load()
	if err != nil {
		log.Printf("failed to load envs: %v", err)
	}

=======
	godotenv.Load()
>>>>>>> origin/main
	r := mux.NewRouter()
	c := cors.New(cors.Options{
    AllowedOrigins: []string{"http://localhost:4321"},
		AllowedHeaders: []string{"*"},
    AllowCredentials: true,
		AllowedMethods: []string{http.MethodGet, http.MethodPost, http.MethodPut, http.MethodDelete, http.MethodOptions},
    Debug: false,
	})

	handler := c.Handler(r)
	r.Use(auth.LoggingMiddleware)

	r.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	})
	routes.UsersRoute(r)
	routes.PostsRoutes(r)
	routes.AuthRoute(r)
	routes.CommentsRoutes(r)


	port := os.Getenv("PORT")
	log.Printf("🚀 server listening at localhost:%s", port)

	err = http.ListenAndServe("0.0.0.0:" + port, handler)
	if err != nil {
		log.Fatal(err)
	}
}