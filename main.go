package main

import (
	"github.com/alanaguiar01/go-rest-api/database"
	"github.com/alanaguiar01/go-rest-api/routes"
	"github.com/gofiber/fiber/v2"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	db, err := database.ConnectDb()
	if err != nil {
		panic(err)
	}
	app := fiber.New()

	routes.CreateRoutes(app, db)

	app.Listen(":3000")
}
