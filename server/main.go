package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/takanome-dev/blog-with-astro-golang/pkg/routes"
)

func main() {
	godotenv.Load()
	r := mux.NewRouter()
	// r.Use(mux.CORSMethodMiddleware(r))
	
	routes.UsersRoute(r)
	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("PORT is not defined")
	}
	
	log.Println("----------------------------- 🚀 --------------------------------")
	log.Printf("server listening at localhost %v", port)
	log.Println("----------------------------- 🚀 --------------------------------")

	err := http.ListenAndServe(":" + port, r)
	if err != nil {
		log.Fatal(err)
	}
}