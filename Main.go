package main

import (
	"agenda-api/database"
	"agenda-api/routes"

	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func main() {
	database.Connect()
	routes.HandleRequest()
}
