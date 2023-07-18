package main

import (
	"fmt"
	"log"

	"github.com/vivekdubey/fiber-api/internal/config"
)

func main() {
	c, err := config.LoadConfig()

	if err != nil {
		log.Fatalln("Failed at config", err)
	}
	fmt.Println(c.DBurl)
	// app := fiber.New()

	// db := db.Init(c.DBurl)

	// app.Get("/", func(ctx *fiber.Ctx) error {
	// 	return ctx.Status(fiber.StatusOK).SendString(c.Port)
	// })

	// books.RegisterRoutes(app, db)

	// app.Listen(c.Port)
}
