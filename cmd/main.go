package main

import (
	"fmt"
	"log"

	"github.com/Mautjee/universal-music-link-api/internal/handlers"
	"github.com/gofiber/fiber/v2"
)

func healtcheck(c *fiber.Ctx) error {
	fmt.Println("Healthcheck: OK")
	return c.SendString("OK")
}

func main() {

	app := fiber.New()

	app.Use("/api", func(c *fiber.Ctx) error {
		fmt.Println("API middleware")

		return c.Next()
	})

	app.Get("/healthcheck", healtcheck)
	app.Post("/api/products", handlers.CreateProduct)
	app.Get("/api/products", handlers.GetAllProducts)

	log.Fatal(app.Listen(":3000"))
}
