package repository

import (
	go_json "github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"log"
	"mcdashboard/internal/controllers"
	"mcdashboard/internal/config"
)

func ConnectDatabase(config *pxgpool.Config) *pgxpool.Pool {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	databaseURL := os.Getenv("DATABASE_URL")
	if databaseURL == "" {
		log.Fatal("DATABASE_URL is missing")
	}

	pool, err := pgxpool.Connect(context.Background(), databaseURL)
	if err != nil {
		log.Fatalf("Unable to connect to database: %v\n", err)
	}
	defer pool.Close()
}

