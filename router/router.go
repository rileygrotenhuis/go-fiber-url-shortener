package router

import (
	"go-fiber-url-shortener/handlers"
	"go-fiber-url-shortener/middleware"

	"github.com/gofiber/fiber/v2"
)

func Initialize(router *fiber.App) {
	router.Use(middleware.Security)
	router.Use(middleware.Json)

	router.Get("/", func(c *fiber.Ctx) error {
		return c.Status(200).SendString("Hello World!")
	})

	router.Get("/:id", handlers.GolyRedirect)

	golies := router.Group("/goly")
	golies.Post("/", handlers.CreateGoly)
	golies.Put("/:id", handlers.UpdateGoly)
	golies.Delete("/:id", handlers.DeleteGoly)
}
