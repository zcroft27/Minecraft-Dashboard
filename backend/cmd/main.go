package main

import (
	"context"
	"log"
	"os"
	"github.com/jackc/pgx/v5"
)

func main() {
	conn, err := pgx.Connect(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}
	defer conn.Close(context.Background())

	// Example query to test connection
	var version string
	if err := conn.QueryRow(context.Background(), "SELECT version()").Scan(&version); err != nil {
		log.Fatalf("Query failed: %v", err)
	}

	log.Println("Connected to:", version)
}


// package main

// import (
// 	go_json "github.com/goccy/go-json"
// 	"github.com/gofiber/fiber/v2"
// 	"log"
// 	"mcdashboard/internal/controllers"
// )

// func main() {
// 	app := SetupApp()

// 	if err := app.Listen(":8080"); err != nil {
// 		log.Fatalf("Failed to start server: %v", err)
// 	}
// }

// func SetupApp() *fiber.App {
// 	app := fiber.New(fiber.Config{
// 		JSONEncoder: go_json.Marshal,
// 		JSONDecoder: go_json.Unmarshal,
// 	})

// 	vmController := controllers.NewVMController()
// 	consoleController := controllers.NewConsoleController()
// 	app.Route("/server", func(r fiber.Router) {
// 		r.Get("/start", func(c *fiber.Ctx) error {
// 			return vmController.ToggleServer(c, true)
// 		})
// 		r.Get("/stop", func(c *fiber.Ctx) error {
// 			return vmController.ToggleServer(c, false)
// 		})
// 		app.Route("/console", func(r fiber.Router) {
// 			r.Get("/list", consoleController.GetPlayerList)
// 		})
// 	})

// 	return app
// }
