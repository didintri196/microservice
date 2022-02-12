package routes

import (
	"ms-fetch/server/http/handlers"
	"ms-fetch/server/http/middlewares"

	"github.com/gofiber/fiber/v2"
)

type ResourceRoute struct {
	RouteGroup fiber.Router
	Handler    handlers.Handler
}

func NewResourceRoute(routeGroup fiber.Router, handler handlers.Handler) ResourceRoute {
	return ResourceRoute{RouteGroup: routeGroup, Handler: handler}
}

func (route ResourceRoute) RegisterRoute() {

	// init
	jwtMiddleware := middlewares.NewJWTMiddleware(route.Handler.Contract)

	fetchRouteBiasa := route.RouteGroup.Group("/user/resource").Use(jwtMiddleware.RoleBiasaOnly)
	fetchRouteAdmin := route.RouteGroup.Group("/admin/resource").Use(jwtMiddleware.RoleAdminOnly)

	// handlers
	resourceHandler := handlers.NewResourceHandler(route.Handler)

	// Resource Route
	fetchRouteBiasa.Get("", resourceHandler.GetResource)
	fetchRouteBiasa.Get("/currency-usd", resourceHandler.GetResourceWithUSD)

	fetchRouteAdmin.Get("/report-weekly", resourceHandler.GetReportWeekResource)

}
