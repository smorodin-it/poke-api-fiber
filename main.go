package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"log"
	"poke-api-fiber/database"
	"poke-api-fiber/routes"
)

func main() {
	database.ConnectDb()
	app := fiber.New()

	app.Use(cors.New())
	app.Get("/list", routes.PokemonList)

	log.Fatal(app.Listen(":3000"))
}
