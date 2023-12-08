package show

import (
	"github.com/eren-ay/golang-REST-API/app/models"
	"github.com/eren-ay/golang-REST-API/pkg/database"
	"github.com/gofiber/fiber/v2"
)

func Create(ctx *fiber.Ctx) error {
	//sample data
	show := models.Show{}
	show.ID = "eren"
	show.Title = "5. The Magic Door"
	showInterface := []interface{}{
		show,
	}
	database.InsertCollection(database.DB, "Show", "Movie", showInterface)

	if err := ctx.BodyParser(&show); err != nil {
		return err
	}

	return nil
}
