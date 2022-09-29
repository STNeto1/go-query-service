package comment

import (
	"comment-api/pkg/common/exceptions"
	"errors"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type handler struct {
	DB        *gorm.DB
	Validator *validator.Validate
}

func RegisterRoutes(r *fiber.App, db *gorm.DB, v *validator.Validate) {
	h := &handler{
		DB:        db,
		Validator: v,
	}

	routes := r.Group("/comments")
	routes.Get("/", h.GetComments)
	routes.Get("/:id", h.ShowPostComments)
	routes.Post("/", h.CreateComment)
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
