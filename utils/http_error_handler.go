package utils

import (
	"github.com/boomauakim/go-todo-clean-arch/domain"
	"github.com/gofiber/fiber/v2"
)

type ErrorResponse struct {
	Error fiber.Map `json:"error"`
}

func ErrorHandler(c *fiber.Ctx, err error) error {
	switch err {
	case fiber.ErrBadRequest:
		return c.Status(fiber.StatusBadRequest).JSON(ErrorResponse{fiber.Map{
			"message": "The request was unacceptable, often due to missing a required parameter.",
		}})
	case fiber.ErrUnprocessableEntity:
		return c.Status(fiber.StatusBadRequest).JSON(ErrorResponse{fiber.Map{
			"message": "The request was unacceptable, often due to missing a required parameter.",
		}})
	case domain.ErrNotFound:
		return c.Status(fiber.StatusNotFound).JSON(ErrorResponse{fiber.Map{
			"message": "The requested resource doesn't exist.",
		}})
	default:
		return c.Status(fiber.StatusInternalServerError).JSON(ErrorResponse{fiber.Map{
			"message": "Something went wrong.",
		}})
	}
}
