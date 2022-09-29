package post

import (
	"post-api/pkg/common/models"

	"github.com/gofiber/fiber/v2"
)

func (h handler) ShowPost(c *fiber.Ctx) error {
	var post models.Post

	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid post id",
		})
	}

	if err := h.DB.First(&post, id).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Post not found",
		})
	}

	return c.Status(fiber.StatusOK).JSON(post)
}
