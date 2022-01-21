package handlers

import (
	"log"

	"github.com/Geeker1/goapi/db"
	"github.com/gofiber/fiber/v2"
)

func GetClubs(c *fiber.Ctx) error {
	apiDB := db.GetDB()

	return c.Render("index", fiber.Map{
		"Title": "Top Football Clubs",
		"Clubs": apiDB.GetClubs(),
	})
}

func GetClub(c *fiber.Ctx) error {
	param := c.Params(("club"))

	apiDB := db.GetDB()
	club, status := apiDB.GetClub(param)

	if status == false {
		log.Panic("Club code not found, returned ", status)
	}

	return c.Render("single", fiber.Map{
		"club": club,
	})
}
