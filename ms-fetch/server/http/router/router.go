package router

import (
	"ms-fetch/server/http/handlers"
	"ms-fetch/server/http/router/routes"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/fiber/v2/middleware/requestid"
)

type Router struct {
	handlers.Handler
}

func NewRouter(handler handlers.Handler) Router {
	return Router{handler}
}

var (
	logFormat = `{"host":"${host}","pid":"${pid}","time":"${time}","request_id":"${locals:requestid}","status":"${status}","method":"${method}","latency":"${latency}","path":"${path}",` +
		`"user-agent":"${ua}","bytes_in":"${bytesReceived}","bytes_out":"${bytesSent}"}`
)

func (router Router) RegisterRoutes() {

	app := router.Contract.App

	// middleware
	app.Use(recover.New())
	app.Use(requestid.New())
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowMethods: "GET,POST,HEAD,PUT,DELETE,PATCH",
	}))
	app.Use(logger.New(logger.Config{
		Format:     logFormat + "\n",
		TimeFormat: time.RFC3339,
		TimeZone:   "Asia/Jakarta",
	}))

	// Root Group
	rootGroup := app.Group("/fetch")

	// Route for check health
	rootGroup.Get("", func(ctx *fiber.Ctx) error {
		return ctx.Status(200).JSON("I am Fine Thanks!")
	})

	// Resource route
	resourceRoute := routes.NewResourceRoute(rootGroup, router.Handler)
	resourceRoute.RegisterRoute()

	// Users route
	usersRoute := routes.NewUsersRoute(rootGroup, router.Handler)
	usersRoute.RegisterRoute()

}
