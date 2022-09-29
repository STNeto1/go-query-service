package main

import (
	"log"
	"os"

	"comment-api/pkg/comment"
	"comment-api/pkg/common/db"
	"comment-api/pkg/common/utils"

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

	comment.RegisterRoutes(app, h, v)

	app.Listen(utils.ParsePort())
}
