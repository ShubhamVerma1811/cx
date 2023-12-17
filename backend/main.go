package main

import (
	dbI "cx/db"
	"cx/model"
	"cx/server"
	"cx/utils"
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

// TODO:: Move these to server package

func SetupServer(db *gorm.DB) (*fiber.App, error) {
	app := fiber.New()

	app.Get("/:url", RedirectByShortUrl)

	api := app.Group("/api")

	api.Route("/urls", func(r fiber.Router) {
		r.Get("/", GetAllLinks)
		r.Post("/", CreateLink)
	})

	api.Route("/:id", func(r fiber.Router) {
		r.Get("/", GetLinkById)
		r.Put("/", UpdateLinkById)
		r.Delete("/", DeleteLinkById)
	})

	return app, nil

}

func GetLinkById(c *fiber.Ctx) error {

	link, err := server.GetLinkByLinkID(db, c.Params("id"))

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(link)

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

	if link.ShortURL == "" {
		link.ShortURL = utils.RandomString(6)
	}

	err := server.CreateLink(db, &link)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(link)

}

func GetAllLinks(c *fiber.Ctx) error {

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

func RedirectByShortUrl(c *fiber.Ctx) error {

	link, err := server.GetLinkByShortURL(db, c.Params("url"))

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Redirect(link.URL, fiber.StatusMovedPermanently)

}

func UpdateLinkById(c *fiber.Ctx) error {

	link := model.Link{
		ID: c.Params("id"),
	}

	if err := c.BodyParser(&link); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Cannot parse JSON",
		})
	}

	link, err := server.UpdateLink(db, &link)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})

	}

	return c.Status(fiber.StatusOK).JSON(link)

}

func DeleteLinkById(c *fiber.Ctx) error {

	link, err := server.DeleteLink(db, c.Params("id"))

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})

	}

	return c.Status(fiber.StatusNoContent).JSON(link)

}
