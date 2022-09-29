package posts

import (
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
)

type handler struct {
	DB *mongo.Database
}

func RegisterRoutes(r *fiber.App, db *mongo.Database) {
	h := &handler{
		DB: db,
	}

	routes := r.Group("/posts")
	routes.Get("/", h.GetAllPosts)
}
