package router

import "github.com/gofiber/fiber/v2"

func Router(app *fiber.App) {
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("hello world")
	})
}
