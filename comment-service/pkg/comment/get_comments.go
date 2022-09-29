package comment

import (
	"comment-api/pkg/common/models"

	"github.com/gofiber/fiber/v2"
)

func (h handler) GetComments(c *fiber.Ctx) error {
	var comments []models.Comment

	if err := h.DB.Find(&comments).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Error getting posts",
		})
	}

	return c.Status(fiber.StatusOK).JSON(comments)
}
