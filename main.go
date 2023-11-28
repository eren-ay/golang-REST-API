package main

import (
	"github.com/eren-ay/golang-REST-API/pkg/database"
	"github.com/eren-ay/golang-REST-API/pkg/utils"
)

func main() {
	utils.CreateServer(3000)
	database.ConnectionDB()
}
