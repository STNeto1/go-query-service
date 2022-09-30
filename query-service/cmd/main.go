package main

import (
	"context"
	"log"
	"os"
	"time"

	"query-api/pkg/common/db"
	"query-api/pkg/common/utils"
	"query-api/pkg/posts"

	"github.com/go-co-op/gocron"
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
	defer func() {
		if err := h.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()

	app := fiber.New()
	posts.RegisterRoutes(app, h.Database("query-api"))

	s := gocron.NewScheduler(time.UTC)
	s.Every(30).Minutes().Do(func() {
		posts.UpdatePostRegisters(h.Database("query-api"))
	})
	s.StartAsync()

	app.Listen(utils.ParsePort())
}
