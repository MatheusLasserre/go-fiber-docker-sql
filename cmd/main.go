package main

import (
	"log"

	"github.com/MatheusLasserre/go-fiber-docker-sqloback/database"
	"github.com/MatheusLasserre/go-fiber-docker-sqloback/routes"
	"github.com/gofiber/fiber/v2"
)

func main() {
	database.InitDBConnection()

	app := fiber.New()

	routes.SetupRoutes(app)

	log.Fatal(app.Listen(":3000"))
}
