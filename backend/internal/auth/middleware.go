package auth

import (
	"mcdashboard/internal/config"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
	"github.com/jackc/pgx/v5/pgxpool"
)

func Middleware(config *config.Supabase, db *pgxpool.Pool) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		token := ctx.Cookies("access_token")

		if token == "" {
			return fiber.NewError(fiber.StatusBadRequest, "Unauthorized, no token")
		}
		payload, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
			return []byte(config.JWTSecret), nil
		})
		if err != nil {

			return fiber.NewError(fiber.StatusBadRequest, "Unauthorized, error parsing")
		}

		// Access the claims
		claims, ok := payload.Claims.(jwt.MapClaims)
		if !ok {
			return fiber.NewError(fiber.StatusBadRequest, "Invalid token claims")
		}

		// Retrieve the subject (user ID or UUID)
		subject, ok := claims["sub"].(string)
		if !ok {
			return fiber.NewError(fiber.StatusBadRequest, "Subject (user ID) not found in token")
		}

		// Set the user ID in the context for later use
		ctx.Locals("userId", subject)
		print(payload)
		return ctx.Next()
	}
}
