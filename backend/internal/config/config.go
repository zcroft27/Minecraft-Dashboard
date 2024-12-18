package config

import (
	"fmt"
	"os"
)

type Config struct {
	DB DB
}

func (c *DB) Connection() string {
	// Get the necessary environment variables
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

	// Build the connection string
	connStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?pgbouncer=true&connection_limit=1",
		dbUser, dbPassword, dbHost, dbPort, dbName)

	return connStr
}
