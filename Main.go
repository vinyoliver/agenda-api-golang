package main

import (
	"agenda-api/agenda"
	"log"
	"net/http"

	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func main() {
	svc := agenda.NewService()
	router := agenda.NewRouter(svc)
	log.Println("Listening on port 8000")
	log.Fatal(http.ListenAndServe(":8000", router))
}
