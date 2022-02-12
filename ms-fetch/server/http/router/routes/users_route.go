package routes

import (
	"ms-fetch/server/http/handlers"
	"ms-fetch/server/http/middlewares"

	"github.com/gofiber/fiber/v2"
)

type UsersRoute struct {
	RouteGroup fiber.Router
	Handler    handlers.Handler
}

func NewUsersRoute(routeGroup fiber.Router, handler handlers.Handler) UsersRoute {
	return UsersRoute{RouteGroup: routeGroup, Handler: handler}
}

func (route UsersRoute) RegisterRoute() {
	// init
	jwtMiddleware := middlewares.NewJWTMiddleware(route.Handler.Contract)

	UserRouteBiasa := route.RouteGroup.Use(jwtMiddleware.RoleBiasaOnly)

	// handlers
	usersHandler := handlers.NewUserHandler(route.Handler)

	// Users Route
	UserRouteBiasa.Get("/status", usersHandler.Status)
}
