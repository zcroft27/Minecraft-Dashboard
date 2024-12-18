package storage

import (
	"context"
	"fmt"
	"log"
	"mcdashboard/internal/config"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

// Repository storage of all repositories.
type Repository struct {
	db *pgxpool.Pool
}

func (r *Repository) Close() error {
	r.db.Close()
	return nil
}

func ConnectDatabase(config config.DB) *pgxpool.Pool {
	dbConfig, err := pgxpool.ParseConfig(config.Connection())
	if err != nil {
		log.Fatal("Failed to create a config, error: ", err)
	}

	conn, err := pgxpool.NewWithConfig(context.Background(), dbConfig)
	if err != nil {
		log.Fatal("Unable to connect to database: ", err)
	}

	err = conn.Ping(context.Background())
	if err != nil {
		log.Fatal("Ping failed:", err)
	}

	log.Print("Connected to database!")

	return conn
}

func NewRepository(config config.DB) *Repository {
	db := ConnectDatabase(config)

	return &Repository{
		db: db,
	}
}

func (r *Repository) TestConnection(ctx context.Context) error {
	var currentTime time.Time
	err := r.db.QueryRow(ctx, `SELECT NOW()`).Scan(&currentTime)
	if err != nil {
		return fmt.Errorf("failed to test database connection: %w", err)
	}
	fmt.Println("Current time from DB:", currentTime)
	return nil
}
