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

	// Test the connection with a simple query
	var dbVersion string
	err = pool.QueryRow(context.Background(), "SELECT version()").Scan(&dbVersion)
	if err != nil {
		log.Fatalf("Query failed: %v\n", err)
	}
	fmt.Printf("Connected to database: %s\n", dbVersion)

	// You can now use `pool` to execute queries or transactions
}

