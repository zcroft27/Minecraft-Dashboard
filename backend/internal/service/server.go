package service

import (
	"mcdashboard/internal/config"
	"mcdashboard/internal/controllers"
	storage "mcdashboard/internal/storage/postgres"

	go_json "github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
)

type App struct {
	Server *fiber.App
	Repo   *storage.Repository
}

func InitApp(config config.Config) *App {
	app := SetupApp()
	repo := storage.NewRepository(config.DB)

	return &App{
		Server: app,
		Repo:   repo,
	}
}

func SetupApp() *fiber.App {
	app := fiber.New(fiber.Config{
		JSONEncoder: go_json.Marshal,
		JSONDecoder: go_json.Unmarshal,
	})

	vmController := controllers.NewVMController()
	consoleController := controllers.NewConsoleController()
	app.Route("/server", func(r fiber.Router) {
		r.Get("/start", func(c *fiber.Ctx) error {
			return vmController.ToggleServer(c, true)
		})
		r.Get("/stop", func(c *fiber.Ctx) error {
			return vmController.ToggleServer(c, false)
		})
		app.Route("/console", func(r fiber.Router) {
			r.Get("/list", consoleController.GetPlayerList)
		})
	})

	return app
}
