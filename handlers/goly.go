package handlers

import (
	"go-fiber-url-shortener/database"
	"go-fiber-url-shortener/model"

	"github.com/gofiber/fiber/v2"
)

func CreateGoly(c *fiber.Ctx) error {
	createNewGolyRequest := new(model.Goly)

	var bodyParserError error = c.BodyParser(createNewGolyRequest)

	if bodyParserError != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "invalid incoming request values",
		})
	}

	var databaseError error = database.DB.Create(createNewGolyRequest).Error

	if databaseError != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "failed to create Goly",
		})
	}

	return c.Status(201).JSON(fiber.Map{
		"message": "success",
	})
}

func UpdateGoly(c *fiber.Ctx) error {
	return nil
}

func DeleteGoly(c *fiber.Ctx) error {
	return nil
}

func GetAllGolies(c *fiber.Ctx) error {
	return nil
}

func GolyRedirect(c *fiber.Ctx) error {
	return nil
}
