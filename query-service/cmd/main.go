package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"query-api/pkg/common/utils"

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

	// h := db.Connect()
	// var v = validator.New()
	s := gocron.NewScheduler(time.UTC)

	app := fiber.New()

	s.Every(5).Seconds().Do(func() {
		fmt.Println("running cron")
	})

	s.StartAsync()

	app.Listen(utils.ParsePort())
}
