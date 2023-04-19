package handlers

import (
	"go-fiber-url-shortener/database"
	"go-fiber-url-shortener/model"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
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
	type UpdateGolyRequest struct {
		ShortUrl string `json:"shortUrl"`
		FullUrl  string `json:"fullUrl"`
	}

	updateGolyRequest := new(UpdateGolyRequest)

	var bodyParserError error = c.BodyParser(updateGolyRequest)

	if bodyParserError != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "invalid incoming request values",
		})
	}

	param := c.Params("id")

	id, err := strconv.Atoi(param)

	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "invalid ID format",
		})
	}

	found := model.Goly{}

	query := model.Goly{
		ID: id,
	}

	err = database.DB.First(&found, &query).Error

	if err == gorm.ErrRecordNotFound {
		return c.Status(400).JSON(fiber.Map{
			"message": "Goly not found",
		})
	}

	if updateGolyRequest.ShortUrl != "" {
		found.ShortUrl = updateGolyRequest.ShortUrl
	}

	if updateGolyRequest.FullUrl != "" {
		found.FullUrl = updateGolyRequest.FullUrl
	}

	database.DB.Save(&found)

	return c.Status(200).JSON(fiber.Map{
		"message": "success",
	})
}

func DeleteGoly(c *fiber.Ctx) error {
	param := c.Params("id")

	id, err := strconv.Atoi(param)

	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "invalid ID format",
		})
	}

	found := model.Goly{}

	query := model.Goly{
		ID: id,
	}

	err = database.DB.First(&found, &query).Error

	if err == gorm.ErrRecordNotFound {
		return c.Status(400).JSON(fiber.Map{
			"message": "Goly not found",
		})
	}

	database.DB.Delete(&found)

	return c.Status(204).JSON(fiber.Map{
		"message": "success",
	})
}

func GolyRedirect(c *fiber.Ctx) error {
	shortUrl := c.Params("id")

	found := model.Goly{}

	query := model.Goly{
		ShortUrl: shortUrl,
	}

	err := database.DB.First(&found, &query).Error

	if err == gorm.ErrRecordNotFound {
		return c.Status(400).JSON(fiber.Map{
			"message": "Goly not found",
		})
	}

	return c.Redirect(found.FullUrl, fiber.StatusTemporaryRedirect)
}
