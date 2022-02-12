package usecase

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/patrickmn/go-cache"
)

type Contract struct {
	Auth struct {
		UserName  string
		UserPhone string
		UserRole  string
	}
	Host struct {
		Efisery  string
		CurrConv string
	}
	Apikey struct {
		CurrConv string
	}
	App       *fiber.App
	Validator *validator.Validate
	Cache     *cache.Cache
	JwtSecret string
}
