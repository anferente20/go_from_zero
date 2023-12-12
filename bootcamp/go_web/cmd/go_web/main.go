package main

import (
	"bootcamp_go_web/api"
	"log"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load("../../.env")
	if err != nil {
		log.Fatal("Error: Error loading .env file")
	}

	api.SetRoutes()
	api.StartApi()

}
