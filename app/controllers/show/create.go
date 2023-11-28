package show

import (
	"github.com/eren-ay/golang-REST-API/app/models"
	"github.com/gofiber/fiber/v2"
)

func Create(ctx *fiber.Ctx) error {
	show := models.Show{}

	if err := ctx.BodyParser(&show); err != nil {
		return ctx.Status(503).JSON(err)
	}
	/*  TO DO  database functions
	    if err := database.ConnectDB().Create(&show).Error; err != nil {
	        return ctx.Status(503).JSON(err)
	    }
	*/

	return ctx.JSON(show)
}
