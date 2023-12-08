package main

import (
	"fmt"
	"log"

	"github.com/eren-ay/golang-REST-API/pkg/routes"
	"github.com/eren-ay/golang-REST-API/pkg/utils"
)

func main() {
	app := utils.CreateServer(3000)
	routes.ShowRoutes(app)
	log.Fatal(app.Listen(fmt.Sprintf(":%d", 3000)))
}
