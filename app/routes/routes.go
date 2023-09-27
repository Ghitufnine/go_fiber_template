package routes

import (
	"github.com/gofiber/fiber/v2"
	"go_fiber_template/tests"
)

func SetUpRoutes(app *fiber.App) {
	//Your routes here
	//Example
	api := app.Group("/api")
	api.Get("/", tests.GetTest)

	// JWT Middleware
	// To use middleware uncomment this
	//middleware.JWTConfig(app)
	// Put your endpoint below here to Make auth with JWT
}
