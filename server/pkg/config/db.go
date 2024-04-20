package config

import (
	"context"
	"database/sql"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/takanome-dev/astro.go.blog/internal/database"
)

var db *database.Queries

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Printf("failed to load envs: %v", err)
	}
	
	dbUrl := os.Getenv("DB_URL")

	conn, err := sql.Open("postgres", dbUrl)
	if err != nil {
		log.Fatal("Can't connect to database:", err)
	}
	// defer conn.Close()

	ctx, cancel := context.WithTimeout(context.Background(), 500*time.Millisecond)
	defer cancel()

	err = conn.PingContext(ctx)
	if err != nil {
		log.Fatal("db connection failed:", err)
	}

	db = database.New(conn)
	log.Printf("🥳 db connected 🥳")
}

func GetDB() *database.Queries {
	return db
}