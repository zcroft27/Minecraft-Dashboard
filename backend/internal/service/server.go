package service

import (
	"mcdashboard/internal/auth"
	"mcdashboard/internal/config"
	"mcdashboard/internal/controllers"
	storage "mcdashboard/internal/storage/postgres"

	go_json "github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v5/pgxpool"
)

type App struct {
	Server *fiber.App
	Repo   *storage.Repository
}

func InitApp(config config.Config) *App {
	repo := storage.NewRepository(config.DB)
	app := SetupApp(config, repo.DB)

	return &App{
		Server: app,
		Repo:   repo,
	}
}

func SetupApp(config config.Config, dbPool *pgxpool.Pool) *fiber.App {
	app := fiber.New(fiber.Config{
		JSONEncoder: go_json.Marshal,
		JSONDecoder: go_json.Unmarshal,
	})

	vmController := controllers.NewVMController()
	consoleController := controllers.NewConsoleController()
	signUpController := controllers.NewAuthController(&config.Supabase)

	authMiddleware := auth.Middleware(&config.Supabase, dbPool)

	app.Post("/sign-up", signUpController.Signup)
	app.Post("login", signUpController.Login)

	app.Use("/vm", authMiddleware)

	app.Route("/vm/server", func(r fiber.Router) {
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
