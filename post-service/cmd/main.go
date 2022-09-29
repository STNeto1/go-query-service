package main

import (
	"log"
	"os"

	"post-api/pkg/common/db"
	"post-api/pkg/common/utils"
	"post-api/pkg/post"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	production := os.Getenv("RAILWAY_ENVIRONMENT") == "production"

	if !production {
		err := godotenv.Load("./pkg/common/envs/.env")
		if err != nil {
			log.Fatal("Error loading .env file")
		}
	}

	h := db.Connect()
	var v = validator.New()

	app := fiber.New()

	post.RegisterRoutes(app, h, v)

	app.Listen(utils.ParsePort())
}
