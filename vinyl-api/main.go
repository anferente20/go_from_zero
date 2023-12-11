package main

import (
	"fmt"
	"vinyl-api/router"
)

func main() {

	fmt.Println("Welcome to Vinyl")

	router.SetRoutes()
	router.RunRouter()
}
