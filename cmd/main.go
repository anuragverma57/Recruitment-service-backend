package main

import (
	"log"
	"recruitment-system/config"
	"recruitment-system/routes"

	"github.com/labstack/echo/v4"
)

func main() {
	config.LoadConfig()

	if config.MongoClient == nil {
		log.Fatal("MongoDB client is not initialized. Exiting application.")
	}

	e := echo.New()

	routes.RegisterRoutes(e)

	e.Logger.Fatal(e.Start(":8080"))
}
