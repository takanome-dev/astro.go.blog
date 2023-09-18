package config

import (
	"database/sql"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/takanome-dev/blog-with-astro-golang/internal/database"
)

var db *database.Queries

func init() {
	godotenv.Load()
	
	dbUrl := os.Getenv("DB_URL")
	if dbUrl == "" {
		log.Fatal("DB_URL is not defined")
	}

	conn, err := sql.Open("postgres", dbUrl)
	if err != nil {
		log.Fatal("Can't connect to database:", err)
	}

	err = conn.Ping()
	if err != nil {
		log.Fatal("db connection failed:", err)
	}

	db = database.New(conn)
	log.Printf("⚙ connected to your database ⚙")
}

func GetDB() *database.Queries {
	return db
}