package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/takanome-dev/blog-with-astro-golang/internal/database"
	"github.com/takanome-dev/blog-with-astro-golang/pkg/routes"
)

type ApiConfig struct {
	DB *database.Queries
}

func main() {
	godotenv.Load()
	r := mux.NewRouter()
	r.Use(mux.CORSMethodMiddleware(r))
	
	routes.UsersRoute(r)
	port := os.Getenv("PORT")
	dbUrl := os.Getenv("DB_URL")

	if port == "" {
		log.Fatal("PORT is not defined")
	}
	if dbUrl == "" {
		log.Fatal("DB_URL is not defined")
	}

	conn, err := sql.Open("postgres", dbUrl)
	if err != nil {
		log.Fatal("Can't connect to database:", err)
	}

	apiCfg := ApiConfig{}
	
	log.Println("----------------------------- 🚀 --------------------------------")
	log.Printf("server listening at localhost %v", port)
	log.Println("----------------------------- 🚀 --------------------------------")

	err := http.ListenAndServe(":" + port, r)
	if err != nil {
		log.Fatal(err)
	}
}