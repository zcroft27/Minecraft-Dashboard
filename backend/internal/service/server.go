package service

import (
	"mcdashboard/internal/auth"
	"mcdashboard/internal/config"
	"mcdashboard/internal/controllers"
	storage "mcdashboard/internal/storage/postgres"

	go_json "github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
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

	app.Use(cors.New(cors.Config{
		AllowOrigins:     "http://localhost:3000", // frontend URL
		AllowCredentials: true,                    // Allow cookies to be sent
		AllowMethods:     "GET,POST,PUT,DELETE",
		AllowHeaders:     "Content-Type,Authorization",
	}))

	vmController := controllers.NewVMController()
	consoleController := controllers.NewConsoleController()
	signUpController := controllers.NewAuthController(&config.Supabase)

	// Add logging on each request
	app.Use(logger.New())

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
