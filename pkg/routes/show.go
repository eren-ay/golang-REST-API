package routes

import (
	"github.com/eren-ay/golang-REST-API/app/controllers/show"
	"github.com/gofiber/fiber/v2"
)

func ShowRoutes(app *fiber.App) {
	// Create show
	app.Post("/show", show.Create)

	// TODO List All shows
	//app.Get("/shows", show.All)

	// TODO Delete chosen index show

	// TODO List chosen index

	// TODO Update chosen index
}
