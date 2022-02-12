package usecase

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type Contract struct {
	Auth struct {
		UserName  string
		UserPhone string
		UserRole  string
	}
	App       *fiber.App
	Validator *validator.Validate
	JwtSecret string
}
