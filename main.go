package main

import (
	"log"

	"github.com/gofiber/fiber/v3"
)

func init() {
	log.SetPrefix("[ARIA] ")
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}

func runAria2cServerMode() {
}

func main() {
	app := fiber.New()
	app.Get("*", func(c fiber.Ctx) error {
		c.Set("Content-Type", "text/html")
		c.SendString(indexHtml)
		return nil
	})
	app.Post("/jsonrpc", func(c fiber.Ctx) error {
		return nil
	})
	log.Fatal(app.Listen(":" + port))
}
