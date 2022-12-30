package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"log"
	"poke-api-fiber/database"
	"poke-api-fiber/routes"
)

func main () {
	database.ConnectDb()
	app := fiber.New()

	app.Get("/list", routes.PokemonList)
	app.Use(cors.New())

	log.Fatal(app.Listen(":3000"))
}