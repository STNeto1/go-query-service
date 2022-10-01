package post

import (
	"errors"
	"post-api/pkg/common/exceptions"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	amqp "github.com/rabbitmq/amqp091-go"
	"gorm.io/gorm"
)

type handler struct {
	DB         *gorm.DB
	Validator  *validator.Validate
	RmqChannel *amqp.Channel
	Queue      amqp.Queue
}

func RegisterRoutes(r *fiber.App, db *gorm.DB, v *validator.Validate, channel *amqp.Channel, queue amqp.Queue) {
	h := &handler{
		DB:         db,
		Validator:  v,
		RmqChannel: channel,
		Queue:      queue,
	}

	routes := r.Group("/posts")
	routes.Get("/", h.GetPosts)
	routes.Get("/:id", h.ShowPost)
	routes.Post("/", h.CreatePost)
}

func (h handler) validate(body interface{}) []exceptions.ApiError {
	err := h.Validator.Struct(body)

	if err != nil {
		var ve validator.ValidationErrors
		if errors.As(err, &ve) {
			out := make([]exceptions.ApiError, len(ve))
			for i, fe := range ve {
				out[i] = exceptions.ApiError{Param: fe.Field(), Message: exceptions.MsgForTag(fe)}
			}

			return out
		}
	}

	return []exceptions.ApiError{}
}
