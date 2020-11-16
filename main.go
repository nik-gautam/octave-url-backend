package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cache"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/joho/godotenv"
	"github.com/nik-gautam/octave-url-backend/database"
	"github.com/nik-gautam/octave-url-backend/handlers"
	"log"
	"os"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic("Cannot read dotenv file")
	}

	if err := database.ConnectDB(os.Getenv("MONGO_URL")); err != nil {
		panic("DB not Connected")
	}

	println("DB connected")

	app := fiber.New()

	app.Use(cors.New())
	app.Use(logger.New())
	app.Use(cache.New())

	app.Get("/", handlers.GetAllUrls)
	app.Get("/:shortCode", handlers.GetLongUrl)
	app.Post("/", handlers.PostAddUrl)
	app.Patch("/", handlers.PatchEditUrl)
	app.Delete("/:id", handlers.DeleteUrl)

	log.Fatal(app.Listen(":3000"))

}
