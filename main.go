package main

import (
    "github.com/gofiber/fiber/v2"
    "github.com/gofiber/fiber/v2/middleware/cache"
    "github.com/gofiber/fiber/v2/middleware/compress"
    "github.com/gofiber/fiber/v2/middleware/cors"
    "github.com/gofiber/fiber/v2/middleware/helmet"
    "github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
    // Start a new fiber app
    app := fiber.New()

    // middlewares
    app.Use(cache.New())
    app.Use(compress.New())
    app.Use(cors.New())
    app.Use(helmet.New())
    app.Use(logger.New())

    app.Get("/", func(c *fiber.Ctx) error {
        err := c.SendString("And the API is UP!")
        return err
    })



    // Listen on PORT 3000
    app.Listen(":3000")
}
