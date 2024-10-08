package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/basicauth"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/proxy"
)

func init() {
	log.SetPrefix("[ARIA] ")
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}

func main() {
	go runAriaForever()
	app := fiber.New()
	app.Use(compress.New())
	app.Use(logger.New())
	if authSet {
		log.Println("Using basic auth")
		app.Use(basicauth.New(basicauth.Config{Users: map[string]string{username: password}}))
	} else {
		log.Println("Not using basic auth")
	}
	app.Get("*", func(c *fiber.Ctx) error {
		c.Set("Content-Type", "text/html")
		c.SendString(indexHtml)
		return nil
	})
	app.Post("/jsonrpc", proxy.Forward("http://localhost:"+ariaPort+"/jsonrpc"))
	log.Fatal(app.Listen(":" + port))
}
