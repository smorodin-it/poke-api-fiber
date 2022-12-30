package routes

import (
	"github.com/gofiber/fiber/v2"
	"poke-api-fiber/database"
	"poke-api-fiber/models"
)

func PokemonList (ctx *fiber.Ctx) error {
	var list []models.Pokemon
	database.DB.Db.Order("pkdx_id").Find(&list)

	return ctx.Status(200).JSON(list)
}