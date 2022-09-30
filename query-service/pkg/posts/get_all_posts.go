package posts

import (
	"query-api/pkg/common/models"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
)

func (h handler) GetAllPosts(c *fiber.Ctx) error {
	col := h.DB.Collection("posts")

	cursor, err := col.Find(c.Context(), bson.D{})
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	var result []models.DBPost
	if err = cursor.All(c.Context(), &result); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	if len(result) == 0 {
		return c.JSON([]models.DBPost{})
	}

	return c.JSON(result)
}
