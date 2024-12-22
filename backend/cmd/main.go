package main

import (
	"context"
	"log"
	"mcdashboard/internal/config"
	"mcdashboard/internal/service"

	"github.com/joho/godotenv"

	"github.com/sethvargo/go-envconfig"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
		return
	}

	var config config.Config
	if err := envconfig.Process(context.Background(), &config); err != nil {
		log.Fatalln("Error processing .env file: ", err)
	}

	app := service.InitApp(config)

	// err = app.Repo.TestConnection(context.Background())
	// if err != nil {
	// 	log.Fatalf("Error connecting to the database: %v", err)
	// 	return
	// }

	defer app.Repo.Close()

	if err := app.Server.Listen(":" + "5432"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
