package main

import (
	go_json "github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
	"log"
	"mcdashboard/internal/controllers"
)

func main() {
	app := SetupApp()

	if err := app.Listen(":8080"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}

func SetupApp() *fiber.App {
	app := fiber.New(fiber.Config{
		JSONEncoder: go_json.Marshal,
		JSONDecoder: go_json.Unmarshal,
	})

	vmController := controllers.NewController()
	app.Route("/vm", func(r fiber.Router) {
		r.Get("/start", vmController.StartServer)
	})

	app.Get("/start-server", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"message": "Hello, world!",
		})
	})

	app.Post("/echo", func(c *fiber.Ctx) error {
		var body map[string]interface{}
		if err := c.BodyParser(&body); err != nil {
			return err
		}
		return c.JSON(fiber.Map{
			"received": body,
		})
	})

	return app
}
