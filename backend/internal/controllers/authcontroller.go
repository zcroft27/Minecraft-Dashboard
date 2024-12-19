package controllers

import (
	"backend/internal/auth"

	"github.com/gofiber/fiber/v2"
)

type AuthController struct{}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func NewAuthController() *AuthController {
	return &AuthController{}
}

func (ac *AuthController) Login(c *fiber.Ctx) error {
	var loginData LoginRequest

	if err := c.BodyParser(&loginData); err != nil {
		return err
	}

	email := loginData.Email
	password := loginData.Password

	resp, err := auth.GetAuthToken(&ac.config, email, password)

	if err != nil {
		return err
	}

	if err := h.store.SetUser(c, resp.User.ID); err != nil {
		return err
	}

	return c.Status(fiber.StatusOK).JSON(resp)
}
