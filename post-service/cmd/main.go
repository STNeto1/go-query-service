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

	conn := utils.CreateRmqConnection("amqp://guest:guest@localhost:5672/")
	ch := utils.CreateRmqChannel(conn)
	q := utils.CreateRmqQueue(ch, "post")
	defer conn.Close()
	defer ch.Close()

	h := db.Connect()
	v := validator.New()

	app := fiber.New()

	post.RegisterRoutes(app, h, v, ch, q)

	app.Listen(utils.ParsePort())
}
