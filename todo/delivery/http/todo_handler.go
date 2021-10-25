package http

import (
	"github.com/boomauakim/go-todo-clean-arch/domain"
	"github.com/gofiber/fiber/v2"
)

type TodoHandler struct {
	todoUC domain.TodoUseCase
}

func NewTodoHandler(uc domain.TodoUseCase) *TodoHandler {
	return &TodoHandler{
		todoUC: uc,
	}
}

func (t *TodoHandler) ListAllTodos(c *fiber.Ctx) error {
	todos, err := t.todoUC.ListAllTodos()

	if err != nil {
		return fiber.ErrServiceUnavailable
	}
	return c.JSON(todos)
}
