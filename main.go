package main

import (
	"go-fiber-url-shortener/database"
	"go-fiber-url-shortener/router"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
)

func getenv(key string, fallback string) string {
	var value string = os.Getenv(key)

	if len(value) == 0 {
		return fallback
	}

	return value
}

func main() {
	godotenv.Load()

	var app *fiber.App = fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))

	database.ConnectDB()

	router.Initialize(app)

	log.Fatal(app.Listen(":" + getenv("PORT", "3000")))
}
