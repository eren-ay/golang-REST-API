package main

import (
	"github.com/eren-ay/golang-REST-API/pkg/database"
	"github.com/eren-ay/golang-REST-API/pkg/routes"
	"github.com/eren-ay/golang-REST-API/pkg/utils"
)

func main() {
	database.ConnectDB()
	app := utils.CreateServer(3000)
	routes.ShowRoutes(app)

}
