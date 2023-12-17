package main

import (
	dbI "cx/db"
	"cx/model"
	"cx/server"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

var db *gorm.DB

func main() {
	db, _ = dbI.SetupDB(db)

	app, err := SetupServer(db)

	PORT := ":8080"

	if err != nil {
		log.Fatal(err)
	}

	app.Listen(PORT)

}

func SetupServer(db *gorm.DB) (*fiber.App, error) {
	app := fiber.New()

	api := app.Group("/api")

	api.Route("/urls", func(r fiber.Router) {
		r.Get("/", GetLinksRoute)
		r.Post("/", CreateLinkRoute)
	})

	api.Route("/:url", func(r fiber.Router) {
		r.Get("/", GetLinkRoute)
		r.Patch("/", UpdateLinkRoute)
		r.Delete("/", DeleteLinkRoute)
	})

	return app, nil

}

func CreateLink(c *fiber.Ctx) error {

	var link = model.Link{
		ID: uuid.New().String(),
	}

	if err := c.BodyParser(&link); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Cannot parse JSON",
		})
	}

	err := server.CreateLink(db, &link)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(link)

}

func GetLinks(c *fiber.Ctx) error {

	links, err := server.GetLinks(db)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(links)
}

func GetLink(c *fiber.Ctx) error {

	link, err := server.GetLinkByShortURL(db, c.Params("url"))

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Redirect(link.URL, fiber.StatusMovedPermanently)

}

func UpdateLink(c *fiber.Ctx) error {

	link, err := server.UpdateLink(db, c.Params("url"))

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})

	}

	return c.Status(fiber.StatusOK).JSON(link)

}

func DeleteLink(c *fiber.Ctx) error {

	err := server.DeleteLink(db, c.Params("url"))

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})

	}

	return c.Status(fiber.StatusNoContent).JSON(fiber.Map{
		"message": "Link deleted",
	})

}

func CreateLinkRoute(c *fiber.Ctx) error {

	var link = model.Link{
		ID: uuid.New().String(),
	}

	if err := c.BodyParser(&link); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Cannot parse JSON",
		})
	}

	err := server.CreateLink(db, &link)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(link)

}

func GetLinksRoute(c *fiber.Ctx) error {

	links, err := server.GetLinks(db)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"count": len(links),
		"links": links,
	})
}

func GetLinkRoute(c *fiber.Ctx) error {

	link, err := server.GetLinkByShortURL(db, c.Params("url"))

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Redirect(link.URL, fiber.StatusMovedPermanently)

}

func UpdateLinkRoute(c *fiber.Ctx) error {

	link, err := server.UpdateLink(db, c.Params("url"))

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})

	}

	return c.Status(fiber.StatusOK).JSON(link)

}

func DeleteLinkRoute(c *fiber.Ctx) error {

	err := server.DeleteLink(db, c.Params("url"))

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})

	}

	return c.Status(fiber.StatusNoContent).JSON(fiber.Map{
		"message": "Link deleted",
	})

}
