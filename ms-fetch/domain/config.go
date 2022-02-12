package domain

import (
	"ms-fetch/domain/constants"
	"os"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

type Config struct {
	App       *fiber.App
	Validator *validator.Validate
}

func LoadConfiguration() (config Config, err error) {
	// load env
	if err = godotenv.Load(constants.EnvironmentDirectory); err != nil {
		return config, err
	}

	// validator
	config.Validator = validator.New()

	// fiber
	config.App = fiber.New()
	return config, err
}

func HttpListen(app *fiber.App) (err error) {
	if err := app.Listen(os.Getenv(constants.EnvironmentAppRestPort)); err != nil {
		return err
	}

	return err
}
