package route

import (
	"gogo/middleware"

	"github.com/gofiber/fiber/v2"
)

func SetUpRoutes(app *fiber.App) {
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	app.Get("/unauthorized",middleware.Protected() ,func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	});
}
