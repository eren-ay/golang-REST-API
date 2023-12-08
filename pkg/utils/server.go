package utils

import (
	"github.com/gofiber/fiber/v2"
)

func CreateServer(port int) *fiber.App {
	app := fiber.New()
	return app
}
