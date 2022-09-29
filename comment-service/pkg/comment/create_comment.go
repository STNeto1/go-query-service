package comment

import (
	"comment-api/pkg/common/models"

	"github.com/gofiber/fiber/v2"
)

type CreateCommentRequestBody struct {
	Comment string `json:"comment" validate:"required,min=3,max=100"`
	PostID  uint   `json:"post_id" validate:"required,number"`
}

func (h handler) CreateComment(c *fiber.Ctx) error {
	var body CreateCommentRequestBody

	if err := c.BodyParser(&body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	out := h.validate(body)
	if len(out) > 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Bad Request",
			"errors":  out,
		})
	}

	comment := models.Comment{
		Comment: body.Comment,
		PostID:  body.PostID,
	}

	if err := h.DB.Create(&comment).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Something went wrong",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(comment)
}
