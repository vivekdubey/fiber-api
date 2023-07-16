package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/vivekdubey/fiber-api/pkg/books"
	"github.com/vivekdubey/fiber-api/pkg/common/config"
	"github.com/vivekdubey/fiber-api/pkg/common/db"
)

func main() {
	c, err := config.LoadConfig()

	if err != nil {
		log.Fatalln("Failed at config", err)
	}

	app := fiber.New()

	db := db.Init(c.DBurl)

	// app.Get("/", func(ctx *fiber.Ctx) error {
	// 	return ctx.Status(fiber.StatusOK).SendString(c.Port)
	// })

	books.RegisterRoutes(app, db)

	app.Listen(c.Port)
}
