package controllers

import (
	"github.com/eren-ay/golang-REST-API/app/models"
	"github.com/gofiber/fiber/v2"
)

func Index(ctx *fiber.Ctx) error {
	posts := []models.Show{}

	return ctx.JSON(posts)
}
