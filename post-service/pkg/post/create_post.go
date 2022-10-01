package post

import (
	"context"
	"encoding/json"
	"fmt"
	"post-api/pkg/common/models"

	"github.com/gofiber/fiber/v2"
	amqp "github.com/rabbitmq/amqp091-go"
)

type CreatePostRequestBody struct {
	Title   string `json:"title" validate:"required,min=3,max=100"`
	Content string `json:"content" validate:"required,min=3,max=1000"`
}

func (h handler) CreatePost(c *fiber.Ctx) error {
	var body CreatePostRequestBody

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

	post := models.Post{
		Title:   body.Title,
		Content: body.Content,
	}

	if err := h.DB.Create(&post).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Something went wrong",
		})
	}

	jsonPost, err := json.Marshal(post)
	if err != nil {
		fmt.Println(err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Something went wrong",
		})
	}

	err = h.RmqChannel.PublishWithContext(context.Background(), "", h.Queue.Name, false, false, amqp.Publishing{
		ContentType: "application/json",
		Body:        jsonPost,
	})

	if err != nil {
		fmt.Println(err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Something went wrong",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(post)
}
