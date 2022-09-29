package comment

import (
	"comment-api/pkg/common/models"

	"github.com/gofiber/fiber/v2"
)

func (h handler) ShowPostComments(c *fiber.Ctx) error {
	var comments []models.Comment

	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid post id",
		})
	}

	if err := h.DB.Where("post_id = ?", id).Find(&comments).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Error getting comments",
		})
	}

	return c.Status(fiber.StatusOK).JSON(comments)
}
