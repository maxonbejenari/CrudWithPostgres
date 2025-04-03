package main

import (
	"github.com/gofiber/fiber/v2"
	"log"
)

func main() {
	app := fiber.New()

	app.Get("api/hello", func(c *fiber.Ctx) error {
		return c.SendString("Hello")
	})

	log.Fatal(app.Listen(":8000"))
}
