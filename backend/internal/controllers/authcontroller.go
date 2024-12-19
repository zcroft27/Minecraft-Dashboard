package controllers

import (
	"mcdashboard/internal/auth"
	"mcdashboard/internal/config"
	"time"

	"github.com/gofiber/fiber/v2"
)

type AuthController struct {
	Supabase *config.Supabase
}

func NewAuthController(supabase *config.Supabase) *AuthController {
	return &AuthController{Supabase: supabase}
}

type SignupRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (ac *AuthController) Signup(ctx *fiber.Ctx) error {
	var req SignupRequest

	// Parse the request body
	if err := ctx.BodyParser(&req); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid request")
	}

	// Sign up the user using Supabase
	response, err := auth.SupabaseSignup(ac.Supabase, req.Email, req.Password)
	if err != nil {
		return err
		// return fiber.NewError(fiber.StatusInternalServerError, "Signup failed")
	}

	// Return access token and user data to the frontend
	return ctx.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message":      "User created successfully",
		"access_token": response.AccessToken,
		"user_id":      response.User.ID,
	})
}

func (ac *AuthController) Login(ctx *fiber.Ctx) error {
	var req LoginRequest

	// Parse the request body
	if err := ctx.BodyParser(&req); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid request")
	}

	// Sign up the user using Supabase
	response, err := auth.Login(ac.Supabase, req.Email, req.Password)
	if err != nil {
		return err
		// return fiber.NewError(fiber.StatusInternalServerError, "Signup failed")
	}

	// Store the access token in the context

	// Set the access token in a cookie
	ctx.Cookie(&fiber.Cookie{
		Name:     "access_token",
		Value:    response.AccessToken,
		Expires:  time.Now().Add(24 * time.Hour),
		HTTPOnly: true,
		Secure:   true,
		SameSite: fiber.CookieSameSiteStrictMode,
	})

	// Return access token and user data to the frontend
	return ctx.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message":      "User logged in successfully",
		"access_token": response.AccessToken,
		"user_id":      response.User.ID,
	})
}
