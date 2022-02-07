package main

import (
	"demo-restful-api/config/database"
	"demo-restful-api/config/http"
	"log"
)

func main() {
	database.Setup()
	r := http.Setup()
	if err := r.Run("127.0.0.1:8080"); err != nil {
		log.Fatal(err)
	}
}
