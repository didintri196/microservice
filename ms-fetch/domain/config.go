package domain

import (
	"ms-fetch/domain/constants"
	"os"
	"time"

	"github.com/patrickmn/go-cache"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

type Config struct {
	App            *fiber.App
	Validator      *validator.Validate
	JwtSecret      string
	HostEfisery    string
	HostCurrConv   string
	ApikeyCurrConv string
	Cache          *cache.Cache
}

func LoadConfiguration() (config Config, err error) {
	// load env
	if err = godotenv.Load(constants.EnvironmentDirectory); err != nil {
		return config, err
	}

	// jwt
	config.JwtSecret = os.Getenv(constants.EnvironmentJWTSecretKey)

	// host
	config.HostEfisery = os.Getenv(constants.EnvironmentHostEfisery)
	config.HostCurrConv = os.Getenv(constants.EnvironmentHostCurrConv)

	// api key
	config.ApikeyCurrConv = os.Getenv(constants.EnvironmentApiKeyCurrConv)

	// validator
	config.Validator = validator.New()

	// fiber
	config.App = fiber.New()

	//cache
	config.Cache = cache.New(5*time.Minute, 10*time.Minute)

	return config, err
}

func HttpListen(app *fiber.App) (err error) {
	if err := app.Listen(os.Getenv(constants.EnvironmentAppRestPort)); err != nil {
		return err
	}

	return err
}
