package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/rs/cors"
	"github.com/takanome-dev/astro.go.blog/internal/auth"
	"github.com/takanome-dev/astro.go.blog/pkg/routes"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Printf("failed to load envs: %v", err)
	}

	// TODO: replace allowed origins with env
	// domains := os.Getenv("DOMAINS")
	// splitDomains := strings.Split(domains, ",")
	// log.Printf("domains: %v", splitDomains)

	r := mux.NewRouter()
	c := cors.New(cors.Options{
		AllowedOrigins: []string{
			"http://localhost:4321",
			"https://astro-go-blog.vercel.app",
			"https://astro-go-blog-takanome-dev.vercel.app",
			"https://astro-go-blog-git-main-takanome-dev.vercel.app",
		},
		AllowedHeaders: []string{
			"Content-Type",
			"x-requested-with",
			"Origin",
			"Referer",
		},
		AllowCredentials: true,
		AllowedMethods:   []string{http.MethodGet, http.MethodPost, http.MethodPut, http.MethodDelete, http.MethodOptions},
		// MaxAge: 86400,
		Debug: (os.Getenv("ENV") == "development"),
	})

	handler := c.Handler(r)

	r.Use(auth.LoggingMiddleware)
	r.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, err := w.Write([]byte("OK"))
		if err != nil {
			log.Printf("failed to write response: %v", err)
		}
	})

	routes.UsersRoute(r)
	routes.PostsRoutes(r)
	routes.AuthRoute(r)
	routes.CommentsRoutes(r)

	port := os.Getenv("PORT")
	log.Printf("🚀 server listening at localhost:%s", port)

	err = http.ListenAndServe("0.0.0.0:"+port, handler)
	if err != nil {
		log.Fatal(err)
	}
}
