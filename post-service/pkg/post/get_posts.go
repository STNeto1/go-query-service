package post

import (
	"post-api/pkg/common/models"

	"github.com/gofiber/fiber/v2"
)

func (h handler) GetPosts(c *fiber.Ctx) error {
	var posts []models.Post

	if err := h.DB.Find(&posts).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Error getting posts",
		})
	}

	return c.Status(fiber.StatusOK).JSON(posts)
}
