package auth

import (
	"context"
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

		userID := ctx.Cookies("user_id")
		var role string
		err := db.QueryRow(context.Background(), "SELECT 'role' FROM auth.users WHERE id = $1", userID).Scan(&role)
		if err != nil {
			return fiber.NewError(fiber.StatusUnauthorized, "Unable to fetch role")
		}

		if role != "admin" {
			return fiber.NewError(fiber.StatusForbidden, "Access forbidden: Admins only")
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
