package routes

import (
	"github.com/eren-ay/golang-REST-API/app/controllers/show"
	"github.com/gofiber/fiber/v2"
)

func ShowRoutes(app *fiber.App) {
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, Fiber!")
	})

	// Create show
	//app.Post("/show", show.Create)

	// TODO List All show id
	app.Get("/show/all", show.AllId)

	// TODO List chosen index
	app.Get("/show/:id", show.GetShowById)

	// TODO Delete chosen index show
	//app.Delete("/show/:id")

	// TODO Update chosen index
	//app.Post("/show/:id")
}
