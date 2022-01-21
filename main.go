package main

import (
	"github.com/Geeker1/goapi/db"
	"github.com/Geeker1/goapi/handlers"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html"
)

func main() {
	engine := html.New("./views", ".html")

	_, err := db.InitDB()
	if err != nil { panic(err) }
	
	fiberConfig := fiber.Config{
		Views: engine,
	}

	app := fiber.New(fiberConfig)

	app.Use(func(c *fiber.Ctx) error {
		return c.Next()
	})

	app.Get("/", handlers.GetClubs)
	app.Get("/clubs", handlers.GetClubs)
	app.Get("/clubs/:club", handlers.GetClub)

	app.Listen(":3000")
}
