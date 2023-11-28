package utils

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
)

func CreateServer(port int) *fiber.App {
	app := fiber.New()
	log.Fatal(app.Listen(fmt.Sprintf(":%d", port)))
	return app
}
