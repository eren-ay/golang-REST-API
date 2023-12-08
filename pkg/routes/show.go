package routes

import (
	"github.com/eren-ay/golang-REST-API/app/controllers/show"
	"github.com/gofiber/fiber/v2"
)

func ShowRoutes(app *fiber.App) {
	// Create show
	app.Get("/show", show.Create)

	// TODO List All show id
	app.Get("/shows", show.AllId)

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, Fiber!")
	})
	// TODO Delete chosen index show

	// TODO List chosen index

	// TODO Update chosen index
}
