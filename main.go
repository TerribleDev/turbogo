package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	token := os.Getenv("API_TOKEN")
	if token == "" {
		log.Fatal("API_TOKEN is required")
		os.Exit(1)
	}
	app := fiber.New()
	app.Use(logger.New(logger.Config{
		// log query string parameters
		Format: "[${ip}]:${port} ${status} - ${method} ${path} - ${queryParams}\n",
	}))
	app.Use(func(c *fiber.Ctx) error {
		authHeader := c.Get("Authorization")
		if authHeader != "Bearer "+token {
			c.Status(401).SendString("Unauthorized")
			return nil
		}
		return c.Next()
	})
	app.Get("/", func(c *fiber.Ctx) error {
		// return hello world
		return c.SendString("Hello, World!")
	})

	// apis for turborepo
	app.Get("/v2/teams", func(c *fiber.Ctx) error {
		return c.SendStatus(fiber.StatusOK)
	})

	app.Get("/v2/user", func(c *fiber.Ctx) error {
		return c.SendStatus(fiber.StatusOK)
	})

	app.Get("/v8/artifacts/:hash", func(c *fiber.Ctx) error {
		// log hello world
		// serialize query params to
		fmt.Println(string(c.Request().URI().QueryString()))
		return c.SendFile("./cache/" + c.Params("hash"))
	})
	app.Put("/v8/artifacts/:hash", func(c *fiber.Ctx) error {
		fmt.Println(string(c.Request().URI().QueryString()))
		return os.WriteFile("./cache/"+c.Params("hash"), c.Request().Body(), 0644)
	})
	port := os.Getenv("PORT")
	if port == "" {
		port = "5002"
	}

	log.Fatal(app.Listen(fmt.Sprintf(":%s", port)))
}
