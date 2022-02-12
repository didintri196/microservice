package main

import (
	"log"
	"ms-fetch/domain"
	"ms-fetch/server/http/handlers"
	"ms-fetch/server/http/router"
	"ms-fetch/usecase"
)

func main() {
	// Load Configuration
	config, err := domain.LoadConfiguration()
	if err != nil {
		log.Fatal("Error while load configuration, ", err.Error())
	}

	// Insert Handler Contract
	var handler = handlers.NewHandler(&usecase.Contract{
		App:       config.App,
		Validator: config.Validator,
		JwtSecret: config.JwtSecret,
		Cache:     config.Cache,
		Host: struct {
			Efisery  string
			CurrConv string
		}{Efisery: config.HostEfisery, CurrConv: config.HostCurrConv},
		Apikey: struct{ CurrConv string }{CurrConv: config.ApikeyCurrConv},
	})

	// Register routes
	router.NewRouter(handler).RegisterRoutes()

	// Listening Http
	if err := domain.HttpListen(config.App); err != nil {
		log.Fatal("Error while listening http protocol, ", err.Error())
	}

}
